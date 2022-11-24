package main

import (
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"log"
)

var Main = gcmd.Command{}

func AddCommand(command *gcmd.Command) {
	_ = Main.AddCommand(command)
}
func main() {
	log.SetFlags(log.Llongfile)
	Main.Run(gctx.New())
}
