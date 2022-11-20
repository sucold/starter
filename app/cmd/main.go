package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/hinego/starter/app/conset/base"
	"github.com/hinego/starter/app/conset/boot"
	"github.com/hinego/starter/app/conset/controller/admin"
	"github.com/hinego/starter/app/conset/controller/auth"
	"github.com/hinego/starter/app/conset/controller/swagger"
	"github.com/hinego/starter/app/conset/database"
	"github.com/hinego/starter/app/conset/response"
	"github.com/hinego/starter/app/conset/service"
	"github.com/hinego/starter/app/conset/tab"
	"github.com/hinego/tox"
)

var Main = gcmd.Command{
	Name:      "main",
	Usage:     "main",
	Brief:     "start main server",
	Arguments: []gcmd.Argument{},
	Func:      mainFunc,
}

func mainWeb() error {
	s := g.Server()
	s.BindHandler("/api.json", swagger.OpenApi)
	s.BindHandler("/swagger", swagger.UI)
	s.Group("/api", func(group *ghttp.RouterGroup) {
		group.Middleware(response.Access)
		group.Middleware(response.Handler)
		group.Bind(
			auth.Auth,
		)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Auth.Middleware)
			group.Bind(
				auth.Authed,
				admin.Token,
			)
			group.Group("/admin", func(group *ghttp.RouterGroup) {
				group.Middleware(service.Auth.MiddlewareWithOption(func(data map[string]any) bool {
					if role, ok := data["role"]; !ok {
						return false
					} else {
						return role == "admin"
					}
				}))
				group.Bind(
					admin.User,
					admin.Config,
				)
			})
		})
	})
	s.BindStatusHandler(404, base.Static)
	s.Run()
	s.Logger()
	return nil
}

func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	app := []func() error{
		database.Init(tab.User{}, tab.Token{}), //连接数据库
		service.StartAuth,                      //启动JWT
		boot.InitUser,
		mainWeb, //启动web服务
	}
	return tox.WithError(app...)
}
func AddCommand(command *gcmd.Command) {
	_ = Main.AddCommand(command)
}
