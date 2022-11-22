package base

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
)

var (
	PassSalt = "ASDGVRgrege258OP:PHGNHG<"
	AppName  = "全球网络加速服务"
	UserKey  = "JWT_UserKey"
)

func Salt(password string) []byte {
	var buffer bytes.Buffer
	buffer.WriteString(password)
	buffer.WriteString("__|__")
	buffer.WriteString(PassSalt)
	return buffer.Bytes()
}
func GeneratorPassword(password string) string {
	if pass, err := bcrypt.GenerateFromPassword(Salt(password), bcrypt.DefaultCost); err != nil {
		return ""
	} else {
		return string(pass)
	}
}
