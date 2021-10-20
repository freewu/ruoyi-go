package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
	"time"
)

// 后台用户 model
type User struct {
	core.Model

	UserBase

	RegisterTime 	time.Time	`json:"registerTime" form:"registerTime" gorm:"default:null;comment:注册时间"` // 注册时间
	Password     	string		`json:"-" form:"password" gorm:"type:varchar(255);default:null;comment:登录密码"` // 登录密码;cmf_password加密'
	LastLoginTime 	time.Time	`json:"lastLoginTime" form:"lastLoginTime" gorm:"default:null;comment:最后登录时间"`  // 最后登录时间
	LastLoginIP   	string		`json:"lastLoginIP" form:"lastLoginIP" gorm:"type:varchar(15);default:'';comment:最后登录IP"`  // 最后登录IP
}

// 后台用户基础结构
type UserBase struct {
	Username    	string          `json:"username" form:"username" gorm:"type:varchar(60);not null;unique;comment:用户名" validate:"required"` // 用户名
	Mobile      	string          `json:"mobile" form:"mobile" gorm:"type:varchar(20);not null;unique;comment:手机号" validate:"required"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Nickname     	string          `json:"nickname" form:"nickname" gorm:"type:varchar(50);not null;comment:用户昵称" validate:"required"`// 用户昵称
	//Birthday     	core.LocalDate 	`json:"birthday" form:"birthday" time_format:"2006-01-02" time_utc:"1" gorm:"default:null;comment:生日"` // 生日
	Birthday     	string  		`json:"birthday" form:"birthday" gorm:"default:null;comment:生日"` // 生日

	Status       	uint            `json:"status" form:"status"  gorm:"type:tinyint(3);default:1;comment:用户状态;0:禁用,1:正常,2:未验证"` // 用户状态;0:禁用,1:正常,2:未验证
	Email        	string          `json:"email" form:"email" gorm:"type:varchar(100);default:'';comment:登录邮箱"` // 登录邮箱
	Gender        	uint			`json:"gender" form:"gender" gorm:"type:tinyint(4);default:0;comment:性别;0:保密,1:男,2:女"` // 性别;0:保密,1:男,2:女
	Avatar        	string			`json:"avatar" form:"avatar" gorm:"type:varchar(255);default:'';comment:用户头像"` // 用户头像
	DepartmentID  	uint			`json:"departmentID" form:"departmentID" gorm:"type:bigint(20);default:0;comment:部门ID"`  // 部门ID
	Remark			string			`json:"remark" form:"remark" gorm:"type:varchar(255);default:'';comment:备注"` // 备注
	IsAdmin			bool			`json:"isAdmin" form:"isAdmin" gorm:"default:0;comment:是否后台管理员 1 是 | 0 否"` // 是否后台管理员 1 是  0 否
}

// 后台用户搜索请求结构体
type UserSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword			string 	`form:"keyword"` // 关键字
	Username		string	`form:"username"` // 用户名
	Mobile			string	`form:"mobile"` // 手机号
	Email			string	`form:"email"` // 登录邮箱
	Nickname		string	`form:"nickname"` // 用户昵称
	Remark			string	`form:"remark"` // 备注
	LastLoginIP		string	`form:"lastLoginIP"` // 最后登录IP
	Status			*uint	`form:"status"` // 用户状态;0:禁用,1:正常,2:未验证
	Gender			*uint	`form:"gender"` // 性别;0:保密,1:男,2:女
	IsAdmin			*bool	`form:"isAdmin"` // 是否后台管理员 1 是  0 否
	DepartmentID  	*uint	`form:"departmentID"` // 部门ID

	DateColumn		string 	`form:"dateColumn"` // 搜索时间字段 lastLoginTime 最后登录时间 | registerTime 注册时间
	BeginTime 		string 	`form:"beginTime" validate:"omitempty,datetime-format"` // 开始时间 格式:yyyy-mm-dd hh:ii:ss
	EndTime 		string 	`form:"endTime" validate:"omitempty,datetime-format"` // 结束时间 格式:yyyy-mm-dd hh:ii:ss
}

type UserAddRequest struct {
	UserBase
}

type UserEditRequest struct {
	request.GetById
	UserBase
}