package cmd

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"os/exec"
	"strings"
)

func init() {
	AddCommand(syn)
}
func Bash(sh string, desc string) {
	output, err := exec.Command("cmd", "/C", sh).CombinedOutput()
	log.Println(desc, err, string(output))
}

var syn = &gcmd.Command{
	Name:  "sync",
	Usage: "sync",
	Brief: "[开发专用]同步更新子模块",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("go.mod") {
			return errors.New("此为开发使用命令")
		}
		log.Println("删除原模块", gfile.Remove("./app/conset"))
		Bash("git submodule update --remote", "同步结果")
		file, err := gfile.ScanDirFile("./app/conset", "*", true)
		if err != nil {
			return
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
