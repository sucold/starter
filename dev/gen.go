package main

import (
	"context"
	"errors"
	"github.com/gogf/gf/v2/encoding/gjson"
	"github.com/gogf/gf/v2/os/gcmd"
	"github.com/gogf/gf/v2/os/gctx"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gproc"
	"github.com/gogf/gf/v2/text/gstr"
	"github.com/hinego/gen"
	"github.com/hinego/types"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"strings"
	"time"
)

var (
	g1        *gen.Generator
	genConfig *gen.CmdParams
	gdb       *gorm.DB
	dsn       = "host=192.168.32.130 user=postgres password=postgres dbname=status port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	typ       = "postgres"
	newLogger = logger.New(log.New(os.Stdout, "\r\n", log.Ltime|log.Lshortfile), logger.Config{
		SlowThreshold:             time.Second,
		LogLevel:                  logger.Silent,
		IgnoreRecordNotFoundError: true,
		Colorful:                  true,
	})
	Config = &gorm.Config{Logger: newLogger}
)

func prepare(ctx context.Context, parser *gcmd.Parser) {
	os.Remove("/etc/gen.db")
	log.SetFlags(log.Ltime | log.Lshortfile)
	genConfig = gen.ArgParse()
	if genConfig == nil {
		log.Fatalf("parse genConfig fail")
	}
	schema.RegisterSerializer("auto", types.AutoSerializer{})
	genConfig.DSN = parser.GetOpt("dsn", dsn).String()
	genConfig.DB = parser.GetOpt("db", typ).String()
	genConfig.OutPath = parser.GetOpt("outPath", "./app/dao").String()
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
		db.Logger = newLogger
		db.Logger.Info(ctx, "working directory: %s", gfile.Pwd())
		if tables, err := db.Migrator().GetTables(); err != nil {
			return
		} else {
			log.Println(gjson.MustEncodeString(tables))
			for _, table := range tables {
				if err := db.Migrator().DropTable(table); err != nil {
					log.Println("删除失败", err)
				}
			}
		}
		if tables, err := db.Migrator().GetTables(); err != nil {
			return
		} else {
			log.Println("清空后", len(tables))
		}
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
		gfile.Remove("app/dao")
		gfile.Remove("app/model")
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
			data := gfile.GetContents(v)
			data = gstr.Replace(data, `"github.com/gogf/gf/v2`, `"github.com/gogf/gf`)
			data = gstr.Replace(data, `"github.com/gogf/gf`, `"github.com/gogf/gf/v2`)
			gfile.PutContents(v, data)
		}
		print(gproc.MustShellExec(ctx, "git add ./app/dao"))
		print(gproc.MustShellExec(ctx, "git add ./app/model"))
		db, err := gdb.DB()
		if err == nil {
			db.Close()
			os.Remove("/etc/gen.db")
		}
		return nil
	},
}

func print(result string) {
	res := strings.Split(result, "\n")
	for _, v := range res {
		if !strings.Contains(v, "CRLF") && v != "" && !strings.Contains(v, "working directory") {
			log.Println(v)
		}
	}
}
func main() {
	gn.Run(gctx.New())
}
