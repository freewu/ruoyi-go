package core

import "time"

// RuoyiModel 若依的公用  Model
type RuoyiModel struct {
	CreatedTime time.Time `json:"create_time" form:"create_time" gorm:"column:create_time;type:datetime;comment:创建时间"` // 创建时间
	CreatedBy   string    `json:"create_by" form:"create_by" gorm:"column:create_by;type:varchar(64);comment:创建者"`     // 创建者
	UpdateTime  time.Time `json:"update_time" form:"update_time" gorm:"column:update_time;type:datetime;comment:更新时间"` // 更新时间
	UpdateBy    string    `json:"update_by" form:"update_by" gorm:"column:update_by;type:varchar(64);comment:更新者"`     // 更新者
	Remark      string    `json:"remark" form:"remark" gorm:"column:remark;type:varchar(500);comment:备注"`              // 备注
}
