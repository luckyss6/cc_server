package dao

import (
	"cc_server/common"
	"cc_server/model"
)

var UserD = &UserDao{}

type UserDao struct {
}

func (u *UserDao) GetUserByNicknameAndPassword(nickname, password string) (*model.UserInfo, error) {
	result := new(model.UserInfo)
	err := common.DB.Table("user_info").Where("nickname = ? and password = ?", nickname, password).Find(&result).Error
	if err != nil {
		return nil, err
	}
	return result, nil
}
