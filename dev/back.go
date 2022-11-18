package main

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
)

func init() {
	AddCommand(bc)
}

var bc = &gcmd.Command{
	Name:  "cc",
	Usage: "cc",
	Brief: "[开发专用]恢复默认",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		err = gfile.Remove("./app/cmd/main.go")
		if err != nil {
			return err
		}
		gfile.PutContents("./app/cmd/main.re.go", `package cmd

import "github.com/gogf/gf/v2/os/gcmd"

var Main = &gcmd.Command{}
`)
		return nil
	},
}
