package models

import (
	"time"

	"gorm.io/gorm"
)

type ProductAdditive struct {
	ID         uint           `gorm:"primaryKey" json:"-"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index" json:"-"`
	ProductID  uint           `json:"product_id"`
	AdditiveID uint           `json:"additive_id"`
	Percentage float64        `gorm:"default:0" json:"percentage"` // Optional
}
