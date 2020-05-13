package main

import (
	"fmt"
	"github.com/jinzhu/gorm"

	"github.com/088haizi/go-blog/config"
)

var err error

func main() {
	fmt.Println("Hello, world!")

	config.DB, err = gorm.Open("mysql", config.DBUrl(config.BuildDBConfig()))
	defer config.DB.Close()

	if err != nil {
		fmt.Println("status: ", err)
	}

	config.DB.AutoMigrate()

	fmt.Printf("Hello, gorm!\n")
}