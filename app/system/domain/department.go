package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
)

var (
	DepartmentStatusActive   = 0
	DepartmentStatusDeactive = 1
)

// Department 部门 model
type Department struct {
	core.Model

	DepartmentBase

	CreateBy string `json:"createBy" form:"createBy" gorm:"type:varchar(64);default:'';comment:创建者"` // 创建者
	UpdateBy string `json:"updateBy" form:"updateBy" gorm:"type:varchar(64);default:'';comment:更新者"` // 更新者
}

// TableName 指定表名
//func (m *Department) TableName() string {
//	return "sys_dept"
//}

// DepartmentBase 部门用户基础结构
type DepartmentBase struct {
	ParentID  uint   `json:"parentID" form:"parentID" gorm:"default:0;comment:父部门ID"`                            // 父部门ID
	Name      string `json:"name" form:"name" gorm:"type:varchar(30);not null;comment:部门名称" validate:"required"` // 部门名称
	Ancestors string `json:"ancestors" form:"ancestors" gorm:"type:varchar(255);default:'';comment:祖级列表"`        // 祖级列表
	Order     uint   `json:"order" form:"order" gorm:"type:int(11);default:0;comment:显示顺序"`                      // 显示顺序

	Leader uint `json:"leader" form:"leader" gorm:"type:int(10);default:0;comment:负责人"`              // 负责人 对应用户表中用户
	Status uint `json:"status" form:"status" gorm:"type:tinyint(3);default:0;comment:部门状态（0正常 1停用）"` // 部门状态（0 正常 / 1 停用）
}

// DepartmentSearchRequest 部门查询请求结构体
type DepartmentSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword   string `form:"keyword"`   // 关键字
	ParentID  *uint  `form:"parentID"`  // 父部门ID
	Name      string `form:"name"`      // 部门名称
	Ancestors string `form:"ancestors"` // 祖级列表
	CreateBy  string `form:"createBy"`  // 创建者
	UpdateBy  string `form:"updateBy"`  // 更新者
	Status    *uint  `form:"status"`    // 部门状态（0正常 1停用）

	IDNotIn    []uint `json:"idNotIn"`
	IDIn       []uint `json:"idIn"`
	ParentIDIn []uint `json:"parentIDIn"`
	// 不包含的ID列表
}

// DepartmentCreateRequest 部门添加请求结构体
type DepartmentCreateRequest struct {
	DepartmentBase
}

// DepartmentUpdateRequest 部门更新请求结构体
type DepartmentUpdateRequest struct {
	request.GetById
	DepartmentBase
}
