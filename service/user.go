package service

import (
	"cc_server/dao"
	"cc_server/model"
)

var UserSrv = &UserService{}

type UserService struct {
}

func (u *UserService) GetUserInfo(nickname, password string) (bool, *model.UserInfo, error) {
	// 从数据库中查询用户信息
	userInfo, err := dao.UserD.GetUserByNicknameAndPassword(nickname, password)
	if err != nil {
		return false, nil, err
	}
	if userInfo.ID == 0 {
		return false, nil, nil
	}
	return true, userInfo, nil
}
