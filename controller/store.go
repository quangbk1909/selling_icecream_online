package controller

import (
	"log"
	"net/http"
	"vinid_project/model"

	"github.com/gin-gonic/gin"
)

type CoordinatesJSON struct {
	Latitude  float64 `json:"latitude" binding:"required"`
	Longitude float64 `json:"longitude" binding:"required"`
	Distance  float64 `json:"distance"`
}

// Lấy tất cả các cửa hàng
func (controller *Controller) GetStores(c *gin.Context) {
	var stores []model.Store
	err := controller.db.Find(&stores).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, stores)
}

// Lấy thông tin chi tiết 1 cửa hàng
func (controller *Controller) DetaiStore(c *gin.Context) {
	var store []model.Store
	err := controller.db.First(&store, c.Param("id")).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, store)
}

// Lấy các items trong 1 cửa hàng
func (controller *Controller) GetItemInStore(c *gin.Context) {
	var items []model.IceCreamItem
	idStore := c.Param("id")
	var store model.Store
	err := controller.db.First(&store, idStore).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Non store exist with id!"})
		return
	}

	controller.db.Model(&store).Related(&items, "IceCreamItems")
	c.JSON(http.StatusOK, items)
}

// Tìm cửa hàng xung quanh đây
func (controller *Controller) GetStoresAroundHere(c *gin.Context) {
	var coordinates CoordinatesJSON
	var stores []model.Store

	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if coordinates.Distance == 0 {
		coordinates.Distance = 0.02
	} else {
		coordinates.Distance = coordinates.Distance / 100
	}

	controller.db.Where("latitude > ? AND latitude < ? AND longitude > ? AND longitude < ?", coordinates.Latitude-coordinates.Distance, coordinates.Latitude+coordinates.Distance, coordinates.Longitude-coordinates.Distance, coordinates.Longitude+coordinates.Distance).Find(&stores)
	c.JSON(http.StatusOK, stores)
}

//Lấy hình ảnh của cửa hàng
func (controller *Controller) GetStoreImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/store_images/" + name)
}

// Search store by text
func (controller *Controller) SearchStore(c *gin.Context) {

}
