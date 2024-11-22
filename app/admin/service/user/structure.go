package user

import "github.com/lfyr/go-api/model"

type LoginReq struct {
	model.User `json:"user"`
	Token      string `json:"token"`
}

type GetPriByRoleIdRes struct {
	Id         int    `json:"id"`
	PriName    string `json:"pri_name"`
	ActionName string `json:"action_name"`
	ParentId   int    `json:"parent_id"`
	RoleId     int    `json:"role_id"`
	PriId      int    `json:"pri_id"`
}
