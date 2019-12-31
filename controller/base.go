package controller

import (
	"net/http"
	"vinid_project/database"

	"github.com/gin-gonic/gin"
)

type Controller struct {
	dao *database.Dao
}

func NewController(dao *database.Dao) *Controller {
	return &Controller{
		dao: dao,
	}
}

func (this *Controller) TestController(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"message": "HI quang"})
}
