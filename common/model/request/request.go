package request

// Paging common input parameter structure
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

type SortInfo struct {
	OrderBy string `json:"orderBy" form:"orderBy"`// 排序
}

// Find by id structure
type GetById struct {
	ID  uint `json:"id" form:"id" validate:"required"` // 主键ID
}

type IdsRequest struct {
	Ids []uint `json:"ids" form:"ids" validate:"required"`
}

type Empty struct{}