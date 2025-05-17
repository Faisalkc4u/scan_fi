package models

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID        uint           `gorm:"primaryKey" json:"-"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`

	Name           string `gorm:"index" json:"name"`
	Barcode        string `gorm:"uniqueIndex:idx_barcode_deleted_at" json:"barcode"`
	NumberOfSearch int    `gorm:"default:0" json:"number_of_search"`

	// Relations
	BrandID *uint         `json:"brand_id,omitempty"`
	Brand   *Manufacturer `gorm:"foreignKey:BrandID" json:"brand,omitempty"`

	Ingredients []Ingredient `gorm:"many2many:product_ingredients;" json:"ingredients,omitempty"`
	Additives   []Additive   `gorm:"many2many:product_additives;" json:"additives,omitempty"`
}
