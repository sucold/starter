package {{.Controller}}

import (
	"context"
	"github.com/hinego/errorx"
	{{.API}} "github.com/sucold/starter/api/{{.API}}"
	"github.com/sucold/starter/app/dao"
	"github.com/sucold/starter/app/model"
	"github.com/hinego/types"
)

type {{.NameLower}}Controller struct{}

var (
	{{.Name}} = {{.NameLower}}Controller{}
)
{{range $k, $v := .Actions}}
func (c *{{$v.NameLower}}Controller) {{$v.Action}}(ctx context.Context, req *{{$v.API}}.{{$v.Name}}{{$v.Action}}Req) (res *{{$v.API}}.{{$v.Name}}{{$v.Action}}Res, err error) {
	{{if $v.Default}}return{{else}}var (
		u = dao.{{$v.Name}}
        {{if eq $v.Action "Get"}}get *model.{{$v.Name}}
		{{end}}//r  = g.RequestFromCtx(ctx)
		//id = r.GetParam(consts.UserKey).Int64()
	){{if eq $v.Action "Create"}}
	req.ID = 0
	if err = u.Create(req.{{$v.Name}}); err != nil {
		return nil, err
	}
    return nil, errorx.NewCode(0, "创建成功", nil){{else if eq $v.Action "Delete"}}
	if _, err = u.Where(u.ID.Eq(req.ID)).Delete(); err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "删除成功", nil){{else if eq $v.Action "Update"}}
	_, err = u.Where(u.ID).Omit(u.ID).Updates(req.{{$v.Name}})
	if err != nil {
		return nil, err
	}
    return nil, errorx.NewCode(0, "更新成功", nil){{else if eq $v.Action "Fetch"}}
    res = &{{$v.API}}.{{$v.Name}}{{$v.Action}}Res{
		PageReq: &req.PageReq,
		PageRes: &types.PageRes{},
		Data:    make([]*model.{{$v.Name}}, 0),
	}
    if res.Total, err = u.Order(u.ID).ScanByPage(&res.Data, res.Offset(), res.Size); err != nil {
		return nil, err
	}
    return{{else if eq $v.Action "Get"}}
    if get, err = u.Where(u.ID.Eq(req.ID)).First(); err != nil {
		return nil, err
	}
	return nil, errorx.NewCode(0, "success", get){{end}}{{end}}
}{{end}}