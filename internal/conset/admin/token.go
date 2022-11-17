package admin

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/errorx"
	"github.com/hinego/starter/api"
	"github.com/hinego/starter/internal/conset/api/v8"
	"github.com/hinego/starter/internal/consts"
	"github.com/hinego/starter/internal/dao"
	"github.com/hinego/starter/internal/model"
	"github.com/hinego/starter/internal/service"
)

type tokenController struct{}

var (
	Token = tokenController{}
)

func (c *tokenController) Fetch(ctx context.Context, req *v8.TokenFetchReq) (res *v8.TokenFetchRes, err error) {
	var (
		r  = g.RequestFromCtx(ctx)
		id = r.GetParam(consts.UserKey).Int64()
	)
	res = &v8.TokenFetchRes{
		PageReq: &req.PageReq,
		PageRes: &api.PageRes{},
		Data:    make([]*model.Token, 0),
	}
	var u = dao.Token
	if res.Total, err = u.Where(u.Role.Like("api"), u.UserID.Eq(id)).Order(u.ID.Desc()).ScanByPage(&res.Data, res.Offset(), res.Size); err != nil {
		return nil, err
	}
	return
}

func (c *tokenController) Create(ctx context.Context, req *v8.TokenCreateReq) (res *v8.TokenCreateRes, err error) {
	var (
		u    = dao.User
		r    = g.RequestFromCtx(ctx)
		id   = r.GetParam(consts.UserKey).Int64()
		user *model.User
	)

	if user, err = u.Where(u.ID.Eq(id)).First(); err != nil {
		return
	}
	if _, err = service.Auth.CreateToken(ctx, map[string]any{consts.UserKey: user.ID, "role": user.Role, "token_type": "api"}); err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "创建成功", nil)
}

func (c *tokenController) Delete(ctx context.Context, req *v8.TokenDeleteReq) (res *v8.TokenDeleteRes, err error) {
	var (
		u     = dao.Token
		r     = g.RequestFromCtx(ctx)
		id    = r.GetParam(consts.UserKey).Int64()
		token *model.Token
	)
	if token, err = u.Where(u.ID.Eq(req.ID), u.UserID.Eq(id)).First(); err != nil {
		return nil, err
	} else {
		if _, err = u.Where(u.ID.Eq(req.ID), u.UserID.Eq(id)).Delete(); err != nil {
			return nil, err
		}
		if _, err = service.Auth.Cache.Remove(ctx, token.Code); err != nil {
			return nil, err
		}
	}
	return nil, errorx.NewCode(0, "删除成功", nil)
}
