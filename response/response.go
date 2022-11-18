package response

import (
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/hinego/errorx"
)

type Body struct {
	Code int    `json:"code"`
	Msg  string `json:"msg"`
	Data any    `json:"data,omitempty"`
}

func Handler(r *ghttp.Request) {
	r.Middleware.Next()
	if r.Response.BufferLength() > 0 {
		return
	}
	var (
		err = r.GetError()
		res = r.GetHandlerResponse()
	)
	var body Body
	if err != nil {
		switch e := err.(type) {
		case *errorx.CodeError:
			body.Code = e.Code
			body.Msg = err.Error()
			body.Data = e.Data
		default:
			body.Code = -1
			body.Msg = err.Error()
		}
	} else {
		body.Msg = "OK"
		body.Data = res
	}
	if body.Code >= 0 {
		r.SetError(nil)
	}
	r.Response.WriteJsonExit(body)
}
