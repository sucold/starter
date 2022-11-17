package ge

import (
	"context"
	_ "embed"
	"fmt"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"strings"
)

//go:embed api/api.go
var apiTemplate string

//go:embed api/controller.go
var controllerTemplate string

//go:embed api/logic.go
var logicTemplate string

type Action struct {
	*Model
	Action      string `json:"action"`
	ActionLower string `json:"action_lower"`
	Default     bool   `json:"default"`
	Path        string `json:"path"`
	Method      string `json:"method"`
	SM          string `json:"SM"`
}

type Model struct {
	Controller string             `json:"controller"`
	API        string             `json:"API"`
	Name       string             `json:"name"` //例如Product
	NameLower  string             `json:"nameLower"`
	Tags       string             `json:"tags"`
	Actions    map[string]*Action `json:"actions"`
}

func (r *Model) init() {
	var names = map[string]string{
		"get":    "获取",
		"fetch":  "列表",
		"create": "创建",
		"delete": "删除",
		"update": "更新",
	}
	r.NameLower = strings.ToLower(r.Name)
	for k, ac := range r.Actions {
		ac.Model = r
		ac.Action = k
		ac.ActionLower = strings.ToLower(ac.Action)
		if sm, ok := names[ac.ActionLower]; ok {
			ac.SM = sm + r.Tags
		} else {
			ac.SM = ac.ActionLower + r.Tags
			ac.Default = true
		}
		if ac.Method == "" {
			ac.Method = "get"
		}
		ac.Path = fmt.Sprintf("/%v/%v", r.NameLower, ac.ActionLower)
	}
}
func (r *Model) Map() map[string]any {
	return map[string]any{
		"Name":       r.Name,
		"NameLower":  r.NameLower,
		"Tags":       r.Tags,
		"Actions":    r.Actions,
		"Controller": r.Controller,
		"API":        r.API,
	}
}

type Group struct {
	Controller string `json:"controller"`
	API        string `json:"API"`
	Model      map[string]*Model
	Logic      map[string]*Model
}

type APP map[string]*Group

func (r *APP) Init() {
	for k, group := range *r {
		group.Controller = k
		for k1, model := range group.Model {
			model.Controller = group.Controller
			model.API = group.API
			model.Name = k1
			model.init()
		}
		for k1, model := range group.Logic {
			model.Controller = group.Controller
			model.API = group.API
			model.Name = k1
			model.init()
		}
	}
	r.generate()
}
func (r *APP) generate() {
	for _, group := range *r {
		for _, model := range group.Model {
			var data = model.Map()
			control, err := g.View().ParseContent(context.TODO(), controllerTemplate, data)
			conPath := fmt.Sprintf("./app/controller/%v/%v.go", model.Controller, model.NameLower)
			log.Println(err, gfile.PutContents(conPath, control))

			control, err = g.View().ParseContent(context.TODO(), apiTemplate, data)
			ApiPath := fmt.Sprintf("./api/%v/%v.go", model.API, model.NameLower)
			log.Println(err, gfile.PutContents(ApiPath, control))
		}
		for _, model := range group.Logic {
			var data = model.Map()
			control, err := g.View().ParseContent(context.TODO(), logicTemplate, data)
			ApiPath := fmt.Sprintf("./app/logic/%v.go", model.NameLower)
			log.Println(err, gfile.PutContents(ApiPath, control))
		}
	}
}
