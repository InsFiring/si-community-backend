package config

import (
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

func DBConnection() (*gorm.DB, error) {
	// connection := "test:test1234@tcp(127.0.0.1:3306)/si_community?charset=utf8mb4&parseTime=True&loc=Local"
	connection := "test:test1234@tcp(127.0.0.1:13306)/si_community?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})
	return db, err
}
