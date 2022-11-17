package main

import (
	"context"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"strings"
)

func init() {
	AddCommand(syn)
}

var syn = &gcmd.Command{
	Name:  "sync",
	Usage: "sync",
	Brief: "[开发专用]同步模块",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		log.Println("删除原模块", gfile.Remove("./app/conset"))
		Bash("git submodule update --remote", "同步结果")
		file, err := gfile.ScanDirFile("./app/conset", "*", true)
		if err != nil {
			return nil
		}
		for _, f := range file {
			content := gfile.GetContents(f)
			if strings.Contains(content, "github.com/hinego/conset") {
				err = gfile.PutContents(f, strings.ReplaceAll(content, "github.com/hinego/conset", "github.com/hinego/starter/app/conset"))
				log.Println(f, err)
			}
		}
		return nil
	},
}
