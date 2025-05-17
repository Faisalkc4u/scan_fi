package models

import (
	"time"

	"gorm.io/gorm"
)

type Ingredient struct {
	ID           uint           `gorm:"primaryKey" json:"-"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index" json:"-"`
	Name         *string        `json:"name,omitempty"`                           // Nullable
	IsHarmful    *bool          `gorm:"default:null" json:"is_harmful,omitempty"` // Nullable
	Allergen     *bool          `gorm:"default:null" json:"allergen,omitempty"`   // Nullable
	AllergenNote *string        `json:"allergen_note,omitempty"`                  // Nullable
}
