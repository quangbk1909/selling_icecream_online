package controller

import (
	"log"
	"net/http"
	"vinid_project/model"

	"github.com/gin-gonic/gin"
)

type UserAuthicationJson struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}

func (controller *Controller) GetUsers(c *gin.Context) {
	var users []model.User
	err := controller.db.Find(&users).Error
	var dataResponse []map[string]interface{}
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
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

	c.JSON(http.StatusOK, dataResponse)

}

func (controller *Controller) DetailUser(c *gin.Context) {
	var user model.User
	id := c.Param("id")
	err := controller.db.First(&user, id).Error
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, user)
}

func (controller *Controller) TestFile(c *gin.Context) {
	c.File("./resources/store_images/1.png")

}
