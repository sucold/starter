package table

import (
	"github.com/hinego/decimal"
	"github.com/sucold/starter/app/conset/tab"
)

type Token struct { //addition.go
	Model
	Code      string `json:"code" gorm:"unique"`
	Expire    int64  `json:"expire"`     //过期时间
	UserID    int64  `json:"user_id"`    //用户ID
	IP        string `json:"ip"`         //登录IP地址
	UserAgent string `json:"user_agent"` //登录时的userAgent
	LogoutAt  int64  `json:"logoutAt"`   //退出登录时间
	DeviceID  int64  `json:"device_id"`  //绑定的设备ID
	Data      string `json:"data"`       //数据内容
	Role      string `json:"role"`
}
type User struct { //addition.go
	Model
	Name     string          `json:"name"`
	Email    string          `json:"email" gorm:"unique"`                   //用户邮箱
	Password string          `json:"password"`                              //用户密码
	Balance  decimal.Decimal `json:"balance" gorm:"type:numeric" dc:"用户余额"` //用户余额USD
	Refer    int64           `json:"refer"`                                 //邀请人ID
	IP       string          `json:"ip" dc:"注册IP"`
	Role     string          `json:"role"`
}
type Service struct {
	Model
	Bind     string         `json:"bind" dc:"绑定的服务"`
	BindID   int64          `json:"bind_id" dc:"绑定的服务ID（主要指绑定的网络ID）"`
	Code     string         `json:"code" gorm:"unique"` //服务代码
	Name     string         `json:"name"`
	Desc     string         `json:"desc"`
	Status   string         `json:"status"`
	Text     string         `json:"text"`
	Auto     bool           `json:"auto" dc:"自动启动"`
	Data     string         `json:"data" dc:"服务的持久化数据"` //例如:同步区块号码
	BootTime int64          `json:"boot_time" dc:"启动时间"`
	StopTime int64          `json:"stop_time" dc:"停止时间"`
	Form     tab.Form       `json:"form" dc:"表单" gorm:"-"`
	Sort     int64          `json:"sort" dc:"排序"`
	Link     map[string]any `json:"link" dc:"链接" gorm:"serializer:auto;type:json" filed:"bytes"`
}

type Log struct {
	Model
	UserID  int64  `json:"user_id"`
	Type    string `json:"type" dc:"日志类型"`
	LinkID  int64  `json:"link_id"`
	Name    string `json:"name"`
	Content string `json:"content"`
}
