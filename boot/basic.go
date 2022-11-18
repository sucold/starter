package boot

import (
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/hinego/conset/base"
	"github.com/hinego/conset/cache"
	"github.com/hinego/conset/service/mail"
	"github.com/hinego/starter/app/dao"
	"github.com/hinego/starter/app/model"
)

func InitUser() error {
	var (
		u = dao.User
	)
	if count, err := u.Count(); err != nil {
		return err
	} else {
		if count != 0 {
			return nil
		}
	}
	data := []*model.User{
		{
			Email:    "admin@qq.com",
			Password: base.GeneratorPassword("AaGG12345678"),
			IP:       "127.0.0.1",
			Role:     "admin",
		},
	}
	return u.Create(data...)
}
func InitEmail(conf *mail.Config) func() error {
	return func() error {
		mailConfig := &mail.Config{
			Endpoint:          "dm.aliyuncs.com",
			AccessKeyId:       "LTAI5tKtV9PNgW2evEADaMte",
			AccessKeySecret:   "rVLeNZyYImp0nS9r93H45dVgltlP6l",
			AccountName:       "admin@goant.xyz",
			AddressType:       1,
			ReplyToAddress:    true,
			FromAlias:         base.AppName,
			ReplyAddress:      "admin@skyqq.cc",
			ReplyAddressAlias: base.AppName,
			Expire:            1800,
			ErrorTimes:        5,
			DiffTime:          60,
		}
		if conf == nil {
			return mail.Init(mailConfig)
		} else {
			return mail.Init(conf)
		}
	}
}
func InitRedis(conf *gredis.Config) func() error {
	return func() error {
		if redis, err := gredis.New(conf); err != nil {
			return err
		} else {
			cache.DefaultCache.SetAdapter(gcache.NewAdapterRedis(redis))
			return nil
		}
	}
}
