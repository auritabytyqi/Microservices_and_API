package storage

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	_ "github.com/go-sql-driver/mysql"
)

var dsn = "auritab91:auritab91@tcp(127.0.0.1:3302)/msapis"
var DB, _ = gorm.Open(mysql.Open(dsn), &gorm.Config{})
