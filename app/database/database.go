package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
	"gorm.io/gorm/schema"
	"log"
	"os"
	"time"
)

var DB *gorm.DB

func Initial(dsn string) error {
	schema.RegisterSerializer("auto", AutoSerializer{})
	newLogger := logger.New(
		log.New(os.Stdout, "\r\n", log.LstdFlags), // io writer（日志输出的目标，前缀和日志包含的内容——译者注）
		logger.Config{
			SlowThreshold:             time.Second,   // 慢 SQL 阈值
			LogLevel:                  logger.Silent, // 日志级别
			IgnoreRecordNotFoundError: true,          // 忽略ErrRecordNotFound（记录未找到）错误
			Colorful:                  true,          // 禁用彩色打印
		},
	)
	var err error
	var db = new(gorm.DB)
	db, err = gorm.Open(sqlite.Open(`/etc/bot.db`), &gorm.Config{Logger: newLogger})
	if err != nil {
		return err
	}
	//log.Println("db.AutoMigrate", db.AutoMigrate(model.Token{}, model.User{}, model.Bin{}, model.File{}, model.Country{}, model.Transform{}))
	DB = db
	return err
}
