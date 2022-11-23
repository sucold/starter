package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/sucold/starter/app/cmd"
	_ "github.com/sucold/starter/app/packed"
	"log"
)

func main() {
	log.SetFlags(log.Llongfile)
	cmd.Main.Run(gctx.New())
}
