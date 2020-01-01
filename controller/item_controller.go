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
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}
	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", items))
}

// Lấy chi tiết một sản phẩm
func (controller *Controller) DetaiItem(c *gin.Context) {
	var dataResponse map[string]interface{}
	var item model.IceCreamItem
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	var itemDao database.ItemDao
	itemDao = controller.dao

	item, err = itemDao.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
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

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", dataResponse))
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
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "No text search!", nil))
	} else {
		textSearch = utility.StringSearchText(textSearch)
		var itemDao database.ItemDao
		itemDao = controller.dao

		items, err := itemDao.SearchFullTextItem(textSearch)
		if err != nil {
			c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error", nil))
			return
		}
		c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", items))
	}

}
