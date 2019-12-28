package controller

import (
	"log"
	"net/http"
	"vinid_project/model"

	"github.com/gin-gonic/gin"
)

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

func (controller *Controller) GetStoreImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/store_images/" + name)
}
