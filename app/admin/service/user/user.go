package user

import (
	"errors"
	"github.com/lfyr/go-api/database/masterdb"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
)

type UserService struct{}

func NewUserService() *UserService {
	return &UserService{}
}

func (this *UserService) List(whereMap map[string]interface{}, fieldSlice []string, page int, size int, withSlice []string) (list []model.User, count int64) {
	list, count = model.NewUser().List(whereMap, fieldSlice, page, size, withSlice)
	return
}

func (this *UserService) First(whereMap map[string]interface{}) (user model.User) {
	user = model.NewUser().First(whereMap)
	return
}

func (this *UserService) GetUserById(id int, withSlice []string) (user model.User) {
	user = model.NewUser().First(map[string]interface{}{"id": id})
	return
}

func (this *UserService) Add(data model.User) (err error) {

	data.Password = utils.HashPassword(data.Password)
	user, _ := model.NewUser().Create(data)
	adminData := model.AppAdmin{
		UserId: user.Id,
		IsUse:  1,
	}
	err = model.NewAppAdmin().Create(&adminData)
	return err
}

func (this *UserService) Update(data model.User) (err error) {

	data.Password = utils.HashPassword(data.Password)
	dataMap := map[string]interface{}{
		"user_name": data.UserName,
		"phone":     data.Phone,
	}
	err = model.NewUser().Update(data.Id, dataMap)
	return err
}

func (this *UserService) Login(userName, password string) (loginReq LoginReq, err error) {
	whereMap := getUserNameWhereMap(userName)
	user := model.NewUser().First(whereMap)
	if user.Id > 0 {
		admin := model.NewAppAdmin().First(map[string]interface{}{"user_id = ?": user.Id}, []string{})
		if admin.Id > 0 && admin.IsUse == 1 {
			loginReq.User = user
			if utils.VerifyPassword(user.Password, password) {
				tk, err := token.SetToken(user)
				if err != nil {
					return loginReq, err
				}
				loginReq.User.Token = tk
				loginReq.Token = tk
				return loginReq, nil
			}
			return loginReq, errors.New("账户或密码不正确")
		}
		return loginReq, errors.New("用户无权限")
	}
	return loginReq, errors.New("未找到该用户")
}

func (this *UserService) Delete(ids []int) (err error) {
	data := map[string]interface{}{"deleted_status": 1}
	err = masterdb.DB.Model(model.User{}).Where("id in (?)", ids).Updates(data).Error
	return err
}

func getUserNameWhereMap(userName string) map[string]interface{} {
	if utils.IsPhoneNumber(userName) {
		return map[string]interface{}{"phone = ?": userName}
	}
	return map[string]interface{}{"user_name = ?": userName}
}
