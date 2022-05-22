package database

import (
	"github.com/sirupsen/logrus"
	"github.com/t0239184/golearn/internal/model"
)

func (d *GormDatabase) CreateUser(user *model.User) (*model.User, error) {
	result := d.DB.Omit("UpdateAt", "Updater", "DeleteAt", "Deleter").Create(user)
	error := result.Error
	logrus.Info("RowsAffected: ", result.RowsAffected)
	logrus.Info("Return Id: ", user.ID)
	return user, error
}

func (d *GormDatabase) FindAllUser() (*[]model.User, error) {
	users := &[]model.User{}
	error := d.DB.Find(users).Error
	return users, error
}

func (d *GormDatabase) FindUserById(id int64) (*model.User, error) {
	user := &model.User{}
	result := d.DB.Take(user, id)
	error := result.Error
	return user, error
}

func (d *GormDatabase) UpdateUser(updateUser *model.User) (*model.User, error) {
	user := &model.User{}
	var error error
	error = d.DB.Take(user, updateUser.ID).Error
	if error != nil {
		logrus.Error("[UpdateUser] ", error)
		return nil, error
	}

	if (updateUser.Password != "") {
		user.Password = updateUser.Password
	}

	if (updateUser.Status != "") {
		user.Status = updateUser.Status
	}

	error = d.DB.Omit("Account", "CreateAt", "Creator", "DeleteAt", "Deleter").Save(user).Error
	if error != nil {
		logrus.Error("[UpdateUser] ", error)
		return nil, error
	}
	return user, error
}

func (d *GormDatabase) DeleteUserById(id int64) (*model.User, error) {
	user := &model.User{}
	error := d.DB.Delete(user, id).Error
	return user, error
}
