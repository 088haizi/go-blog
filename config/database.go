package config

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var DB *gorm.DB

type DBConfig struct {
	host		string
	port		int
	user		string
	dbName		string
	password	string
}

func BuildDBConfig() *DBConfig {
	dbConfig := DBConfig{
		host:     "0.0.0.0",
		port:     3306,
		user:     "root",
		dbName:   "todos",
		password: "hiamroot",
	}
	return &dbConfig
}

func DBUrl(dbConfig *DBConfig) string {
	return fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		dbConfig.user,
		dbConfig.password,
		dbConfig.host,
		dbConfig.port,
		dbConfig.dbName,
	)
}