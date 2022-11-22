//go:build mysql

package database

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func GetDial(dsn string) gorm.Dialector { //sqlite
	return mysql.Open(dsn)
}
