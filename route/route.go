package route

import (
	"vinid_project/authentication"
	"vinid_project/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, c *controller.Controller) {

	r.POST("/register", c.Register)
	r.POST("/login", c.Login)

	userR := r.Group("/users")
	{
		userR.GET("", c.GetUsers)
		userR.GET("/:id", c.DetailUser)
		userR.GET("/:id/orders", c.GetOrderOfUser)
		userR.PUT("/:id", c.UpdateInfo)

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

	orderR := r.Group("/orders")
	{
		//orderR.GET("", c.GetItems)
		orderR.GET("/:id", c.DetaiOrder)
		orderR.POST("", c.CreatOrder)

	}

	searchR := r.Group("/search")
	{
		searchR.GET("/item", c.SearchItem)
		searchR.GET("/store", c.SearchStore)
	}

	r.GET("/ping", authentication.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, c.GetInt("userID"))
	})

	r.GET("/resources/store_images/:name", c.GetStoreImage)
	r.GET("/resources/item_images/:name", c.GetItemImage)
}
