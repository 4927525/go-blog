package service

import (
	"errors"
	"go-blog/dao"
	"go-blog/models"
	"go-blog/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	passwd = utils.Md5Crypt(passwd, "")
	user := dao.Login(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码错误")
	}
	uid := user.Uid
	// token
	token,err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token未能生成")
	}
	var userInfo models.UserInfo
	userInfo.Uid = uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	lr := &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
