package models

import (
	"time"

	"gorm.io/gorm"
)

type Manufacturer struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name *string `gorm:"uniqueIndex" json:"name,omitempty"`

	// Self-referencing structure
	ParentID     *uint          `json:"parent_id,omitempty"`                               // Optional parent manufacturer
	Parent       *Manufacturer  `gorm:"foreignKey:ParentID" json:"parent,omitempty"`       // Direct parent
	Subsidiaries []Manufacturer `gorm:"foreignKey:ParentID" json:"subsidiaries,omitempty"` // Brands or child manufacturers
}
