package controller

import (
	"net/http"
	"strconv"
	"vinid_project/database"
	"vinid_project/model"
	"vinid_project/utility"

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
	var storeDao database.StoreDao = controller.dao

	stores, err := storeDao.FetchStore()
	if err != nil {
		message := "Internal server error. Database error!"
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, message, nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", stores))
}

// Lấy thông tin chi tiết 1 cửa hàng
func (controller *Controller) DetaiStore(c *gin.Context) {
	var dataResponse map[string]interface{}
	var store model.Store
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	var storeDao database.StoreDao = controller.dao
	store, err = storeDao.GetStoreByID(id)

	if err != nil {
		message := "Non store exist with id!"
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(404, message, nil))
		return
	}
	dataResponse = map[string]interface{}{
		"id":         store.ID,
		"name":       store.Name,
		"address":    store.Address,
		"image_path": store.ImagePath,
		"latitude":   store.Latitude,
		"longitude":  store.Longitude,
		"created_at": store.CreatedAt,
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", dataResponse))
}

// Lấy các items trong 1 cửa hàng
func (controller *Controller) GetItemInStore(c *gin.Context) {
	var items []model.IceCreamItem
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	var storeDao database.StoreDao = controller.dao

	items, err = storeDao.GetItemInStore(id)
	if err != nil {
		message := "Non store exist with id!"
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, message, nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", items))
}

// Tìm cửa hàng xung quanh đây
func (controller *Controller) GetStoresAroundHere(c *gin.Context) {
	var coordinates CoordinatesJSON
	var stores []model.Store

	if err := c.ShouldBindJSON(&coordinates); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request! Not enough info for request.", nil))
		return
	}

	if coordinates.Distance == 0 {
		coordinates.Distance = 0.02
	} else {
		coordinates.Distance = coordinates.Distance / 110
	}

	var storeDao database.StoreDao = controller.dao
	stores, err := storeDao.GetStoreAroundHere(coordinates.Latitude, coordinates.Longitude, coordinates.Distance)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", stores))
}

//Lấy hình ảnh của cửa hàng
func (controller *Controller) GetStoreImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/store_images/" + name)
}

// Search store by text
func (controller *Controller) SearchStore(c *gin.Context) {

}
