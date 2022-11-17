package admin

import (
	"context"
	"errors"
	"github.com/hinego/conset/api/v8"
	"github.com/hinego/errorx"
	"github.com/hinego/gen/field"
	"github.com/hinego/starter/app/consts"
	"github.com/hinego/starter/app/dao"
	"github.com/hinego/types"
)

type userController struct{}

var (
	User = userController{}
)

func (c *userController) Fetch(ctx context.Context, req *v8.UserFetchReq) (res *v8.UserFetchRes, err error) {
	//var (
	//r  = g.RequestFromCtx(ctx)
	//id = r.GetParam(consts.UserKey).Int64()
	//)
	res = &v8.UserFetchRes{
		PageReq: &req.PageReq,
		PageRes: &types.PageRes{},
		Data:    make([]*v8.FetchUser, 0),
	}
	var u = dao.User
	if res.Total, err = u.Order(u.ID.Desc()).ScanByPage(&res.Data, res.Offset(), res.Size); err != nil {
		return nil, err
	}
	return
}

func (c *userController) Update(ctx context.Context, req *v8.UserUpdateReq) (res *v8.UserUpdateRes, err error) {
	var (
		u    = dao.User
		data []field.AssignExpr
	)
	if req.ID == 1 && req.Role != "admin" {
		return nil, errors.New("ID为1的用户必须为管理员（防止全部都成为普通用户后无法管理后台）")
	}
	if req.Name != "" {
		data = append(data, u.Name.Value(req.Name))
	}
	if req.Password != "" {
		data = append(data, u.Name.Value(consts.GeneratorPassword(req.Password)))
	}
	if req.Role != "" {
		data = append(data, u.Role.Value(req.Role))
	}
	_, err = u.Where(u.ID.Eq(req.ID)).UpdateSimple(data...)
	if err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "更新成功", nil)
}
