//go:build !postgres && !mysql

package database

import (
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

func GetDial(dsn string) gorm.Dialector { //sqlite
	return sqlite.Open(dsn)
}
