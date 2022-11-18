package base

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
)

func Static(request *ghttp.Request) {
	request.Response.Status = 200
	file := gres.Get("public/index.html")
	if file != nil {
		request.Response.Write(file.Content())
	} else if data := gfile.GetBytes("public/index.html"); data != nil {
		request.Response.Write(data)
	} else {
		request.Response.Status = 404
	}
}
