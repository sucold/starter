package utility

import (
	"context"
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/frame/g"
	"strings"
)

func Get(r *resty.Client, method string, path string, body any, data any) error {
	method = strings.ToLower(method)
	debug := false
	if strings.Contains(path, "/do_debug/") {
		debug = true
		path = strings.ReplaceAll(path, "/do_debug/", "/")
	}
	req := r.R()
	if body != nil {
		switch method {
		case "postform":
			if body1, ok := body.(map[string]string); ok {
				req.SetFormData(body1)
			} else {
				return errors.New("输入错误")
			}
		default:
			req.SetBody(body)
			req.SetHeader("Content-Type", "application/json")
		}
	}
	var resp *resty.Response
	var err error
	switch method {
	case "get":
		resp, err = req.Get(path)
		break
	case "post", "postform":
		resp, err = req.Post(path)
		break
	case "put":
		resp, err = req.Put(path)
		break
	case "delete":
		resp, err = req.Put(path)
		break
	default:
		return errors.New("不支持的请求方式")
	}
	if debug {
		g.Log().Debug(context.TODO(), "Request URL ", req.URL)
		g.Log().Debug(context.TODO(), "Request Body ", string(resp.Body()))
		g.Log().Debug(context.TODO(), "Request Body ", gjson.MustEncodeString(resp.RawResponse))
	}
	if err != nil {
		return err
	}

	return json.Unmarshal(resp.Body(), data)
}
