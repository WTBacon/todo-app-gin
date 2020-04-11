package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
)

var DB *gorm.DB

type dbConfig struct {
	host     string
	port     int
	user     string
	dbName   string
	password string
}

func NewDbConfig() *dbConfig {
	dbConfig := dbConfig{
		host:     "localhost",
		port:     3306,
		user:     "root",
		dbName:   "todo",
		password: "password",
	}
	return &dbConfig
}

func DbUrl(dbConfig *dbConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.user,
		dbConfig.password,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbName,
	)
}
