package impl

import (
	"github.com/xilanhuaer/http-client/model/entity"
)

type UserService struct{}

func (u *UserService) Create(user entity.User) error {
	return nil
}
func (u *UserService) List(query map[string]string, limit, offset int) (users []entity.User, total int64, err error) {
	return nil, 0, nil
}
