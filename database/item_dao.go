package database

import (
	"vinid_project/model"
)

type ItemDao interface {
	FetchItems() ([]model.IceCreamItem, error)
	GetItemByID(id int) (model.IceCreamItem, error)
	SearchFullTextItem(text string) ([]model.IceCreamItem, error)
	// Update(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Store(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Delete(id int64) (bool, error)
}

func (dao *Dao) FetchItems() ([]model.IceCreamItem, error) {
	var items []model.IceCreamItem
	err := dao.db.Find(&items).Error
	if err != nil {
		return nil, err
	}

	return items, nil
}

func (dao *Dao) GetItemByID(id int) (model.IceCreamItem, error) {
	var item model.IceCreamItem
	err := dao.db.First(&item, id).Error
	if err != nil {
		return model.IceCreamItem{}, err
	}

	return item, nil
}

func (dao *Dao) SearchFullTextItem(text string) ([]model.IceCreamItem, error) {
	var items []model.IceCreamItem
	query := "SELECT * FROM ice_cream_item WHERE MATCH (name,type) AGAINST ('" + text + "' IN BOOLEAN MODE);"
	err := dao.db.Raw(query).Scan(&items).Error
	if err != nil {
		return nil, err
	}
	return items, nil
}
