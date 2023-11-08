package impl

import (
	"fmt"

	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/model/entity"
)

type UserService struct{}

func (u *UserService) Create(user entity.User) error {
	err := global.DB.Where("account = ?", user.Account).First(&user).Error
	if err == nil {
		return fmt.Errorf("用户已存在")
	}
	return global.DB.Create(&user).Error
}
