package config

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"

	"github.com/088haizi/go-blog/database"
)

var db *gorm.DB

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
		dbName:   "go_blog",
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

func CreateMysqlConn() *gorm.DB {
	var err error

	db, err = gorm.Open("mysql", DBUrl(BuildDBConfig()))
	if err != nil {
		fmt.Println("status: ", err)
	}
	fmt.Println("create mysql database connection successful")

	return db
}

func DBConn() *gorm.DB {
	return db
}

func DBClose() {
	db.Close()
}

func Migrate() {
	CreateMysqlConn()
	defer DBClose()
	db.Set("gorm:table_options", "ENGINE=InnoDB").AutoMigrate(&database.Post{})
	//db.Model(&database.Post{}).AddIndex("created_at")
}