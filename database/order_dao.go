package database

import (
	"errors"
	"strconv"
	"vinid_project/model"
	"vinid_project/utility"
)

type OrderDao interface {
	GetDetailOrderByID(id int) (model.OrderDetail, error)
	// Update(item *model.IceCreamItem) (*model.IceCreamItem, error)
	StoreOrder(orderJson model.OrderJson) (model.OrderDetail, error)
	// Delete(id int64) (bool, error)
}

func (dao *Dao) GetDetailOrderByID(id int) (model.OrderDetail, error) {
	var orderDetail model.OrderDetail
	var order model.Order
	var orderItems []model.OrderItem
	err := dao.db.First(&order, id).Error
	if err != nil {
		return model.OrderDetail{}, err
	} else {
		// dao.db.Model(&order).Related(&items, "IceCreamItems")
		// order.IceCreamItems = items
		// return order, nil
		orderDetail.OrderInfo = order
		err = dao.db.Where("order_id = ?", id).Find(&orderItems).Error
		if err != nil {
			orderDetail.Items = nil
		} else {
			for _, orderItem := range orderItems {
				var item model.IceCreamItem
				err = dao.db.First(&item, orderItem.IceCreamItemId).Error
				if err != nil {
					continue
				} else {
					itemInOrder := model.ItemInOrder{ItemInfo: item, Quantity: orderItem.Quantity}
					orderDetail.Items = append(orderDetail.Items, itemInOrder)
				}
			}
		}
		return orderDetail, nil

	}
}

func (dao *Dao) StoreOrder(orderJson model.OrderJson) (model.OrderDetail, error) {
	var order model.Order
	var orderItems []model.OrderItem
	if len(orderJson.Items) != 0 {
		for _, item := range orderJson.Items {
			var itemInfo model.IceCreamItem
			err := dao.db.First(&itemInfo, item.ItemID).Error
			if err != nil {
				return model.OrderDetail{}, errors.New("No item found with this id :" + strconv.Itoa(item.ItemID))
			}
			orderItems = append(orderItems, model.OrderItem{IceCreamItemId: item.ItemID, Quantity: item.Quantity})
		}
	}

	order.UserID = orderJson.UserID
	order.ShipFee = orderJson.ShipFee
	order.TotalFee = orderJson.TotalFee
	order.Status = orderJson.Status
	deliveryAddress, _ := utility.GetAddressFromCoordinates(orderJson.Coordinates.Latitude, orderJson.Coordinates.Longitude)
	order.DeliveryAddress = deliveryAddress

	err := dao.db.Create(&order).Error
	if err != nil {
		return model.OrderDetail{}, errors.New("Internal server error")
	}

	if len(orderItems) > 0 {
		for _, orderItem := range orderItems {
			orderItem.OrderID = order.ID
			dao.db.Create(&orderItem)
		}
	}

	return dao.GetDetailOrderByID(order.ID)

}
