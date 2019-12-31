package controller

import (
	"net/http"
	"strconv"
	"vinid_project/database"
	"vinid_project/model"
	"vinid_project/utility"

	"github.com/gin-gonic/gin"
)

//Lấy danh sách tất cả các sản phẩm
func (controller *Controller) GetItems(c *gin.Context) {
	var items []model.IceCreamItem
	var itemDao database.ItemDao
	itemDao = controller.dao
	items, err := itemDao.FetchItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, items)
}

// Lấy chi tiết một sản phẩm
func (controller *Controller) DetaiItem(c *gin.Context) {
	var dataResponse map[string]interface{}
	var item model.IceCreamItem
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var itemDao database.ItemDao
	itemDao = controller.dao

	item, err = itemDao.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	dataResponse = map[string]interface{}{
		"id":         item.ID,
		"name":       item.Name,
		"type":       item.Type,
		"image_path": item.ImagePath,
		"price":      item.Price,
		"created_at": item.CreatedAt,
	}

	c.JSON(http.StatusOK, dataResponse)
}

// Lấy hình ảnh của 1 item
func (controller *Controller) GetItemImage(c *gin.Context) {
	name := c.Param("name")
	c.File("./resources/item_images/" + name)
}

// Search item theo text
func (controller *Controller) SearchItem(c *gin.Context) {

	textSearch := c.Query("text")
	if textSearch == "" {
		c.JSON(http.StatusBadRequest, gin.H{"message": "No text search"})
	} else {
		textSearch = utility.StringSearchText(textSearch)
		var itemDao database.ItemDao
		itemDao = controller.dao

		items, err := itemDao.SearchFullTextItem(textSearch)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, items)
	}

}
