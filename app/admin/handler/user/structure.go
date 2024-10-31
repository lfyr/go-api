package user

// User
type (
	AddUserReq struct {
		UserName string `json:"user_name" binding:"required"`
		Password string `json:"password"  binding:"required"`
		Email    string `json:"email"`
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
	UpdateUserReq struct {
		Id       int    `json:"id" binding:"required"`
		UserName string `json:"user_name"`
		Phone    string `json:"phone"`
	}
	ToAssignReq struct {
		Id int `form:"id" binding:"required"`
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
		RoleId []int `form:"roleId" binding:"required"`
		UserId int   `form:"userId" binding:"required"`
	}
)
