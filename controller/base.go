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
	text := c.Query("text")
	c.JSON(http.StatusOK, text)
}
