package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hinego/starter/app/cmd"
	_ "github.com/hinego/starter/app/packed"
	"log"
)

func main() {
	//test
	log.SetFlags(log.Llongfile)
	cmd.Main.Run(gctx.New())
}
