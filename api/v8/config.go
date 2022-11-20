package v8

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hinego/conset/base"
)

type ConfigGetReq struct {
	g.Meta `path:"/config/get" method:"get" sm:"更新配置" tags:"系统设置"`
}
type ConfigGetRes base.Setting

type ConfigUpdateReq struct {
	g.Meta   `path:"/config/update" method:"post" sm:"更新配置" tags:"系统设置"`
	Icon     string `json:"icon" dc:"ICON图标"`
	Register bool   `json:"register" dc:"开放注册"`
	Forget   bool   `json:"forget" dc:"开放找回"`
	Verify   bool   `json:"verify" dc:"注册验证"`
	Desc     string `json:"desc" dc:"登录说明"`
	Logo     string `json:"logo" dc:"Logo地址"`
	Title    string `json:"title" dc:"网站名称"`
}
type ConfigUpdateRes struct {
}

type ConfigUploadReq struct {
	g.Meta `path:"/config/upload" method:"post" sm:"更新配置" tags:"系统设置"`
	File   *ghttp.UploadFile `json:"files" type:"file" dc:"上传文件"`
	Name   string            `json:"name" dc:"上传类型（logo.svg|favicon.ico）" v:"in:logo.svg,favicon.ico"`
}
type ConfigUploadRes struct {
}
