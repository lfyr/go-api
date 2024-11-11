package model

import (
	"github.com/lfyr/go-api/config/global"
	"github.com/lfyr/go-api/database/masterdb"
	"time"
)

type AppAdmin struct {
	global.Model
	UserId int            `json:"user_id"`
	IsUse  int            `json:"is_use"`
	Role   []AppAdminRole `json:"role" gorm:"foreignKey:AdminId;references:Id"`
	User   *User          `json:"user" gorm:"foreignKey:Id;references:UserId"`
}

type ListRsp struct {
	Id        int            `json:"id"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	IsUse     int            `json:"is_use"`
	UserId    int            `json:"user_id"`
	AdminRole []AppAdminRole `json:"admin_role" gorm:"foreignKey:AdminId;references:Id"`
	User      *User          `json:"user" gorm:"foreignKey:Id;references:UserId"`
	UserName  string         `json:"user_name" `
	Nickname  string         `json:"nickname"`
	Phone     string         `json:"phone"`
}

func (u *AppAdmin) TableName() string {
	return "app_admin"
}

func NewAppAdmin() *AppAdmin {
	return &AppAdmin{}
}

func (this *AppAdmin) List(whereMap map[string]interface{}, fieldSlice []string, page, size int, withSlice []string) (list []ListRsp, count int64) {
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
	tx.Joins("left join user on app_admin.user_id = user.id").Count(&count).Order("id asc").Offset((page - 1) * size).Limit(size).Find(&list)
	return
}

func (this *AppAdmin) First(whereMap map[string]interface{}, withSlice []string) (data AppAdmin) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	if len(withSlice) > 0 {
		for _, v := range withSlice {
			tx = tx.Preload(v)
		}
	}
	tx.Order("id desc").First(&data)
	return
}

func (this *AppAdmin) Many(whereMap map[string]interface{}, withSlice []string) (list []AppAdmin) {
	tx := masterdb.DB.Model(this)
	if len(whereMap) > 0 {
		for k, v := range whereMap {
			tx = tx.Where(k, v)
		}
	}
	if len(withSlice) > 0 {
		for _, v := range withSlice {
			tx = tx.Preload(v)
		}
	}
	tx.Find(&list)
	return
}

func (this *AppAdmin) Create(data *AppAdmin) (err error) {
	tx := masterdb.DB.Model(this)
	err = tx.Create(&data).Error
	return
}

func (this *AppAdmin) Update(Id int, user map[string]interface{}) (err error) {
	err = masterdb.DB.Model(this).Where("id = ?", Id).Updates(&user).Error
	return
}
