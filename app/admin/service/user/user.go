package user

import (
	"errors"
	"github.com/lfyr/go-api/model"
	"github.com/lfyr/go-api/utils"
	"github.com/lfyr/go-api/utils/token"
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

func (this *UserService) GetUserPri(adminId int) (data model.AppAdmin) {
	data = model.NewAppAdmin().First(map[string]interface{}{"user_id=?": adminId}, []string{"Role.AppRolePrivilege.AppPrivilege"}) // , "Role.Pri", "Role.Pri.Pri"
	return
}

func (this *UserService) Login(phone, password string) (loginReq LoginReq, err error) {
	user := model.NewUser().First(map[string]interface{}{"phone = ?": phone})
	if user.Id > 0 {
		admin := model.NewAppAdmin().First(map[string]interface{}{"user_id = ?": user.Id}, []string{""})
		if admin.Id > 0 {
			loginReq.User = user
			loginReq.AdminId = admin.Id
			if user.Password == utils.HashPassword(password, user.Salt) {
				tk, err := token.SetToken(user)
				if err != nil {
					return loginReq, err
				}
				loginReq.User.Token = tk
				return loginReq, nil
			}
			return loginReq, errors.New("账户或密码不正确")
		}
		return loginReq, errors.New("用户无权限")
	}
	return loginReq, errors.New("未找到该用户")
}
