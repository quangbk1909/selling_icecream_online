package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

type Controller struct {
	db *gorm.DB
}

func NewController(db *gorm.DB) *Controller {
	return &Controller{
		db: db,
	}
}
func (this *Controller) TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HI quang"})
}
