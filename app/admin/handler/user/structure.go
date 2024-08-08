package user

type AddUserReq struct {
	UserName string `json:"user_name" binding:"required"`
	Password string `json:"password"  binding:"required"`
	Email    string `json:"email"  binding:"required"`
	Phone    string `json:"phone"  binding:"required"`
}

type GetUserReq struct {
	Id int `form:"id" binding:"required"`
}

type LoginReq struct {
	Phone    string `json:"phone" binding:"required"`
	Password string `json:"password"  binding:"required"`
}
