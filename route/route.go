package route

import (
	"vinid_project/authentication"
	"vinid_project/controller"

	"github.com/gin-gonic/gin"
)

func InitRoute(r *gin.Engine, c *controller.Controller) {

	r.GET("/", func(c *gin.Context) {
		c.JSON(200, "Helu! Đà đa đa đa đá ... sâu coollll!")
	})

	r.POST("/register", c.Register)
	r.POST("/login", c.Login)

	apiR := r.Group("/api")
	apiR.Use(authentication.AuthMiddleware())
	{
		userR := apiR.Group("/users")
		{
			userR.GET("", c.GetUsers)
			userR.GET("/:id", c.DetailUser)
			userR.GET("/:id/orders", c.GetOrderOfUser)
			userR.PUT("", c.UpdateInfo)
			userR.PUT("/deposit", c.Deposite)

		}
		storeR := apiR.Group("/stores")
		{
			storeR.GET("", c.GetStores)
			storeR.GET("/:id", c.DetaiStore)
			storeR.GET("/:id/items", c.GetItemInStore)
		}

		apiR.GET("/around_here", c.GetStoresAroundHere)

		itemR := apiR.Group("/items")
		{
			itemR.GET("", c.GetItems)
			itemR.GET("/:id", c.DetaiItem)
			itemR.GET("/:id/ratings", c.GetRatingsOfItem)

		}

		ratingR := apiR.Group("/ratings")
		{
			ratingR.POST("", c.CreateRating)
		}

		orderR := apiR.Group("/orders")
		{
			//orderR.GET("", c.GetItems)
			orderR.GET("/:id", c.DetaiOrder)
			orderR.POST("", c.CreateOrder)
		}

		searchR := apiR.Group("/search")
		{
			searchR.GET("/item", c.SearchItem)
			searchR.GET("/store", c.SearchStore)
		}

	}

	r.GET("/resources/store_images/:name", c.GetStoreImage)
	r.GET("/resources/item_images/:name", c.GetItemImage)

	r.GET("/ping", authentication.AuthMiddleware(), func(c *gin.Context) {
		c.JSON(200, c.GetInt("userID"))
	})

	r.GET("test_geocoder", c.TestGeoCoder)

}
