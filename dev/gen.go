package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/hinego/gen"
	"github.com/hinego/starter/dev/table"
	"log"
	"os"
)

func init() {
	AddCommand(gn)
}

var g1 *gen.Generator
var genConfig *gen.CmdParams

func prepare() {
	os.Remove("gorm.db")
	log.SetFlags(log.Llongfile)
	genConfig = gen.ArgParse()
	if genConfig == nil {
		log.Fatalf("parse genConfig fail")
	}
	//schema.RegisterSerializer("auto", database.AutoSerializer{})
	if db, err := gen.Connect(gen.DBType(genConfig.DB), genConfig.DSN); err != nil {
		log.Fatalf("connect db server fail: %v", err)
	} else {
		g1 = gen.NewGenerator(gen.Config{
			Mode:              gen.WithDefaultQuery | gen.WithoutContext,
			OutPath:           genConfig.OutPath,
			OutFile:           genConfig.OutFile,
			ModelPkgPath:      genConfig.ModelPkgName,
			WithUnitTest:      genConfig.WithUnitTest,
			FieldNullable:     false,
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			FieldSignable:     genConfig.FieldSignable,
		})
		g1.UseDB(db)
	}
}

var gn = &gcmd.Command{
	Name:  "gen",
	Usage: "gen",
	Brief: "[开发专用]生成dao",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("go.mod") {
			return errors.New("此为开发使用命令")
		}
		prepare()
		m := []any{
			table.Token{},
			table.User{},
		}
		g1.LinkModel(m...)
		if !genConfig.OnlyModel {
			g1.ApplyBasic(g1.GenerateAllTable()...)
		}
		g1.Execute()
		os.Remove("gorm.db")
		return nil
	},
}
