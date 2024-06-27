package model

import (
	"github.com/lfyr/go-api/config"
	"github.com/lfyr/go-api/database/masterdb"
)

type User struct {
	config.Model
	UserName string
	Password string
	Email    string
	Phone    string
	Ip       string
	Status   string
	Token    string
	Avatar   string
}

func (u *User) TableName() string {
	return "user"
}

func NewUser() *User {
	return &User{}
}

func (this *User) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []User, count int64) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}

	if len(fieldSlice) > 0 {
		tx = tx.Select(fieldSlice)
	}

	if len(withSlice) > 0 {
		for _, v := range withSlice {
			tx = tx.Preload(v)
		}
	}
	tx.Count(&count).Order("id asc").Offset((page - 1) * size).Limit(size).Find(&list)
	return
}

func (this *User) First(whereMap map[string]interface{}) (nodeVisitCall User) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&nodeVisitCall)
	return
}

func (this *User) Many(whereMap map[string]interface{}) (app []User) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Preload("Category").Find(&app)
	return
}

func (this *User) Create(data *User) (err error) {
	tx := masterdb.DB.Model(this)
	err = tx.Create(&data).Error
	return
}

func (this *User) Update(Id int, user map[string]interface{}) (err error) {
	err = masterdb.DB.Model(this).Where("id = ?", Id).Updates(&user).Error
	return
}
