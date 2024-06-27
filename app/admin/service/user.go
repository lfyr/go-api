package service

import (
	"github.com/lfyr/go-api/model"
)

func GetUserById(id int) (user model.User) {
	user = model.NewUser().First(map[string]interface{}{"id": id})
	return
}

func Add(data model.User) (err error) {
	err = model.NewUser().Create(&data)
	return err
}
