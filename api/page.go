package api

type PageRes struct {
	Total int64 `json:"total" dc:"总计记录"`
}
type PageReq struct {
	Page   int               `json:"page" dc:"页码"`
	Size   int               `json:"size" dc:"分页"`
	Sorter map[string]string `json:"sorter" dc:"排序规则"`
}

func (r *PageReq) Offset() int {
	if r.Page < 1 {
		r.Page = 1
	}
	if r.Size < 10 {
		r.Size = 10
	}
	return r.Size * (r.Page - 1)
}
