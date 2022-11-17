package {{.API}}

import (
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/starter/api"
	"github.com/hinego/starter/internal/model"
)
{{range $k, $v := .Actions}}
type {{$v.Name}}{{$v.Action}}Req struct {
	g.Meta `path:"{{$v.Path}}" method:"{{$v.Method}}" sm:"{{$v.SM}}" tags:"{{$v.Tags}}"`
	{{if eq $v.Action "Fetch"}}types.PageReq{{else if eq $v.Action "Create"}}*model.{{$v.Name}}{{else if eq $v.Action "Update"}}*model.{{$v.Name}}{{else if eq $v.Action "Delete"}}ID     int64 `json:"id"`{{else if eq $v.Action "Get"}}ID     int64 `json:"id"`{{end}}
}
{{if  eq $v.Action "Get"}}type {{$v.Name}}{{$v.Action}}Res model.{{$v.Name}}
{{else if eq $v.Action "Fetch"}}type {{$v.Name}}{{$v.Action}}Res struct {
	{{if eq $v.Action "Fetch"}}*types.PageReq
	*types.PageRes
	Data []*model.{{$v.Name}} `json:"data"`{{end}}
}
{{else}}type {{$v.Name}}{{$v.Action}}Res struct{}
{{end}}{{end}}