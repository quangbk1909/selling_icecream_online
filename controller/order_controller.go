package controller

import (
	"net/http"
	"strconv"
	"vinid_project/database"
	"vinid_project/model"
	"vinid_project/utility"

	"github.com/gin-gonic/gin"
)

func (controller *Controller) CreateOrder(c *gin.Context) {
	userID, ok := c.Get(model.UserIDInClaimJWT)
	if !ok {
		c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Get no user id from header", nil))
		return
	}

	var orderInfo model.OrderJson
	var orderDetail model.OrderDetail
	if err := c.ShouldBindJSON(&orderInfo); err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Not enough info to create order", nil))
		return
	}

	if orderInfo.TotalFee < 0 || orderInfo.ShipFee < 0 {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Fee can not less than 0", nil))
		return
	}

	orderInfo.UserID = userID.(int)

	var userDao database.UserDao = controller.dao
	user, err := userDao.GetUserByID(orderInfo.UserID)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "User does not exist!", nil))
		return
	}

	if orderInfo.Status == 1 {
		if user.VinidPoint < orderInfo.TotalFee {
			c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Vinid point does not enough to pay the order!", nil))
			return
		} else {
			var orderDao database.OrderDao = controller.dao
			orderDetail, err = orderDao.StoreOrder(orderInfo)
			if err != nil {
				c.JSON(http.StatusBadRequest, utility.MakeResponse(404, err.Error(), nil))
				return
			}
			user.VinidPoint -= orderInfo.TotalFee
			_, err = userDao.Update(&user)
			if err != nil {
				c.JSON(http.StatusInternalServerError, utility.MakeResponse(500, "Internal server error!", nil))
				return
			}
		}
	} else {
		var orderDao database.OrderDao = controller.dao
		orderDetail, err = orderDao.StoreOrder(orderInfo)
		if err != nil {
			c.JSON(http.StatusBadRequest, utility.MakeResponse(404, err.Error(), nil))
			return
		}
	}

	c.JSON(http.StatusOK, utility.MakeResponse(200, "Create order successfully!", orderDetail))

}

func (controller *Controller) DetaiOrder(c *gin.Context) {
	var orderDetail model.OrderDetail
	var orderDao database.OrderDao = controller.dao

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Bad request. Can not convert id parameter to int", nil))
		return
	}

	orderDetail, err = orderDao.GetDetailOrderByID(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, utility.MakeResponse(404, "Order not exist with the id!", nil))
		return
	}
	c.JSON(http.StatusOK, utility.MakeResponse(200, "Request successful", orderDetail))
}
