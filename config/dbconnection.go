package config

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type DBConfig struct {
	Server   string
	Port     int
	Dbname   string
	User     string
	Password string
}

func DBConnection(tomlConfig TomlConfig) (*gorm.DB, error) {
	connection := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		tomlConfig.Database.User, tomlConfig.Database.Password, tomlConfig.Database.Server, tomlConfig.Database.Port, tomlConfig.Database.Dbname)
	db, err := gorm.Open(mysql.Open(connection), &gorm.Config{})

	return db, err
}
