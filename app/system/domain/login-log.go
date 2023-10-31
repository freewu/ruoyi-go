package domain

import (
	"ruoyi-go/common/model/request"
)

// LoginLog 访问日志 model
type LoginLog struct {
	//core.Model
	ID uint `json:"id" form:"id" gorm:"column:info_id;type:bigint(20) auto_increment;primary_key;not null;comment:访问ID"` // 访问ID

	LoginLogBase
}

// TableName 指定表名
func (m *LoginLog) TableName() string {
	return "sys_logininfor"
}

type LoginLogBase struct {
	Name      string `json:"login_name" form:"login_name" gorm:"column:login_name;type:varchar(50);default:'';comment:登录账号"`             // 登录账号
	IPAddress string `json:"ipaddr" form:"ipaddr" gorm:"column:ipaddr;type:varchar(50);default:'';comment:登录IP地址"`                       // 登录IP地址
	Location  string `json:"login_location" form:"login_location" gorm:"column:login_location;type:varchar(50);default:'';comment:登录地点"` // 登录地点
	Browser   string `json:"browser" form:"browser" gorm:"column:browser;type:varchar(50);default:'';comment:浏览器类型"`                     // 浏览器类型
	OS        string `json:"os" form:"os" gorm:"column:os;type:varchar(50);default:'';comment:操作系统"`                                     // 操作系统
	Message   string `json:"msg" form:"msg" gorm:"column:msg;type:varchar(50);default:'';comment:提示消息"`                                  // 提示消息
	Module    string `json:"module" form:"module" gorm:"column:module;type:varchar(30);default:'';comment:登录模块"`                         // 登录模块
	Time      string `json:"login_time" form:"login_time" gorm:"column:login_time;type:bigint(20);default:'0';comment:访问时间 13位时间戳"`      // 访问时间

	Status uint `json:"status" form:"status" gorm:"column:status;type:tinyint(4);default:0;comment:登录状态（0成功 1失败）"` // 登录状态（0成功 1失败）
}

// LoginLogSearchRequest 后台访问日志搜索请求结构体
type LoginLogSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword   string `json:"keyword" form:"keyword"`               // 关键字
	Name      string `json:"login_name" form:"login_name"`         // 登录账号
	IPAddress string `json:"ipaddr" form:"ipaddr"`                 // 登录IP地址
	Location  string `json:"login_location" form:"login_location"` // 登录地点
	Browser   string `json:"browser" form:"browser"`               // 浏览器类型
	OS        string `json:"os" form:"os"`                         // 操作系统
	Message   string `json:"msg" form:"msg"`                       // 提示消息
	Module    string `json:"module" form:"module"`                 // 登录模块
	ID        *uint  `json:"id" form:"id"`                         // 访问ID
	Status    *uint  `json:"status" form:"status"`                 // 登录状态（0成功 1失败）

	BeginDate string `json:"begin_date" form:"begin_date"` // 开始时间
	EndDate   string `json:"end_date" form:"end_date"`     // 结束时间

	IDNotIn []uint `json:"idNotIn"` // 不包含的ID列表
	IDIn    []uint `json:"idIn"`    // 包含的ID列表

}

// LoginLogCreateRequest 访问日志添加请求结构体
type LoginLogCreateRequest struct {
	LoginLogBase
}

// LoginLogUpdateRequest 访问日志编辑请求结构体
type LoginLogUpdateRequest struct {
	request.GetById
	LoginLogBase
}

/**
CREATE TABLE `sys_login_log` (
  `info_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '访问ID',
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` bigint(20) DEFAULT '0' COMMENT '访问时间',
  `module` varchar(30) DEFAULT NULL COMMENT '登录模块',
  PRIMARY KEY (`info_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=96 DEFAULT CHARSET=utf8mb4 COMMENT='系统访问记录';
*/
