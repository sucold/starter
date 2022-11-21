package auth

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	v7 "github.com/hinego/conset/api/v7"
	"github.com/hinego/conset/base"
	"github.com/hinego/errorx"
	"github.com/hinego/gen/field"
	"github.com/hinego/starter/app/dao"
	"github.com/hinego/starter/app/model"
)

type userController struct{}

var (
	User = userController{}
)

func (c *userController) Get(ctx context.Context, req *v7.UserGetReq) (res *v7.UserGetRes, err error) {
	var (
		u    = dao.User
		r    = g.RequestFromCtx(ctx)
		id   = r.GetParam(base.UserKey).Int64()
		user *model.User
	)
	user, err = u.Where(u.ID.Eq(id)).First()
	if err != nil {
		return nil, err
	}
	return &v7.UserGetRes{
		ID:        user.ID,
		CreatedAt: user.CreatedAt,
		UpdatedAt: user.UpdatedAt,
		Name:      user.Name,
		Email:     user.Email,
		Role:      user.Role,
	}, nil

}

func (c *userController) Update(ctx context.Context, req *v7.UserUpdateReq) (res *v7.UserUpdateRes, err error) {
	var (
		u     = dao.User
		r     = g.RequestFromCtx(ctx)
		id    = r.GetParam(base.UserKey).Int64()
		where = make([]field.AssignExpr, 0)
	)
	where = append(where, u.Name.Value(req.Name))
	if req.Password != "" {
		if req.Password == req.Password1 {
			return nil, errorx.NewCode(-1, "新密码不能与旧密码相同", nil)
		}
		where = append(where, u.Password.Value(base.GeneratorPassword(req.Password1)))
	}
	_, err = u.Where(u.ID.Eq(id)).UpdateSimple(where...)
	if err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "更新成功", nil)
}
