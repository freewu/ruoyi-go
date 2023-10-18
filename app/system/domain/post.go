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

/*
CREATE TABLE `sys_post` (
  `post_id` bigint(20) NOT NULL AUTO_INCREMENT COMMENT '岗位ID',
  `post_code` varchar(64) NOT NULL COMMENT '岗位编码',
  `post_name` varchar(50) NOT NULL COMMENT '岗位名称',
  `post_sort` int(11) NOT NULL COMMENT '显示顺序',
  `status` char(1) NOT NULL COMMENT '状态（0正常 1停用）',
  `create_by` varchar(64) DEFAULT '' COMMENT '创建者',
  `create_time` datetime DEFAULT NULL COMMENT '创建时间',
  `update_by` varchar(64) DEFAULT '' COMMENT '更新者',
  `update_time` datetime DEFAULT NULL COMMENT '更新时间',
  `remark` varchar(500) DEFAULT NULL COMMENT '备注',
  PRIMARY KEY (`post_id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='岗位信息表';
*/

// 后台职位搜索请求结构体
type PostSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword string `form:"keyword"` // 关键字
}

type PostAddRequest struct {
	PostBase
}

type PostEditRequest struct {
	request.GetById
	PostBase
}
