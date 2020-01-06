package database

import (
	"vinid_project/model"
)

type StoreDao interface {
	FetchStore() ([]model.Store, error)
	GetStoreByID(id int) (model.Store, error)
	GetItemInStore(id int) ([]model.IceCreamItem, error)
	GetStoreAroundHere(latitude float64, longitude float64, distance float64) ([]model.Store, error)

	// Update(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Store(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Delete(id int64) (bool, error)
}

func (dao *Dao) FetchStore() ([]model.Store, error) {
	var stores []model.Store
	err := dao.db.Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}

func (dao *Dao) GetStoreByID(id int) (model.Store, error) {
	var store model.Store
	err := dao.db.First(&store, id).Error
	if err != nil {
		return model.Store{}, err
	}

	return store, nil
}

func (dao *Dao) GetItemInStore(id int) ([]model.IceCreamItem, error) {
	var items []model.IceCreamItem
	var store model.Store
	err := dao.db.First(&store, id).Error
	if err != nil {
		return nil, err
	}
	dao.db.Model(&store).Related(&items, "IceCreamItems")

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

func (dao *Dao) GetStoreAroundHere(latitude float64, longitude float64, distance float64) ([]model.Store, error) {
	var stores []model.Store
	err := dao.db.Where("latitude > ? AND latitude < ? AND longitude > ? AND longitude < ?", latitude-distance, latitude+distance, longitude-distance, longitude+distance).Find(&stores).Error
	if err != nil {
		return nil, err
	}
	return stores, nil
}
