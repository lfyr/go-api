package user

// User
type (
	AddUserReq struct {
		Phone string `json:"phone"  binding:"required"`
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
	DelUserReq struct {
		Ids []int `json:"ids" binding:"required"`
	}
	ToAssignReq struct {
		Id int `form:"id" binding:"required"`
	}
	InfoReq struct {
		Id       int         `json:"id"`
		UserId   int         `json:"user_id"`
		IsUse    int         `json:"is_use"`
		UserName string      `json:"user_name"`
		Nickname string      `json:"nickname"`
		Password string      `json:"password"`
		Email    string      `json:"email"`
		Phone    string      `json:"phone"`
		Ip       string      `json:"ip"`
		Token    string      `json:"token"`
		Avatar   string      `json:"avatar"`
		Routes   interface{} `json:"routes"`
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
		RoleName string `form:"role_name"`
		Page     int    `form:"page" binding:"required"`
		PageSize int    `form:"page_size" binding:"required"`
	}

	DelRoleReq struct {
		Id int `form:"id" binding:"required"`
	}

	AddAdminRoleReq struct {
		RoleId  []int `form:"roleId" binding:"required"`
		AdminId int   `form:"adminId" binding:"required"`
	}

	GetRolePrivilegeReq struct {
		RoleId int `form:"role_id" binding:"required"`
	}

	PrivilegeTree struct {
		Id         int             `json:"id"`
		PriName    string          `json:"pri_name"`
		ActionName string          `json:"action_name"`
		MenuName   string          `json:"menu_name"`
		ParentId   int             `json:"parent_id"`
		Children   []PrivilegeTree `json:"children"`
	}

	AddRolePrivilegeReq struct {
		RoleId int   `json:"roleId" binding:"required"`
		PriId  []int `json:"priId" binding:"required"`
	}
)

// Admin
type (
	GetAdminListReq struct {
		UserName string `form:"user_name"`
		Page     int    `form:"page" binding:"required"`
		PageSize int    `form:"page_size" binding:"required"`
	}

	GetAdminListRsp struct {
		Id       int    `json:"id"`
		UserId   int    `json:"user_id"`
		IsUse    int    `json:"is_use"`
		UserName string `json:"user_name"`
		Nickname string `json:"nickname"`
		Phone    string `json:"phone"`
		RoleStr  string `json:"role_str"`
	}
	AddAdminReq struct {
		Phone string `json:"phone"  binding:"required"`
	}
	UpdateAdminReq struct {
		Id    int `json:"id"`
		IsUse int `json:"is_use"`
	}
)

type (
	AddPrivilegeReq struct {
		PriName    string `json:"pri_name" binding:"required"`
		ActionName string `json:"action_name"`
		MenuName   string `json:"menu_name"`
		Pid        int    `json:"pid"`
	}

	UpDatePrivilegeReq struct {
		Id         int    `json:"id" binding:"required"`
		PriName    string `json:"pri_name" binding:"required"`
		ActionName string `json:"action_name"`
		MenuName   string `json:"menu_name"`
		ParentId   int    `json:"parent_id"`
	}

	DelPrivilegeReq struct {
		Id int `form:"id" binding:"required"`
	}
)
