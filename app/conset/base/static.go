package base

import (
	"bytes"
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
			if bytes.HasPrefix(ret, []byte(`<?xml`)) {
				request.Response.Header().Set("Content-Type", "image/svg+xml")
			}
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
	if ret := sniff("public/index.html"); ret != nil {
		ret = bytes.Replace(ret, []byte(`默认标题`), []byte(DefaultSetting.Title), 1)
		ret = bytes.Replace(ret, []byte(`/favicon.ico`), []byte(DefaultSetting.Icon), 1)
		request.Response.Write(ret)
		return
	}
	request.Response.Status = 404
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
