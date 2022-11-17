package main

import (
	"github.com/hinego/systemd/internal/database"
	"github.com/hinego/systemd/internal/table"
	"gorm.io/gen"
	"gorm.io/gorm/schema"
	"log"
	"os"
)

var g *gen.Generator
var config *gen.CmdParams

func init() {
	os.Remove("gorm.db")
	log.SetFlags(log.Llongfile)
	config = gen.ArgParse()
	if config == nil {
		log.Fatalf("parse config fail")
	}
	schema.RegisterSerializer("auto", database.AutoSerializer{})
	if db, err := gen.Connect(gen.DBType(config.DB), config.DSN); err != nil {
		log.Fatalf("connect db server fail: %v", err)
	} else {
		g = gen.NewGenerator(gen.Config{
			Mode:              gen.WithDefaultQuery | gen.WithoutContext,
			OutPath:           config.OutPath,
			OutFile:           config.OutFile,
			ModelPkgPath:      config.ModelPkgName,
			WithUnitTest:      config.WithUnitTest,
			FieldNullable:     false,
			FieldWithIndexTag: true,
			FieldWithTypeTag:  true,
			FieldSignable:     config.FieldSignable,
		})
		g.UseDB(db)
	}
}

func main() {
	m := []any{
		table.Token{},
		table.User{},
		table.Bin{},
		table.File{},
		table.Country{},
		table.Transform{},
	}
	g.LinkModel(m...)
	if !config.OnlyModel {
		g.ApplyBasic(g.GenerateAllTable()...)
	}
	g.Execute()
}
