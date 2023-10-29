package domain

import (
	"ruoyi-go/common/model/core"
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
	return "sys_login_log"
}

/*
  `login_name` varchar(50) DEFAULT '' COMMENT '登录账号',
  `ipaddr` varchar(50) DEFAULT '' COMMENT '登录IP地址',
  `login_location` varchar(255) DEFAULT '' COMMENT '登录地点',
  `browser` varchar(50) DEFAULT '' COMMENT '浏览器类型',
  `os` varchar(50) DEFAULT '' COMMENT '操作系统',
  `status` tinyint(4) DEFAULT '0' COMMENT '登录状态（0成功 1失败）',
  `msg` varchar(255) DEFAULT '' COMMENT '提示消息',
  `login_time` bigint(20) DEFAULT '0' COMMENT '访问时间',
  `module` varchar(30) DEFAULT NULL COMMENT '登录模块',
*/

type LoginLogBase struct {
	Name  string `json:"name" form:"name" validate:"required" gorm:"column:config_name;type:varchar(100);not null;comment:参数名称"` // 参数名称
	Key   string `json:"key" form:"key" validate:"required" gorm:"column:config_key;type:varchar(100);not null;comment:参数键名"`    // 参数键名
	Value string `json:"value" form:"value" gorm:"column:config_value;type:varchar(500);default:'';comment:参数键值"`                // 参数键名

	Type string `json:"type" form:"type" gorm:"column:config_type;type:char(1);default:'N';comment:系统内置（Y是 | N否）"` // 是否系统内置（Y是 | N否）

	core.RuoyiModel
}

// LoginLogSearchRequest 后台访问日志搜索请求结构体
type LoginLogSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword string `json:"keyword" form:"keyword"` // 关键字
	Value   string `json:"value" form:"value"`     // 参数键值
	Key     string `json:"key" form:"key"`         // 参数键名
	Name    string `json:"name" form:"name"`       // 参数名称
	ID      *uint  `json:"id" form:"id"`           // 参数编号
	Type    string `json:"type" form:"type"`       // 是否系统内置（Y是 | N否）

	IDNotIn []uint `json:"idNotIn"` // 不包含的ID列表
	IDIn    []uint `json:"idIn"`    // 包含的ID列表

	core.RouyiSearchRequest
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
