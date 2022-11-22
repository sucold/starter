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
		//log.Println("删除原模块", gfile.Remove("./app/conset"))
		//log.Println("删除原模块", gfile.Remove(".git/modules/app"))
		//Bash("git submodule sync", "更新")
		//Bash("git submodule sync --recursive", "更新")
		//Bash("git submodule add https://github.com/hinego/conset   app/conset", "同步结果")
		//Bash("git submodule update --remote", "同步结果")
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
		if !gfile.Exists("./app/cmd/main.go") {
			err = gfile.Copy("./app/conset/command/main.go", "./app/cmd/main.go")
			log.Println("复制main.go", err)
			err = gfile.ReplaceFile("package command", "package cmd", "./app/cmd/main.go")
			log.Println("替换main.go", err)
			Bash("git add app/cmd/main.go", "git add")
			gfile.Remove("./app/cmd/main.re.go")
		}
		return nil
	},
}
