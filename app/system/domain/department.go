package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
)

// 部门 model
type Department struct {
	core.Model

	DepartmentBase

	//DeleteFlag	uint	`json:"deleteFlag" form:"deleteFlag" gorm:"type:tinyint(3);default:0;comment:删除标志（0代表存在 2代表删除）"` // 删除标志（0代表存在 2代表删除）
	CreateBy	string	`json:"createBy" form:"createBy" gorm:"type:varchar(64);default:'';comment:创建者"` // 创建者
	UpdateBy	string	`json:"updateBy" form:"updateBy" gorm:"type:varchar(64);default:'';comment:更新者"` // 更新者
}

type DepartmentBase struct {
	ParentID	uint	`json:"parentID" form:"parentID" gorm:"default:0;comment:父部门ID"` // 父部门ID
	Name		string	`json:"name" form:"name" gorm:"type:varchar(30);not null;comment:部门名称" validate:"required"` // 部门名称
	Ancestors	string	`json:"ancestors" form:"ancestors" gorm:"type:varchar(50);default:'';comment:祖级列表"` // 祖级列表
	Order		uint	`json:"order" form:"order" gorm:"type:int(11);default:0;comment:显示顺序"` // 显示顺序

	Leader		string	`json:"leader" form:"leader" gorm:"type:varchar(20);default:'';comment:负责人"` // 负责人
	Mobile		string	`json:"mobile" form:"mobile" gorm:"type:varchar(20);default:'';comment:联系电话"` // 联系电话
	Email		string	`json:"email" form:"email" gorm:"type:varchar(50);default:'';comment:邮箱"` // 邮箱

	Status      uint	`json:"status" form:"status" gorm:"type:tinyint(3);default:0;comment:部门状态（0正常 1停用）"` // 部门状态（0正常 1停用）
}

type DepartmentSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword		string 	`form:"keyword"` // 关键字
	ParentID	*uint 	`form:"parentID"` // 父部门ID
	Name		string 	`form:"name"` // 部门名称
	Ancestors	string 	`form:"ancestors"` // 祖级列表
	Leader		string 	`form:"leader"` // 负责人
	Mobile		string 	`form:"mobile"` // 联系电话
	Email		string 	`form:"email"` // 邮箱
	CreateBy	string	`form:"createBy"` // 创建者
	UpdateBy	string	`form:"updateBy"` // 更新者
	Status		*uint	`form:"status"` // 部门状态（0正常 1停用）
}

type DepartmentAddRequest struct {
	DepartmentBase
}

type DepartmentEditRequest struct {
	request.GetById
	DepartmentBase
}