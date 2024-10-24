package user

// User
type (
	AddUserReq struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password"  binding:"required"`
		Email    string `json:"email"  binding:"required"`
		Phone    string `json:"phone"  binding:"required"`
	}
	GetUserReq struct {
		Page     int `form:"page" binding:"required"`
		PageSize int `form:"page_size" binding:"required"`
	}

	LoginReq struct {
		UserName string `json:"username" binding:"required"`
		Password string `json:"password"  binding:"required"`
	}
)

// Role
type (
	AddRoleReq struct {
		RoleName string `json:"role_name" binding:"required"`
	}

	UpdateRoleReq struct {
		Id       int    `form:"id" binding:"required"`
		RoleName string `form:"id" json:"role_name" binding:"required"`
	}

	GetRoleReq struct {
		Page     int `form:"page" binding:"required"`
		PageSize int `form:"page_size" binding:"required"`
	}

	DelRoleReq struct {
		Id int `form:"id" binding:"required"`
	}

	AddAdminRoleReq struct {
		RoleIds []int `form:"role_ids" binding:"required"`
		UserId  int   `form:"user_id" binding:"required"`
	}
)
