package core

import (
	"gorm.io/gorm"
	"time"
)

type Model struct {
	ID        uint 				`json:"id" form:"id" gorm:"primarykey;comment:自增ID"`
	CreatedAt time.Time			`json:"createdAt" form:"createdAt" gorm:"comment:创建时间"`
	UpdatedAt time.Time			`json:"updatedAt" form:"updatedAt" gorm:"comment:更新时间"`
	DeletedAt gorm.DeletedAt	`json:"-" gorm:"index"`
}
