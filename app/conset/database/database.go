package database

import (
	"context"
	"github.com/gogf/gf/v2/frame/g"
	"github.com/hinego/starter/app/dao"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var newLogger = logger.New(
	log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
	logger.Config{
		SlowThreshold:             time.Second,   // 慢 SQL 阈值
		LogLevel:                  logger.Silent, // 日志级别
		IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
		Colorful:                  true,          // 禁用彩色打印
	},
)
var Config = &gorm.Config{Logger: newLogger}

func Init(mod ...any) func() error {
	return func() error {
		schema.RegisterSerializer("auto", AutoSerializer{})
		var err error
		var db = new(gorm.DB)
		dsn := g.Cfg().MustGet(context.TODO(), "gorm").String()
		db, err = gorm.Open(GetDial(dsn), Config)
		if err != nil {
			return err
		}
		if err = db.AutoMigrate(mod...); err != nil {
			return err
		}
		dao.SetDefault(db)
		return err
	}
}
