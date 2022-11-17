package cmd

import (
	"bytes"
	"crypto/tls"
	"github.com/gogf/gf/v2/database/gredis"
	"github.com/gogf/gf/v2/net/ghttp"
	"github.com/gogf/gf/v2/os/gcache"
	"github.com/gogf/gf/v2/os/gfile"
	"github.com/gogf/gf/v2/os/gres"
	"github.com/hinego/systemd/internal/cache"
	"github.com/hinego/systemd/internal/consts"
	"github.com/hinego/systemd/internal/dao"
	"github.com/hinego/systemd/internal/database"
	"github.com/hinego/systemd/internal/model"
	"github.com/hinego/systemd/internal/service/mail"
	"golang.org/x/crypto/bcrypt"
	"strings"
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
func (r *configCommand) Proxy(request *ghttp.Request) {
	request.Response.Status = 200
	file := gres.Get("public/index.html")
	if file != nil {
		request.Response.Write(file.Content())
	} else if data := gfile.GetBytes("public/index.html"); data != nil {
		request.Response.Write(data)
	} else {
		request.Response.Status = 404
	}
}
func (r *configCommand) initEmail() error {
	mailConfig := &mail.Config{
		Endpoint:          "dm.aliyuncs.com",
		AccessKeyId:       "LTAI5tKtV9PNgW2evEADaMte",
		AccessKeySecret:   "rVLeNZyYImp0nS9r93H45dVgltlP6l",
		AccountName:       "admin@goant.xyz",
		AddressType:       1,
		ReplyToAddress:    true,
		FromAlias:         consts.AppName,
		ReplyAddress:      "admin@skyqq.cc",
		ReplyAddressAlias: consts.AppName,
		Expire:            1800,
		ErrorTimes:        5,
		DiffTime:          60,
	}
	return mail.Init(mailConfig)
}
func (r *configCommand) initHosts() error {
	var (
		path = "/etc/hosts"
		host = "127.0.0.1 " + consts.ServerName
	)
	hosts := gfile.GetContents(path)
	if !strings.Contains(hosts, host) {
		return gfile.PutContentsAppend(path, "\n"+host+"\n")
	}
	return nil
}
func (r *configCommand) initUser() error {
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
			ID:       1,
			Email:    "admin@qq.com",
			Password: r.salt("AaGG12345678"),
			IP:       "127.0.0.1",
			Role:     "admin",
		},
		{
			Email:    "12345@qq.com",
			Password: r.salt("s51234512348"),
			IP:       "127.0.0.1",
			Role:     "user",
			ID:       2,
		},
	}
	return u.Create(data...)
}
func (r *configCommand) salt(password string) string {
	var buffer bytes.Buffer
	buffer.WriteString(password)
	buffer.WriteString("__|__")
	buffer.WriteString(consts.PassSalt)
	pass, _ := bcrypt.GenerateFromPassword(buffer.Bytes(), bcrypt.DefaultCost)
	return string(pass)
}
func (r *configCommand) TLSConfig() (*tls.Config, error) {
	pair, err := tls.X509KeyPair([]byte(consts.ServerCert), []byte(consts.ServerKey))
	if err != nil {
		return nil, err
	}
	return &tls.Config{
		Certificates: []tls.Certificate{pair},
	}, nil
}
