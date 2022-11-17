package base

import (
	"bytes"
	"golang.org/x/crypto/bcrypt"
)

const (
	PassSalt = "ASDGVRgrege258OP:PHGNHG<"
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
