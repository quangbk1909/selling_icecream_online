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

	for i, item := range items {
		rows, err := dao.db.Table("item_image").Where("ice_cream_item_id = ?", item.ID).Select("image_path").Rows()
		if err != nil {
			continue
		}

		for rows.Next() {
			var image_path string
			rows.Scan(&image_path)
			items[i].ImagePaths = append(items[i].ImagePaths, image_path)
		}
		rows.Close()
	}
	return items, nil
}

func (dao *Dao) GetItemByID(id int) (model.IceCreamItem, error) {
	var item model.IceCreamItem
	err := dao.db.First(&item, id).Error
	if err != nil {
		return model.IceCreamItem{}, err
	}

	rows, err := dao.db.Table("item_image").Where("ice_cream_item_id = ?", item.ID).Select("image_path").Rows()
	if err != nil {
		item.ImagePaths = nil
	}

	for rows.Next() {
		var image_path string
		rows.Scan(&image_path)
		item.ImagePaths = append(item.ImagePaths, image_path)
	}
	rows.Close()

	return item, nil
}

func (dao *Dao) SearchFullTextItem(text string) ([]model.IceCreamItem, error) {
	var items []model.IceCreamItem
	query := "SELECT * FROM ice_cream_item WHERE MATCH (name,type) AGAINST ('" + text + "' IN BOOLEAN MODE);"
	err := dao.db.Raw(query).Scan(&items).Error
	if err != nil {
		return nil, err
	}

	for i, item := range items {
		rows, err := dao.db.Table("item_image").Where("ice_cream_item_id = ?", item.ID).Select("image_path").Rows()
		if err != nil {
			continue
		}

		for rows.Next() {
			var image_path string
			rows.Scan(&image_path)
			items[i].ImagePaths = append(items[i].ImagePaths, image_path)
		}
		rows.Close()
	}
	return items, nil
}
