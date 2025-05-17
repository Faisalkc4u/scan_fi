package models

import (
	"time"

	"gorm.io/gorm"
)

type Additive struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Code        *string `gorm:"uniqueIndex" json:"code,omitempty"`        // Nullable
	CommonName  *string `json:"common_name,omitempty"`                    // Nullable
	Category    *string `json:"category,omitempty"`                       // Nullable
	IsHarmful   *bool   `gorm:"default:null" json:"is_harmful,omitempty"` // Nullable bool
	Description *string `json:"description,omitempty"`                    // Nullable
}
