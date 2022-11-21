package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"log"
	"os/exec"
)

func Bash(sh string, desc string) {
	output, err := exec.Command("cmd", "/C", sh).CombinedOutput()
	log.Println(desc, err, string(output))
}

var Main = &gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "[开发专用]更新umiset",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("package.json") {
			return errors.New("请在项目根目录执行")
		}
		cmd := `git clone https://github.com/hinego/uminew.git /projects/git/uminew`
		Bash(cmd, "clone 到本地")
		Bash("git -C /projects/git/uminew reset --hard", "重置")
		Bash("git -C /projects/git/uminew fetch --all", "获取最新版本")
		Bash("git -C /projects/git/uminew pull", "合并到最新版本")
		Bash("git rm -rf src/umiset", "删除旧文件")
		err = gfile.CopyDir("/projects/git/uminew/src/umiset", "src/umiset")
		if err != nil {
			return errors.New("复制文件失败" + err.Error())
		}
		Bash("git add src/umiset", "git add")
		Bash("git -C /projects/git/uminew log -1 --format=%cd", "显示最新版本")
		return nil
	},
}

func main() {
	Main.Run(gctx.New())
}
