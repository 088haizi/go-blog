package main

import (
	"fmt"

	"github.com/088haizi/go-blog/config"
)

var err error

func main() {
	fmt.Println("Hello, world!")

	config.CreateMysqlConn()
	config.Migrate()
	defer config.DBClose()

	fmt.Printf("Hello, gorm!\n")
}