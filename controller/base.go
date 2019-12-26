package controller

import (
	"vinid_project/model"

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
	// var item model.IceCreamItem
	// //var orderItem model.OrderItem
	// var orders []model.Order
	// // this.db.Preload("order")
	// this.db.First(&item)
	// this.db.Model(&item).Related(&orders, "Orders")
	// c.JSON(200, orders)

	var items []model.IceCreamItem
	//var orderItem model.OrderItem
	var order model.Order
	// this.db.Preload("order")
	this.db.First(&order)
	this.db.Model(&order).Related(&items, "IceCreamItems")
	c.JSON(200, items)

	// rows, err := db.Raw(`
	//  	SELECT full_name,phone_number
	//  	FROM user
	// `).Rows()
	// for rows.Next() {
	// 	var fullName string
	// 	rows.Scan(&fullName)

	// 	fmt.Println(fullName)
	// }

}
