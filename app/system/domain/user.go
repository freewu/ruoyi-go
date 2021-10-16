package domain

import (
	"ruoyi-go/common/core"
	"ruoyi-go/common/model/request"
)

// 后台用户 model
type User struct {
	core.Model

	UserBase
}

// 后台用户基础结构
type UserBase struct {
	Username	string	`json:"username" gorm:"type:varchar(60);not null;unique;comment:用户名" validate:"required"` // 用户名
	Mobile		string	`json:"mobile" gorm:"type:varchar(20);not null;unique;comment:手机号" validate:"required"` // 中国手机不带国家代码，国际手机号格式为：国家代码-手机号
	Nickname	string	`json:"username" gorm:"type:varchar(50);not null;comment:用户昵称" validate:"required"` // 用户昵称
}

/*
CREATE TABLE `user` (
  `birthday` int(11) NOT NULL DEFAULT '0' COMMENT '生日',
  `create_time` int(11) NOT NULL DEFAULT '0' COMMENT '注册时间',
  `user_password` varchar(255) NOT NULL DEFAULT '' COMMENT '登录密码;cmf_password加密',
  `user_status` tinyint(3) unsigned NOT NULL DEFAULT '1' COMMENT '用户状态;0:禁用,1:正常,2:未验证',
  `user_email` varchar(100) NOT NULL DEFAULT '' COMMENT '用户登录邮箱',
  `sex` tinyint(4) NOT NULL DEFAULT '0' COMMENT '性别;0:保密,1:男,2:女',
  `avatar` varchar(255) NOT NULL DEFAULT '' COMMENT '用户头像',
  `last_login_time` int(11) NOT NULL DEFAULT '0' COMMENT '最后登录时间',
  `last_login_ip` varchar(15) NOT NULL DEFAULT '' COMMENT '最后登录ip',
  `dept_id` bigint(20) unsigned NOT NULL DEFAULT '0' COMMENT '部门id',
  `remark` varchar(255) NOT NULL DEFAULT '' COMMENT '备注',
  `is_admin` tinyint(4) NOT NULL DEFAULT '1' COMMENT '是否后台管理员 1 是  0   否',
  PRIMARY KEY (`id`) USING BTREE,
  UNIQUE KEY `user_login` (`user_name`) USING BTREE,
  UNIQUE KEY `mobile` (`mobile`) USING BTREE,
  KEY `user_nickname` (`user_nickname`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=2 DEFAULT CHARSET=utf8mb4 COMMENT='用户表';
*/

// 后台用户搜索请求结构体
type UserSearchRequest struct {
	request.PageInfo
	request.SortInfo

	Keyword	string 	`json:"keyword"` // 关键字
}

type UserAddRequest struct {
	UserBase
}

type UserEditRequest struct {
	request.GetById
	UserBase
}