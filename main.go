package main

import (
	"fmt"
	"github.com/WTBacon/todo-app-gin/routers"

	"github.com/WTBacon/todo-app-gin/config"
	"github.com/WTBacon/todo-app-gin/models"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var err error

func main() {
	config.DB, err = gorm.Open("mysql", config.DbUrl(config.NewDbConfig()))
	defer config.DB.Close()

	if err != nil {
		fmt.Println("Can't connect DB! Error:", err)
	}
	config.DB.AutoMigrate(&models.Todo{})

	r := routers.SetUpRouter()
	r.Run()
}
