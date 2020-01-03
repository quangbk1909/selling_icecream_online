package database

import (
	"vinid_project/model"
)

type UserDao interface {
	FetchUser() ([]model.User, error)
	GetUserByID(id int) (model.User, error)
	CheckUserExistByPhone(phone string) bool
	GetOrderOfUser(idUser int) ([]model.Order, error)
	Update(user *model.User) (*model.User, error)
	Store(user *model.User) (*model.User, error)
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

func (dao *Dao) CheckUserExistByPhone(phone string) bool {
	var user model.User
	dao.db.Where("phone_number = ?", phone).First(&user)
	if user != (model.User{}) {
		return true
	}
	return false
}

func (dao *Dao) GetOrderOfUser(idUser int) ([]model.Order, error) {
	var orders []model.Order
	err := dao.db.Where("user_id = ?", idUser).Find(&orders).Error
	if err != nil {
		return nil, err
	}

	return orders, nil

}

func (dao *Dao) Update(user *model.User) (*model.User, error) {
	err := dao.db.Save(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}

func (dao *Dao) Store(user *model.User) (*model.User, error) {
	err := dao.db.Create(user).Error
	if err != nil {
		return user, err
	}
	return user, nil
}
