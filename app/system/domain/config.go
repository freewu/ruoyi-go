package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
)

// Config 系统配置 model
type Config struct {
	//core.Model
	ID uint `json:"id" form:"id" gorm:"column:config_id;type:bigint(20) auto_increment;primary_key;not null;comment:参数主键"` // 参数主键

	ConfigBase
}

// TableName 指定表名
func (m *Config) TableName() string {
	return "sys_config"
}

/*
  `config_name` varchar(100) DEFAULT '' COMMENT '参数名称',
  `config_key` varchar(100) DEFAULT '' COMMENT '参数键名',
  `config_value` varchar(500) DEFAULT '' COMMENT '参数键值',
  `config_type` char(1) DEFAULT 'N' COMMENT '系统内置（Y是 N否）',
*/

type ConfigBase struct {
	Name  string `json:"name" form:"name" validate:"required" gorm:"column:config_name;type:varchar(100);not null;comment:参数名称"` // 参数名称
	Key   string `json:"key" form:"key" validate:"required" gorm:"column:config_key;type:varchar(100);not null;comment:参数键名"`    // 参数键名
	Value string `json:"value" form:"value" gorm:"column:config_value;type:varchar(500);default:'';comment:参数键值"`                // 参数键名

	Type string `json:"type" form:"type" gorm:"column:config_type;type:char(1);default:'N';comment:系统内置（Y是 | N否）"` // 是否系统内置（Y是 | N否）

	core.RuoyiModel
}

// ConfigSearchRequest 后台系统配置搜索请求结构体
type ConfigSearchRequest struct {
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

// ConfigCreateRequest 系统配置添加请求结构体
type ConfigCreateRequest struct {
	ConfigBase
}

// ConfigUpdateRequest 系统配置编辑请求结构体
type ConfigUpdateRequest struct {
	request.GetById
	ConfigBase
}
