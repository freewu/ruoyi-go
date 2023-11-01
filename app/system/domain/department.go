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
	//core.Model
	ID uint `json:"id" form:"id" gorm:"column:dept_id;type:bigint(20) auto_increment;primary_key;not null;comment:部门id"` // 部门id

	DepartmentBase
}

// TableName 指定表名
func (m *Department) TableName() string {
	return "sys_dept"
}

// DepartmentBase 部门用户基础结构
type DepartmentBase struct {
	ParentID  uint   `json:"parent_id" form:"parent_id" gorm:"column:parent_id;type:bigint(20);default:0;comment:父部门ID"`                      // 父部门ID
	Ancestors string `json:"ancestors" form:"ancestors" validate:"required" gorm:"column:ancestors;type:varchar(64);default:'';comment:祖级列表"` // 祖级列表
	Name      string `json:"name" form:"name" validate:"required" gorm:"column:dept_name;type:varchar(30);not null;comment:部门名称"`             // 部门名称
	Sort      uint   `json:"sort" form:"sort" gorm:"column:order_num;type:int(4);default:0;comment:显示顺序"`                                     // 显示顺序
	Status    uint   `json:"status" form:"status" gorm:"column:status;type:char(1);default:0;comment:部门状态（0正常 1停用）"`                          // 部门状态（0正常 1停用）

	Leader string `json:"leader" form:"leader" gorm:"column:leader;type:varchar(20);comment:负责人"` // 负责人
	Phone  string `json:"phone" form:"phone" gorm:"column:phone;type:varchar(11);comment:联系电话"`   // 联系电话
	EMail  string `json:"email" form:"email" gorm:"column:email;type:varchar(50);comment:邮箱"`     // 邮箱

	core.RuoyiModel
}

// DepartmentSearchRequest 部门查询请求结构体
type DepartmentSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword   string `form:"keyword"`   // 关键字
	ParentID  *uint  `form:"parentID"`  // 父部门ID
	Name      string `form:"name"`      // 部门名称
	Leader    string `form:"leader"`    // 负责人
	Phone     string `form:"phone"`     // 联系电话
	EMail     string `form:"email"`     // 邮箱
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
