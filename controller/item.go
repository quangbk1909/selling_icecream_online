package controller

import (
	"net/http"
	"vinid_project/model"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) GetItems(c *gin.Context) {
	var items []model.IceCreamItem
	err := controller.db.Find(&items).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

func (controller *Controller) DetaiItem(c *gin.Context) {
	var item []model.IceCreamItem
	err := controller.db.First(&item, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

func (controller *Controller) GetItemImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/item_images/" + name)
}
