package model

import (
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/database/masterdb"
)

type User struct {
	global.Model
	UserName      string `json:"user_name"`
	Nickname      string `json:"nickname"`
	Password      string `json:"password"`
	Email         string `json:"email"`
	Phone         string `json:"phone"`
	Ip            string `json:"ip"`
	Status        string `json:"status"`
	Token         string `json:"token"`
	Avatar        string `json:"avatar"`
	DeletedStatus int    `json:"deleted_status"`
}

func (u User) TableName() string {
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

func (this *User) First(whereMap map[string]interface{}) (user User) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	tx.Order("id desc").First(&user)
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

func (this *User) Create(data User) (User, error) {
	tx := masterdb.DB.Model(this)
	err := tx.Create(&data).Error
	return data, err
}

func (this *User) Update(Id int, user map[string]interface{}) (err error) {
	err = masterdb.DB.Model(this).Where("id = ?", Id).Updates(&user).Error
	return
}

func (this *User) Delete(Ids []int) (err error) {
	err = masterdb.DB.Model(this).Where("id in (?)", Ids).Delete(&this).Error
	return
}
