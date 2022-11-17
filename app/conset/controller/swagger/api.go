package swagger

import (
	"context"
	"fmt"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/net/goai"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/samber/lo"
)

var openapi *goai.OpenApiV3

func initOpenApi(r *ghttp.Request) {
	if openapi != nil {
		return
	}
	openapi = goai.New()
	var (
		ctx    = context.TODO()
		err    error
		method string
	)
	for _, item := range r.Server.GetRoutes() {
		switch item.Type {
		case ghttp.HandlerTypeMiddleware, ghttp.HandlerTypeHook:
			continue
		}
		method = item.Method
		if gstr.Equal(method, "ALL") {
			method = ""
		}
		//log.Println(gjson.MustEncodeString(item))
		if !lo.Contains([]string{
			"/api/bin/fetch",
			"/api/bin/get",
		}, item.Route) {
			continue
		}
		if item.Handler.Info.Func == nil {
			err = openapi.Add(goai.AddInput{
				Path:   item.Route,
				Method: method,
				Object: item.Handler.Info.Value.Interface(),
			})
			if err != nil {
				r.Server.Logger().Fatalf(ctx, `%+v`, err)
			}
		}
	}
}
func OpenApi(r *ghttp.Request) {
	initOpenApi(r)
	r.Response.WriteJson(openapi)
}

const (
	swaggerUIDocName            = `redoc.standalone.js`
	swaggerUIDocNamePlaceHolder = `{SwaggerUIDocName}`
	swaggerUIDocURLPlaceHolder  = `{SwaggerUIDocUrl}`
	swaggerUITemplate           = `
<!DOCTYPE html>
<html>
	<head>
	<title>API Reference</title>
	<meta charset="utf-8"/>
	<meta name="viewport" content="width=device-width, initial-scale=1">
	<style>
		body {
			margin:  0;
			padding: 0;
		}
	</style>
	</head>
	<body>
		<redoc spec-url="{SwaggerUIDocUrl}" show-object-schema-examples="true"></redoc>
		<script src="https://unpkg.com/redoc@2.0.0-rc.70/bundles/redoc.standalone.js"> </script>
	</body>
</html>
`
)

func UI(r *ghttp.Request) {
	content := gstr.ReplaceByMap(swaggerUITemplate, map[string]string{
		swaggerUIDocURLPlaceHolder:  "/api.json",
		swaggerUIDocNamePlaceHolder: gstr.TrimRight(fmt.Sprintf(`//%s%s`, r.Host, "/swagger"), "/") + "/" + swaggerUIDocName,
	})
	r.Response.Write(content)
	r.ExitAll()
}
