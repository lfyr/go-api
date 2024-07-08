package service

import (
	"github.com/lfyr/go-api/model"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}
func (this *UserService) GetUserById(id int) (user model.User) {
	user = model.NewUser().First(map[string]interface{}{"id": id})
	return
}

func (this *UserService) Add(data model.User) (err error) {
	err = model.NewUser().Create(&data)
	return err
}
