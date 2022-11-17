package logic

type {{.NameLower}}Logic struct{}

var (
	{{.Name}} = {{.NameLower}}Logic{}
)

func (r *{{.NameLower}}Logic) test() {
	return
}