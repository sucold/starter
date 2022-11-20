package admin

import (
	"context"
	v8 "github.com/hinego/conset/api/v8"
	"github.com/hinego/conset/base"
	"github.com/hinego/errorx"
)

type configController struct{}

var (
	Config = configController{}
)

func (c *configController) Get(ctx context.Context, req *v8.ConfigGetReq) (res *v8.ConfigGetRes, err error) {
	return nil, errorx.NewCode(0, "success", base.DefaultSetting)
}
func (c *configController) Update(ctx context.Context, req *v8.ConfigUpdateReq) (res *v8.ConfigUpdateRes, err error) {
	base.DefaultSetting.Desc = req.Desc
	base.DefaultSetting.Register = req.Register
	base.DefaultSetting.Forget = req.Forget
	base.DefaultSetting.Verify = req.Verify
	base.DefaultSetting.Icon = req.Icon
	base.DefaultSetting.Logo = req.Logo
	base.DefaultSetting.Title = req.Title
	base.SaveConfig()
	return nil, errorx.NewCode(0, "更新成功", nil)
}
