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

	c := controller.NewController(db)

	r.GET("/ping", c.TestController)
	userR := r.Group("/users")
	{
		userR.GET("", c.GetUsers)
		userR.GET("/:id", c.DetailUser)

	}

	storeR := r.Group("/stores")
	{
		storeR.GET("", c.GetStores)
		storeR.GET("/:id", c.DetaiStore)
		storeR.GET("/:id/items", c.GetItemInStore)

	}

	itemR := r.Group("/items")
	{
		itemR.GET("", c.GetItems)
		itemR.GET("/:id", c.DetaiItem)

	}

	r.GET("/resources/store_images/:name", c.GetStoreImage)
	r.GET("/resources/item_images/:name", c.GetItemImage)
	//r.StaticFile("resources/store_images/1.png", "./resources/store_images/1.png")
	r.Run(":8080")
}
