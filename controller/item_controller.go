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
	var itemDao database.ItemDao = controller.dao
	var userDao database.UserDao = controller.dao

	items, err := itemDao.FetchItems()
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}

	for i, item := range items {
		var ratings []model.Rating
		ratings, err = itemDao.GetAllCommentOfItem(item.ID)
		if err != nil {
			continue
		}

		var dataRatings []map[string]interface{}

		for _, rating := range ratings {
			user, err := userDao.GetUserByID(rating.UserID)
			if err != nil {
				continue
			}
			dataRatings = append(dataRatings, map[string]interface{}{
				"id":          rating.ID,
				"rating_star": rating.RatingStar,
				"comment":     rating.Comment,
				"user_name":   user.FullName,
				"user_avatar": "https://www.takadada.com/wp-content/uploads/2019/07/avatar-one-piece-1.jpg",
				"created_at":  rating.CreatedAt,
			})
		}

		items[i].Ratings = dataRatings

	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", items))
}

// Lấy chi tiết một sản phẩm
func (controller *Controller) DetaiItem(c *gin.Context) {
	var item model.IceCreamItem
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	var itemDao database.ItemDao = controller.dao
	var userDao database.UserDao = controller.dao

	item, err = itemDao.GetItemByID(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}

	var ratings []model.Rating
	ratings, err = itemDao.GetAllCommentOfItem(item.ID)
	if err != nil {
		c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", item))
	}

	var dataRatings []map[string]interface{}

	for _, rating := range ratings {
		user, err := userDao.GetUserByID(rating.UserID)
		if err != nil {
			continue
		}
		dataRatings = append(dataRatings, map[string]interface{}{
			"id":          rating.ID,
			"rating_star": rating.RatingStar,
			"comment":     rating.Comment,
			"user_name":   user.FullName,
			"user_avatar": "https://www.takadada.com/wp-content/uploads/2019/07/avatar-one-piece-1.jpg",
			"created_at":  rating.CreatedAt,
		})
	}

	item.Ratings = dataRatings

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", item))
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

func (controller *Controller) GetRatingsOfItem(c *gin.Context) {
	var ratings []model.Rating
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}
	var itemDao database.ItemDao = controller.dao
	var userDao database.UserDao = controller.dao

	ratings, err = itemDao.GetAllCommentOfItem(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}

	var dataResponse []map[string]interface{}

	for _, rating := range ratings {
		user, err := userDao.GetUserByID(rating.UserID)
		if err != nil {
			continue
		}
		dataResponse = append(dataResponse, map[string]interface{}{
			"id":          rating.ID,
			"rating_star": rating.RatingStar,
			"comment":     rating.Comment,
			"user":        user.FullName,
			"created_at":  rating.CreatedAt,
		})
	}
	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", dataResponse))

}

func (controller *Controller) CreateRating(c *gin.Context) {
	userID, ok := c.Get("userID")
	if !ok {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Get no user id from header", nil))
		return
	}

	var rating model.Rating
	if err := c.ShouldBindJSON(&rating); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Not enough info to create rating!", nil))
		return
	}

	rating.UserID = userID.(int)
	var itemDao database.ItemDao = controller.dao
	_, err := itemDao.GetItemByID(rating.ItemID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Non item exist with the id", nil))
		return
	}

	if rating.RatingStar < 0 || rating.RatingStar > 5 {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Rating star invalid!", nil))
		return
	}

	rating, err = itemDao.AddRating(rating)
	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", rating))

}
