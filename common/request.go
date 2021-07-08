package common

// 统一的翻页请求
type PageInfo struct {
	Page     int `json:"page" form:"page"`         // 页码
	PageSize int `json:"pageSize" form:"pageSize"` // 每页大小
}

// 统一的排序请求
type SortInfo struct {
	OrderBy string `json:"orderBy" form:"orderBy"`// 排序
}

// 单个ID请求
type GetById struct {
	ID uint `json:"id" form:"id"` // 主键ID
}

// 批量ID请求
type IdsRequest struct {
	Ids []string `json:"ids" form:"ids"`
}