package database

import (
	"vinid_project/model"
)

type UserDao interface {
	FetchUser() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	//GetItemInStore(id int) ([]model.IceCreamItem, error)
	//GetStoreAroundHere(latitude float64, longitude float64, distance float64) ([]model.Store, error)
	// Update(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Store(item *model.IceCreamItem) (*model.IceCreamItem, error)
	// Delete(id int64) (bool, error)
}

func (dao *Dao) FetchUser() ([]model.User, error) {
	var users []model.User
	err := dao.db.Find(&users).Error
	if err != nil {
		return nil, err
	}
	return users, nil
}

func (dao *Dao) GetUserByID(id int) (model.User, error) {
	var user model.User
	err := dao.db.First(&user, id).Error
	if err != nil {
		return model.User{}, err
	}

	return user, nil
}
