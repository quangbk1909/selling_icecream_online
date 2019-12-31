package route

import (
	"vinid_project/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, c *controller.Controller) {
	userR := r.Group("/users")
	{
		userR.GET("", c.GetUsers)
		userR.GET("/:id", c.DetailUser)
		//userR.POST("/:id/deposit".c.Deposit)

	}

	storeR := r.Group("/stores")
	{
		storeR.GET("", c.GetStores)
		storeR.GET("/:id", c.DetaiStore)
		storeR.GET("/:id/items", c.GetItemInStore)
	}

	r.GET("/around_here", c.GetStoresAroundHere)

	itemR := r.Group("/items")
	{
		itemR.GET("", c.GetItems)
		itemR.GET("/:id", c.DetaiItem)
	}

	searchR := r.Group("/search")
	{
		searchR.GET("/item", c.SearchItem)
		searchR.GET("/store", c.SearchStore)
	}

	r.GET("ping", c.TestController)

	r.GET("/resources/store_images/:name", c.GetStoreImage)
	r.GET("/resources/item_images/:name", c.GetItemImage)
}
