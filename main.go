package main

import (
	"vinid_project/controller"
	"vinid_project/database"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()
	db := database.ConnectDB()
	db.SingularTable(true)
	defer db.Close()
	conn := controller.NewController(db)
	r.GET("/ping", conn.TestController)
	r.Run(":8080")
}
