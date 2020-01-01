package controller

import (
	"net/http"
	"strconv"
	"vinid_project/database"
	"vinid_project/model"
	"vinid_project/utility"

	"github.com/gin-gonic/gin"
)

type UserAuthicationJson struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (controller *Controller) GetUsers(c *gin.Context) {
	var users []model.User
	var userDao database.UserDao = controller.dao

	users, err := userDao.FetchUser()
	var dataResponse []map[string]interface{}
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
		return
	}
	for _, user := range users {
		dataResponse = append(dataResponse, map[string]interface{}{
			"id":           user.ID,
			"full_name":    user.FullName,
			"phone_number": user.PhoneNumber,
			"address":      user.Address,
			"vinid_point":  user.VinidPoint,
			"created_at":   user.CreatedAt,
		})
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", dataResponse))

}

func (controller *Controller) DetailUser(c *gin.Context) {
	var user model.User
	var userDao database.UserDao = controller.dao
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	user, err = userDao.GetUserByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Non user exist with the id", nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", user))
}

func (controller *Controller) GetOrderOfUser(c *gin.Context) {
	var orders []model.Order
	id, err := strconv.Atoi(c.Param("id"))

	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	var userDao database.UserDao = controller.dao

	orders, err = userDao.GetOrderOfUser(id)
	if err != nil {
		message := "Non orders exist with user id!"
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, message, nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful!", orders))

}

func (controller *Controller) TestFile(c *gin.Context) {
	c.File("./resources/store_images/1.png")

}
