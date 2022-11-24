package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/sucold/starter/dev/fun"
)

var Main = &gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "[开发专用]更新umiset",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("package.json") {
			return errors.New("请在项目根目录执行")
		}
		cmd := `git clone https://github.com/sucold/starter-ui.git /projects/git/uminew`
		fun.Bash(cmd, "clone 到本地")
		fun.Bash("git -C /projects/git/uminew reset --hard", "重置")
		fun.Bash("git -C /projects/git/uminew fetch --all", "获取最新版本")
		fun.Bash("git -C /projects/git/uminew pull", "合并到最新版本")
		fun.Bash("git rm -rf src/umiset", "删除旧文件")
		err = gfile.CopyDir("/projects/git/uminew/src/umiset", "src/umiset")
		if err != nil {
			return errors.New("复制文件失败" + err.Error())
		}
		fun.Bash("git add src/umiset", "git add")
		fun.Bash("git -C /projects/git/uminew log -1 --format=%cd", "显示最新版本")
		return nil
	},
}

func main() {
	Main.Run(gctx.New())
}
