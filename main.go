package main

import (
	"vinid_project/controller"
	"vinid_project/database"
	"vinid_project/route"

	"github.com/gin-gonic/gin"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := gin.Default()

	db := database.ConnectDB()
	db.SingularTable(true)	
	defer db.Close()

	dao := database.NewDao(db)

	c := controller.NewController(dao)

	route.InitRoute(r, c)

	r.Run(":8080")
}
