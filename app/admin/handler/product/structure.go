package product

type (
	IdReq struct {
		Id int `json:"id" form:"id" binding:"required"`
	}
)

// User
type (
	GetBrandReq struct {
		Page     int `form:"page" binding:"required"`
		PageSize int `form:"pageSize" binding:"required"`
	}
	AddBrandReq struct {
		BrandName string `json:"brandName" form:"brandName" binding:"required"`
		Logo      string `json:"logo" form:"logo"`
	}
	UpdateBrandReq struct {
		Id        int    `json:"id" form:"id" binding:"required"`
		BrandName string `json:"brandName" form:"brandName" binding:"required"`
		Logo      string `json:"logo" form:"logo"`
	}
)

type (
	AddCategoryReq struct {
		CatName  string `json:"cat_name" form:"cat_name" binding:"required"`
		ParentId int    `json:"parent_id" form:"parent_id"`
	}
	UpdateCategoryReq struct {
		Id       int    `json:"id" form:"id"`
		CatName  string `json:"cat_name" form:"cat_name" binding:"required"`
		ParentId int    `json:"parent_id" form:"parent_id"`
	}
)
