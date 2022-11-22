package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/sucold/starter/dev/ge"
)

func init() {
	AddCommand(gm)
}

var gm = &gcmd.Command{
	Name:  "gcc",
	Usage: "gcc",
	Brief: "[开发专用]生成model",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("go.mod") {
			return errors.New("此为开发使用命令")
		}
		var data = ge.APP{
			"test": {
				API: "v2",
				Model: map[string]*ge.Model{
					"User": {
						Tags: "用户",
						Actions: map[string]*ge.Action{
							"Fetch":  {},
							"Delete": {},
							"Create": {},
							"Update": {},
							"Get":    {},
							"Test":   {},
						},
					},
				},
				Logic: map[string]*ge.Model{
					"User": {},
				},
			},
		}
		data.Init()
		return nil
	},
}
