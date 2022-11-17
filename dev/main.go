package main

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
	"os/exec"
)

var Main = gcmd.Command{}

func Bash(sh string, desc string) {
	output, err := exec.Command("cmd", "/C", sh).CombinedOutput()
	log.Println(desc, err, string(output))
}

func AddCommand(command *gcmd.Command) {
	_ = Main.AddCommand(command)
}
func main() {
	log.SetFlags(log.Llongfile)
	Main.Run(gctx.New())
}
