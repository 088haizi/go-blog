package main

import (
	"fmt"

	"github.com/088haizi/go-blog/config"
	"github.com/088haizi/go-blog/routes"
)

var err error

func main() {
	fmt.Println("Hello, world!")

	config.Migrate()

	config.CreateMysqlConn()
	defer config.DBClose()

	fmt.Printf("Hello, gorm!\n")

	r := routes.SetupRouter()

	r.Run()
}