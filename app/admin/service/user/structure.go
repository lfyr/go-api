package user

import "github.com/lfyr/go-api/model"

type LoginReq struct {
	model.User `json:"user"`
	Token      string `json:"token"`
}
