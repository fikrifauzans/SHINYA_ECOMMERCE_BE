package product

import (
	"time"

	"gorm.io/gorm"
)

type Product struct {
	ID          uint    `gorm:"primaryKey" json:"id"`
	Name        string  `json:"name" binding:"required"`
	Description string  `json:"description" binding:"required, max=255"`
	SKU         string  `json:"sku" binding:"required"`
	Price       float64 `json:"price" binding:"required"`

	CreatedBy uint           `json:"created_by"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedBy uint           `json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
