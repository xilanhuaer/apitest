package impl

import (
	"context"
	"fmt"
	"github.com/xilanhuaer/http-client/common/response"
	"github.com/xilanhuaer/http-client/dal/model"
	"github.com/xilanhuaer/http-client/dal/query"
	"github.com/xilanhuaer/http-client/dal/vo"
	"github.com/xilanhuaer/http-client/global"
	"github.com/xilanhuaer/http-client/utils"
)

type UserService struct{}

func (userService *UserService) Register(user model.User) error {
	u := query.User
	_, err := u.WithContext(context.Background()).Where(u.Account.Eq(user.Account)).First()
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
		return query.User.WithContext(context.Background()).Create(&user)
	}
	return fmt.Errorf("用户已存在，请重试")
}

func (userService *UserService) Login(account, password string) (userinfo vo.Userinfo, err error) {
	u := query.User
	user, err := u.WithContext(context.Background()).Where(u.Account.Eq(account)).First()
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

func (userService *UserService) List(params map[string]string, limit, offset int) (response.Page, error) {
	query1 := global.DB.Model(&model.User{})
	var userList []model.User
	var total int64
	for key, value := range params {
		if value != "" {
			query1.Where(fmt.Sprintf("%s = ?", key), value)
		}
	}
	err := query1.Count(&total).Error
	if err != nil {
		return response.Page{}, err
	}

	err = query1.Limit(limit).Offset(offset).Find(&userList).Error
	if err != nil {
		return response.Page{}, err
	}
	return response.Page{
		List:  userList,
		Total: total,
	}, err
}

// Find 根据id查询用户信息
func (userService *UserService) Find(id int) (vo.Userinfo, error) {
	var (
		userinfo vo.Userinfo
	)
	user, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(int32(id))).First()
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

func (userService *UserService) UpdatePassword(oldPassword, newPassword string, id int32) error {
	user, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(id)).First()
	if err != nil {
		return err
	}
	if utils.RSA_Decrypt(user.Password, "private.pem") == oldPassword {
		_, err = query.User.WithContext(context.Background()).Where(query.User.ID.Eq(id)).
			Updates(map[string]interface{}{
				"password": utils.RSA_Encrypt(newPassword, "public.pem"),
			})
		if err != nil {
			return err
		}
	}
	return fmt.Errorf("密码错误")
}

func (userService *UserService) Update(id int32, message interface{}) error {
	data := utils.StructToMap(message)
	_, err := query.User.WithContext(context.Background()).Where(query.User.ID.Eq(id)).
		Updates(data)
	return err
}
