package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
)

// 职位 model
type Post struct {
	//core.Model
	ID uint `json:"id" form:"id" gorm:"column:post_id;type:bigint(20);primary_key;autoIncrement;not null;comment:岗位ID"` // 岗位ID

	PostBase
}

// TableName 指定表名
func (m *Post) TableName() string {
	return "sys_post"
}

type PostBase struct {
	Code   string `json:"code" form:"code" validate:"required" gorm:"column:post_code;type:varchar(64);not null;comment:岗位编码"` // 岗位编码
	Name   string `json:"name" form:"name" validate:"required" gorm:"column:post_name;type:varchar(50);not null;comment:岗位名称"` // 岗位名称
	Sort   uint   `json:"sort" form:"sort" gorm:"column:post_sort;type:int(11);not null;comment:显示顺序"`                         // 显示顺序
	Status uint   `json:"status" form:"status" gorm:"column:status;type:char(1);not null;comment:状态(0正常 1停用)"`                 // 状态（0正常 1停用）

	core.RuoyiModel
}

// PostSearchRequest 后台职位搜索请求结构体
type PostSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword string `json:"keyword" form:"keyword"` // 关键字
	Code    string `json:"code" form:"code"`       // 岗位编码
	Name    string `json:"name" form:"name"`       // 岗位编码
	ID      *uint  `json:"id" form:"id"`           // 岗位编号
	Status  *uint  `json:"status" form:"status"`   // 状态

	IDNotIn []uint `json:"idNotIn"` // 不包含的ID列表
	IDIn    []uint `json:"idIn"`    // 包含的ID列表

	core.RouyiSearchRequest
}

// PostCreateRequest 职位添加请求结构体
type PostCreateRequest struct {
	PostBase
}

// PostUpdateRequest 职位编辑请求结构体
type PostUpdateRequest struct {
	request.GetById
	PostBase
}
