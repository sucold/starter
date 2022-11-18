package cmd

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hinego/starter/app/cache"
	"github.com/hinego/starter/app/dao"
	"github.com/hinego/starter/app/database"
)

type configCommand struct{}

var (
	config = configCommand{}
)

func (r *configCommand) initDatabase() error {
	dsn := "host=database.local user=postgres password=postgres dbname=sock port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	if err := database.Initial(dsn); err != nil {
		return err
	}
	dao.SetDefault(database.DB)
	return nil
}
func (r *configCommand) initRedis() error {
	redisConfig := &gredis.Config{
		Address: "192.168.32.130:6379",
		Db:      0,
		Pass:    "skyqqcc",
	}
	if redis, err := gredis.New(redisConfig); err != nil {
		return err
	} else {
		cache.DefaultCache.SetAdapter(gcache.NewAdapterRedis(redis))
		return nil
	}
}
