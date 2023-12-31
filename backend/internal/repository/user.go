package repository

import (
	"github.com/luckyss6/cc_server/pkg/model"
	"github.com/luckyss6/cc_server/pkg/storage"
	"gorm.io/gorm"
)

type UserRepo struct {
}

func NewUserRepo() *UserRepo {
	return &UserRepo{}
}

func (u *UserRepo) GetUserInfoByUsername(username string) (result *model.UserInfo, err error) {
	err = storage.DB.Where("username = ?", username).First(&result).Error
	if err == gorm.ErrRecordNotFound {
		return nil, nil
	}
	return
}

func (u *UserRepo) CreateUserInfo(userInfo *model.UserInfo) error {
	return storage.DB.Create(userInfo).Error
}
