package controller

import (
	"net/http"
	"strconv"
	"vinid_project/authentication"
	"vinid_project/database"
	"vinid_project/model"
	"vinid_project/utility"

	"golang.org/x/crypto/bcrypt"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) Register(c *gin.Context) {
	var registerJson model.AuthenticationJson

	if err := c.ShouldBindJSON(&registerJson); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Not enough info to register", nil))
		return
	}

	var userDao database.UserDao = controller.dao
	if userDao.CheckUserExistByPhone(registerJson.PhoneNumber) {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Phone number already exist!", nil))
		return
	}

	hash, err := bcrypt.GenerateFromPassword([]byte(registerJson.Password), bcrypt.MinCost)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error", nil))
		return
	}
	registerJson.Password = string(hash)

	user := model.User{
		PhoneNumber: registerJson.PhoneNumber,
		Password:    registerJson.Password,
	}

	_, err = userDao.Store(&user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error", nil))
		return
	}

	token, err := authentication.MakeJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error. Login again to continue!", nil))
		return
	}
	c.JSON(http.StatusOK, utility.MakeResponse(200, "Register successful!", gin.H{"token": token}))

}

func (controller *Controller) UpdateInfo(c *gin.Context) {
	//var user model.User
	var userInfoJson model.UserInfoJson
	if err := c.ShouldBindJSON(&userInfoJson); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "There is no data to update!", nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Update user info successful!", userInfoJson))
}

func (controller *Controller) Login(c *gin.Context) {
	var authenticationJson model.AuthenticationJson
	if err := c.ShouldBindJSON(&authenticationJson); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Authenticate fail!", nil))
		return
	}

	var userDao database.UserDao = controller.dao

	user, err := userDao.Authenticate(authenticationJson)
	if err != nil {
		c.JSON(http.StatusUnauthorized, utility.MakeResponse(401, err.Error(), nil))
		return
	}

	token, err := authentication.MakeJWT(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error. Login again to continue!", nil))
		return
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Authenticate success!", gin.H{"user": user, "token": token}))
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
