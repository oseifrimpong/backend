package model

import (
	"time"

	"gorm.io/gorm"
)

type Base struct {
	Version   int32          `json:"version" gorm:"version;"`
	CreatedAt *time.Time     `json:"created_at" gorm:"column:created_at"`
	UpdatedAt *time.Time     `json:"updated_at" gorm:"column:updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" swaggertype:"integer"`
	ID        int            `json:"id" gorm:"column:id;primary_key;type:varchar(30);"`
}
