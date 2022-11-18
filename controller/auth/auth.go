package auth

import (
	"bytes"
	"context"
	"errors"
	"fmt"
	"github.com/gogf/gf/v2/container/gvar"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/util/gconv"
	"github.com/gogf/gf/v2/util/grand"
	"github.com/hinego/authentic"
	"github.com/hinego/conset/api"
	"github.com/hinego/conset/base"
	"github.com/hinego/conset/cache"
	"github.com/hinego/conset/service"
	"github.com/hinego/conset/service/mail"
	"github.com/hinego/errorx"
	"github.com/hinego/starter/app/dao"
	"github.com/hinego/starter/app/model"
	"github.com/samber/lo"
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
	"strings"
	"time"
)

type authController struct{}
type authedController struct{}

var (
	Auth                 = authController{}
	Authed               = authedController{}
	ErrIncorrectUserName = errorx.New("账号或密码错误")
)

func (c *authController) Login(ctx context.Context, req *api.AuthLoginReq) (res *api.AuthLoginRes, err error) {
	var (
		user *model.User
	)
	if user, err = dao.User.Where(dao.User.Email.Like(req.Username)).First(); err != nil {
		return nil, ErrIncorrectUserName
	} else if err = bcrypt.CompareHashAndPassword([]byte(user.Password), base.Salt(req.Password)); err != nil {
		return nil, ErrIncorrectUserName
	} else {
		var token *authentic.Context
		if token, err = service.Auth.CreateToken(ctx, map[string]any{base.UserKey: user.ID, "role": user.Role}); err != nil {
			return nil, err
		} else {
			ret := map[string]any{
				"expire": token.Expire.Unix(),
				"token":  token.Token.Raw,
			}
			return nil, errorx.NewCode(0, "登录成功", ret)
		}
	}
}
func (c *authController) Register(ctx context.Context, req *api.AuthRegisterReq) (res *api.AuthRegisterRes, err error) {
	var (
		r        = g.RequestFromCtx(ctx)
		password []byte
	)

	user := &model.User{
		Name:  req.Username,
		Email: req.Username,
		IP:    r.GetClientIp(),
	}
	if password, err = bcrypt.GenerateFromPassword(base.Salt(req.Password), bcrypt.DefaultCost); err != nil {
		return nil, err
	} else {
		user.Password = string(password)
	}
	send := &api.AuthSendReq{
		Mail: req.Username,
		Type: "register",
	}
	if err = c.verifyCode(ctx, send, req.Code); err != nil {
		return nil, err
	}
	if err = dao.User.Create(user); err != nil {
		if strings.Contains(err.Error(), "duplicate key") {
			return nil, errorx.New("邮箱已在系统中存在")
		}
		return nil, err
	} else {
		return nil, errorx.NewCode(0, "注册成功", nil)
	}
}
func (c *authController) Send(ctx context.Context, req *api.AuthSendReq) (res *api.AuthSendRes, err error) {
	if !lo.Contains([]string{"register", "forget"}, req.Type) {
		return nil, errorx.New("不支持的验证码类型")
	}
	key := c.sendKey(req.Mail, req.Type)
	var code = gconv.String(grand.N(100000, 999999))
	var expire = time.Duration(mail.Conf.Expire) * time.Second
	var balance time.Duration
	if balance, err = cache.GetExpire(ctx, key); err == nil {
		ex := int64(expire.Seconds()) - int64(balance.Seconds())
		if ex < mail.Conf.DiffTime {
			return nil, errorx.New("验证码已发送，请过一会儿再请求")
		}
	}
	params := map[string]any{
		"name":   base.AppName,
		"code":   code,
		"expire": int64(expire.Seconds()) / 60,
		"url":    "https://baidu.com",
	}
	u := dao.User
	if _, err = u.Where(u.Email.Like(req.Mail)).First(); err != nil {
		if req.Type == "forget" {
			if errors.Is(err, gorm.ErrRecordNotFound) {
				return nil, errorx.New("邮箱未注册")
			}
			return nil, err
		}
	} else {
		if req.Type == "register" {
			return nil, errorx.New("邮箱已在系统中存在")
		}
	}
	if err = mail.Send(ctx, req.Mail, mail.RegisterTPL, params); err != nil {
		return nil, err
	}
	if err = cache.Set(ctx, c.sendKey(req.Mail, req.Type), code, expire); err != nil {
		return nil, errorx.New("缓存验证码失败")
	}
	errKey := c.sendKey(req.Mail, req.Type+"|error")
	_, _ = cache.Remove(ctx, errKey)
	return nil, errorx.NewCode(0, "发送成功", code)
}
func (c *authController) verifyCode(ctx context.Context, req *api.AuthSendReq, code string) error {
	var expire = time.Duration(mail.Conf.Expire) * time.Second
	if get, err := cache.Get(ctx, c.sendKey(req.Mail, req.Type)); err != nil {
		return errorx.New("请先获取邮箱验证码")
	} else {
		if !gconv.Bool(get) {
			return errorx.New("请先获取邮箱验证码")
		}
		errKey := c.sendKey(req.Mail, req.Type+"|error")
		var fail *gvar.Var
		if fail, err = cache.Get(ctx, errKey); err == nil {
			if fail.Int64() >= mail.Conf.ErrorTimes {
				return errorx.New(fmt.Sprintf("错误次数过多，请重新获取"))
			}
		}
		if get.String() != code {
			var times = fail.Int64()
			if err = cache.Set(ctx, errKey, times+1, expire); err != nil {
				return errorx.New("更新验证码错误次数失败" + err.Error())
			}
			return errorx.New("验证码错误")
		}
		return nil
	}

}
func (c *authController) sendKey(email string, typ string) string {
	var buffer bytes.Buffer
	buffer.WriteString("send|")
	buffer.WriteString(typ)
	buffer.WriteString("|")
	buffer.WriteString(email)
	return buffer.String()
}
func (c *authController) Forget(ctx context.Context, req *api.AuthForgetReq) (res *api.AuthForgetRes, err error) {
	var password []byte
	send := &api.AuthSendReq{
		Mail: req.Username,
		Type: "forget",
	}
	if err = c.verifyCode(ctx, send, req.Code); err != nil {
		return nil, err
	}
	u := dao.User
	if password, err = bcrypt.GenerateFromPassword(base.Salt(req.Password), bcrypt.DefaultCost); err != nil {
		return nil, err
	} else if _, err = u.Where(u.Email.Like(req.Username)).UpdateSimple(u.Password.Value(string(password))); err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "重置成功", nil)
}
func (c *authedController) Session(ctx context.Context, req *api.AuthSessionReq) (res *api.AuthSessionRes, err error) {
	return
}
func (c *authedController) User(ctx context.Context, req *api.AuthUserReq) (res *api.AuthUserRes, err error) {
	var (
		u    = dao.User
		r    = g.RequestFromCtx(ctx)
		id   = r.GetParam(base.UserKey).Int64()
		user *model.User
	)
	if user, err = u.Where(u.ID.Eq(id)).First(); err != nil {
		return nil, err
	}
	data := api.AuthUserRes(user)
	res = &data
	return
}
func (c *authedController) Logout(ctx context.Context, req *api.AuthLogoutReq) (res *api.AuthLogoutRes, err error) {
	return nil, service.Auth.LogoutHandler(ctx)
}
func (c *authedController) Refresh(ctx context.Context, req *api.AuthRefreshReq) (res *api.AuthRefreshRes, err error) {
	return nil, service.Auth.RefreshHandler(ctx)
}
func (c *authedController) Reset(ctx context.Context, req *api.AuthResetReq) (res *api.AuthResetRes, err error) {
	var (
		u        = dao.User
		r        = g.RequestFromCtx(ctx)
		id       = r.GetParam(base.UserKey).Int64()
		password []byte
	)
	if password, err = bcrypt.GenerateFromPassword(base.Salt(req.Password), bcrypt.DefaultCost); err != nil {
		return nil, err
	} else if _, err = u.Where(u.ID.Eq(id)).UpdateSimple(u.Password.Value(string(password))); err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "修改成功", nil)
}
