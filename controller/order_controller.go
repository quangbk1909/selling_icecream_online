package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type ItemOrderJson struct {
	ItemID   int `json:"item_id"`
	Quantity int `json:"quantity"`
}

type OrderJson struct {
	UserID   int             `json:"user_id"`
	TotalFee int             `json:"total_fee"`
	Items    []ItemOrderJson `json:"items"`
}

func (controller *Controller) CreatOrder(c *gin.Context) {
	var orderInfo OrderJson
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusOK, orderInfo)

}
