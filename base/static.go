package base

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"strings"
)

func Static(request *ghttp.Request) {
	request.Response.Status = 200
	url := strings.ToLower(request.RequestURI)
	if strings.HasPrefix(url, "/logo.svg") {
		if ret := sniff("public/save_logo.svg"); ret != nil {
			request.Response.Write(ret)
			return
		}
	}
	if strings.HasPrefix(url, "/favicon.ico") {
		if ret := sniff("public/save_favicon.ico"); ret != nil {
			request.Response.Write(ret)
			return
		}
	}
	file := gres.Get("public/index.html")
	if file != nil {
		request.Response.Write(file.Content())
	} else if data := gfile.GetBytes("public/index.html"); data != nil {
		request.Response.Write(data)
	} else {
		request.Response.Status = 404
	}
}
func sniff(name string) []byte {
	if ret := gfile.GetBytes(name); ret != nil {
		return ret
	}
	if file := gres.Get(name); file != nil {
		return file.Content()
	}
	return nil
}
