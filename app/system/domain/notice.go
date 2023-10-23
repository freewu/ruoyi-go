package domain

import (
	"ruoyi-go/common/model/core"
	"ruoyi-go/common/model/request"
)

// Notice 公告 model
type Notice struct {
	//core.Model
	ID uint `json:"id" form:"id" gorm:"column:notice_id;type:bigint(20) auto_increment;primary_key;not null;comment:公告ID"` // 公告ID

	NoticeBase
}

// TableName 指定表名
func (m *Notice) TableName() string {
	return "sys_notice"
}

type NoticeBase struct {
	Title   string `json:"title" form:"title" validate:"required" gorm:"column:notice_title;type:varchar(50);not null;comment:公告标题"` // 公告标题
	Content string `json:"content" form:"content" validate:"required" gorm:"column:notice_content;type:longblob;comment:公告内容"`       // 公告内容
	Type    uint   `json:"type" form:"type" gorm:"column:notice_type;type:char(1);not null;default:0;comment:公告类型（1通知 2公告）"`         // 公告类型（1通知 2公告）                         // 显示顺序
	Status  uint   `json:"status" form:"status" gorm:"column:status;type:char(1);not null;comment:公告状态（0正常 1关闭）"`                    // 公告状态（0正常 1关闭）

	core.RuoyiModel
}

// NoticeSearchRequest 后台公告搜索请求结构体
type NoticeSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword string `json:"keyword" form:"keyword"` // 关键字
	Title   string `json:"title" form:"title"`     // 公告标题
	Content string `json:"content" form:"content"` // 公告内容
	ID      *uint  `json:"id" form:"id"`           // 公告编号
	Status  *uint  `json:"status" form:"status"`   // 公告状态（0正常 1关闭）
	Type    *uint  `json:"type" form:"type"`       // 公告类型（1通知 2公告）

	IDNotIn []uint `json:"idNotIn"` // 不包含的ID列表
	IDIn    []uint `json:"idIn"`    // 包含的ID列表

	core.RouyiSearchRequest
}

// NoticeCreateRequest 公告添加请求结构体
type NoticeCreateRequest struct {
	NoticeBase
}

// NoticeUpdateRequest 公告编辑请求结构体
type NoticeUpdateRequest struct {
	request.GetById
	NoticeBase
}
