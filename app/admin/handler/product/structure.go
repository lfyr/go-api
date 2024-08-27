package product

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
