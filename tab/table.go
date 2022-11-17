package tab

import (
	"github.com/hinego/decimal"
)

type Model struct {
	ID        int64 `gorm:"primarykey" json:"id"`
	CreatedAt int64 `json:"-"`
	UpdatedAt int64 `json:"-"`
}
type Token struct {
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

type User struct {
	Model
	Name     string          `json:"name"`
	Email    string          `json:"email" gorm:"unique"`                   //用户邮箱
	Password string          `json:"password"`                              //用户密码
	Balance  decimal.Decimal `json:"balance" gorm:"type:numeric" dc:"用户余额"` //用户余额USD
	Refer    int64           `json:"refer"`                                 //邀请人ID
	IP       string          `json:"ip" dc:"注册IP"`
	Role     string          `json:"role"`
}
