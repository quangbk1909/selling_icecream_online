package controller

import (
	"net/http"
	"vinid_project/model"
	"vinid_project/utility"

	"github.com/gin-gonic/gin"
)

//Lấy danh sách tất cả các sản phẩm
func (controller *Controller) GetItems(c *gin.Context) {
	var items []model.IceCreamItem
	err := controller.db.Find(&items).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, items)
}

// Lấy chi tiết một sản phẩm
func (controller *Controller) DetaiItem(c *gin.Context) {
	var item []model.IceCreamItem
	err := controller.db.First(&item, c.Param("id")).Error
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, item)
}

// Lấy hình ảnh của 1 item
func (controller *Controller) GetItemImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/item_images/" + name)
}

// Search item theo text
func (controller *Controller) SearchItem(c *gin.Context) {
	var items []model.IceCreamItem
	textSearch := c.Query("text")
	if textSearch == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No text search"})
	} else {
		textSearch = utility.StringSearchText(textSearch)
		query := "SELECT * FROM ice_cream_item WHERE MATCH (name,type) AGAINST ('" + textSearch + "' IN BOOLEAN MODE);"
		err := controller.db.Raw(query).Scan(&items).Error
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		c.JSON(http.StatusOK, items)
	}

}
