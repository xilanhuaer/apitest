package impl

import (
	"fmt"

	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/model/entity"
	"github.com/xilanhuaer/http-client/utils"
)

type UserService struct{}

func (u *UserService) Register(user entity.User) error {
	err := global.DB.Where("account = ?", user.Account).First(&entity.User{}).Error
	if err != nil {
		user.Password = utils.RSA_Encrypt(user.Password, "./public.pem")
		return global.DB.Create(&user).Error
	}
	return fmt.Errorf("用户已存在，请重试")
}

func (u *UserService) Login(account, password string) (user entity.User, err error) {
	err = global.DB.Where("account = ?", account).First(&user).Error
	if err != nil {
		return entity.User{}, err
	}
	if utils.RSA_Decrypt(user.Password, "./private.pem") == password {
		return user, nil
	}
	return entity.User{}, fmt.Errorf("密码错误")
}

func (u *UserService) List(query map[string]string, limit, offset int) (users []entity.User, total int64, err error) {
	return nil, 0, nil
}
func (u *UserService) Find(id string) (entity.UserInfo, error) {
	var (
		userinfo entity.UserInfo
		user     entity.User
	)
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return entity.UserInfo{}, err
	}
	{
		userinfo.ID = user.ID
		userinfo.Account = user.Account
		userinfo.Name = user.Name
		userinfo.Avatar = user.Avatar
		userinfo.Email = user.Email
		userinfo.Phone = user.Phone
		userinfo.Description = user.Description
	}
	return userinfo, err
}
