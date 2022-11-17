package main

import (
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/hinego/starter/internal/cmd"
	_ "github.com/hinego/starter/internal/packed"
	"log"
)

func main() {
	//test
	log.SetFlags(log.Llongfile)
	cmd.Main.Run(gctx.New())
}
