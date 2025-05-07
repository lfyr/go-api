package structure

type (
	IdReq struct {
		Id int `json:"id" form:"id" binding:"required"`
	}
	PageReq struct {
		Page     int `form:"page" binding:"required"`
		PageSize int `form:"pageSize" binding:"required"`
	}
)

// User
type (
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
		CatName  string `json:"catName" form:"catName" binding:"required"`
		ParentId int    `json:"parentId" form:"parentId"`
	}
	UpdateCategoryReq struct {
		Id       int    `json:"id" form:"id" binding:"required"`
		CatName  string `json:"catName" form:"catName"`
		ParentId int    `json:"parentId" form:"parentId"`
	}

	CategoryTree struct {
		Id       int            `json:"id"`
		CatName  string         `json:"cat_name"`
		ParentId int            `json:"parent_id"`
		Children []CategoryTree `json:"children"`
	}
)

type (
	AddGoodsReq struct {
		GoodsName      string      `json:"goodsName"`
		CatId          int         `json:"catId"`
		BrandId        int         `json:"brandId"`
		ShopPrice      float64     `json:"shopPrice"`
		Logo           string      `json:"logo"`
		SmLogo         string      `json:"smLogo"`
		IsHot          int         `json:"isHot"`
		IsNew          int         `json:"isNew"`
		IsBest         int         `json:"isBest"`
		IsOnSale       int         `json:"isOnSale"`
		SeoKeyword     string      `json:"seoKeyWord"`
		SeoDescription string      `json:"seoDescription"`
		TypeId         int         `json:"typeId"`
		SortNum        int         `json:"sortNum"`
		IsDelete       int         `json:"isDelete"`
		GoodsDesc      string      `json:"goodsDesc"`
		Addtime        int         `json:"addtime"`
		GoodsPicsData  []GoodsPics `json:"goods_pics"`
	}

	GoodsPics struct {
		Pic     string `json:"pic"`
		SmPic   string `json:"smPic"`
		GoodsId int    `json:"goodsId"`
	}

	UpdateGoodsReq struct {
		Id             int     `json:"id" form:"id" binding:"required"`
		GoodsName      string  `json:"goodsName"`
		CatId          int     `json:"catId"`
		BrandId        int     `json:"brandId"`
		ShopPrice      float64 `json:"shopPrice"`
		Logo           string  `json:"logo"`
		SmLogo         string  `json:"smLogo"`
		IsHot          int     `json:"isHot"`
		IsNew          int     `json:"isNew"`
		IsBest         int     `json:"isBest"`
		IsOnSale       int     `json:"isOnSale"`
		SeoKeyword     string  `json:"seoKeyWord"`
		SeoDescription string  `json:"seoDescription"`
		TypeId         int     `json:"typeId"`
		SortNum        int     `json:"sortNum"`
		IsDelete       int     `json:"isDelete"`
		GoodsDesc      string  `json:"goodsDesc"`
		Addtime        int     `json:"addtime"`
	}
)
