package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/hinego/starter/app/conset/controller"
	"github.com/hinego/starter/app/conset/controller/admin"
	"github.com/hinego/starter/app/conset/controller/swagger"
	"github.com/hinego/starter/app/response"
	"github.com/hinego/starter/app/service"
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
			controller.Auth,
		)
		group.Group("/", func(group *ghttp.RouterGroup) {
			group.Middleware(service.Auth.Middleware)
			group.Bind(
				controller.Authed,
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
				)
			})
		})
	})
	s.BindStatusHandler(404, config.Proxy)
	s.Run()
	s.Logger()
	return nil
}
func mainFunc(ctx context.Context, parser *gcmd.Parser) (err error) {
	app := []func() error{
		config.initDatabase, //连接数据库
		service.StartAuth,   //启动JWT
		config.initUser,     //初始化用户
		mainWeb,             //启动web服务
	}
	return tox.WithError(app...)
}
func AddCommand(command *gcmd.Command) {
	_ = Main.AddCommand(command)
}
