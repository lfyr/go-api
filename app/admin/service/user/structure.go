package user

import "github.com/lfyr/go-api/model"

type LoginReq struct {
	model.User `json:"user"`
	AdminId    int    `json:"admin_id"`
	Token      string `json:"token"`
}
