package cmd

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gcmd"
)

func init() {
	AddCommand(main)
}

var main = &gcmd.Command{
	Name:  "test",
	Usage: "test",
	Brief: "start test server",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		s := g.Server()
		s.EnableAdmin()
		s.Run()
		return nil
	},
}
