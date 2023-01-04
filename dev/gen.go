package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hinego/gen"
	"gorm.io/gorm"
	"log"
	"os"
)

var g1 *gen.Generator
var genConfig *gen.CmdParams
var gdb *gorm.DB

func prepare(ctx context.Context, parser *gcmd.Parser) {
	os.Remove("/etc/gen.db")
	log.SetFlags(log.Ldate | log.Ltime)
	genConfig = gen.ArgParse()
	if genConfig == nil {
		log.Fatalf("parse genConfig fail")
	}
	genConfig.DSN = parser.GetOpt("dsn", "/etc/gen.db").String()
	genConfig.DB = parser.GetOpt("db", "sqlite").String()
	genConfig.OutPath = parser.GetOpt("outPath", "./app/dao").String()
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
		gdb = db
	}
}

var gn = &gcmd.Command{
	Name:  "main",
	Usage: "main",
	Brief: "[开发专用]生成dao",
	Func: func(ctx context.Context, parser *gcmd.Parser) (err error) {
		if !gfile.Exists("go.mod") {
			return errors.New("此为开发使用命令")
		}
		prepare(ctx, parser)
		g1.LinkModel(Md...)
		if !genConfig.OnlyModel {
			g1.ApplyBasic(g1.GenerateAllTable()...)
		}
		g1.Execute()
		file, err := gfile.ScanDirFile("app/model", "*", true)
		if err != nil {
			log.Println(err)
			return nil
		}
		for _, v := range file {
			log.Println(v)
			data := gfile.GetContents(v)
			data = gstr.Replace(data, `"github.com/gogf/gf/v2`, `"github.com/gogf/gf`)
			data = gstr.Replace(data, `"github.com/gogf/gf`, `"github.com/gogf/gf/v2`)
			gfile.PutContents(v, data)
		}
		gproc.MustShellRun(ctx, "git add ./app/dao")
		gproc.MustShellRun(ctx, "git add ./app/model")
		db, err := gdb.DB()
		if err == nil {
			db.Close()
			os.Remove("/etc/gen.db")
		}
		return nil
	},
}

func main() {
	gn.Run(gctx.New())
}
