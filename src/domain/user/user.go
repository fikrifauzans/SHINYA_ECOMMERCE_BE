package user

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	ID        uint           `gorm:"primaryKey" json:"id"`
	Name      string         `json:"name" binding:"required"`
	Email     string         `json:"email" gorm:"unique" binding:"required"`
	Password  string         `json:"password" binding:"required"`
	BirthDate *time.Time     `json:"birth_date" form:"birth_date" time_format:"2006-01-02"`
	Address   string         `json:"address" binding:"required"`
	CreatedBy uint           `json:"created_by"`
	UpdatedBy uint           `json:"updated_by"`
	DeletedBy uint           `json:"deleted_by"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at" gorm:"index"`
}
