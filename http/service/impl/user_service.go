package impl

import (
	"fmt"
	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
	"github.com/xilanhuaer/http-client/dal/vo"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

type UserService struct{}

func (u *UserService) Register(user model.User) error {
	err := global.DB.Where("account = ?", user.Account).First(&model.User{}).Error
	if err != nil {
		{
			if !utils.IsAccount(user.Account) {
				return fmt.Errorf("账号格式错误")
			}
			if !utils.IsPassword(user.Password) {
				return fmt.Errorf("密码格式错误")
			}
			if user.Email != "" && !utils.IsEmail(user.Email) {
				return fmt.Errorf("邮箱格式错误")
			}
			if user.Phone != "" && !utils.IsPhone(user.Phone) {
				return fmt.Errorf("手机号格式错误")
			}
		}
		user.Password = utils.RSA_Encrypt(user.Password, "./public.pem")
		return global.DB.Create(&user).Error
	}
	return fmt.Errorf("用户已存在，请重试")
}

func (u *UserService) Login(account, password string) (userinfo vo.Userinfo, err error) {
	var user model.User
	err = global.DB.Where("account = ?", account).First(&user).Error
	if err != nil {
		return vo.Userinfo{}, err
	}
	if utils.RSA_Decrypt(user.Password, "./private.pem") == password {
		{
			userinfo.ID = user.ID
			userinfo.Name = user.Name
			userinfo.Account = user.Account
			userinfo.Avatar = user.Avatar
			userinfo.Email = user.Email
			userinfo.Phone = user.Phone
			userinfo.Description = user.Description
			userinfo.Token, _ = utils.GenJWT(user.ID, user.Name)
		}
		return userinfo, nil
	}
	return vo.Userinfo{}, fmt.Errorf("密码错误")
}

func (u *UserService) List(params map[string]string, limit, offset int) (response.Page, error) {
	query := global.DB.Model(&model.User{})
	var userList []model.User
	var total int64
	for key, value := range params {
		if value != "" {
			query.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	err := query.Count(&total).Error
	if err != nil {
		return response.Page{}, err
	}

	err = query.Limit(limit).Offset(offset).Find(&userList).Error
	if err != nil {
		return response.Page{}, err
	}
	return response.Page{
		List:  userList,
		Total: total,
	}, err
}

// Find 根据id查询用户信息
func (u *UserService) Find(id string) (vo.Userinfo, error) {
	var (
		userinfo vo.Userinfo
		user     model.User
	)
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return vo.Userinfo{}, err
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

func (u *UserService) UpdatePassword(oldPassword, newPassword, id string) error {
	var user model.User
	err := global.DB.Where("id = ?", id).First(&user).Error
	if err != nil {
		return err
	}
	if utils.RSA_Decrypt(user.Password, "./private.pem") != oldPassword {
		return fmt.Errorf("密码错误请重试")
	}
	err = global.DB.Model(&model.User{}).Update("password", utils.RSA_Encrypt(newPassword, "./public.pem")).Error
	return err
}

func (u *UserService) Update(id string, message interface{}) error {
	data := utils.StructToMap(message)
	return global.DB.Model(&model.User{}).Where("id = ?", id).Updates(data).Error
}
