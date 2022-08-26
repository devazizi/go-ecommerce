package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID         uint           `gorm:"primarykey" json:"id"`
	Name       string         `json:"name"`
	CellNumber string         `json:"cell_number"`
	CreatedAt  time.Time      `json:"created_at"`
	UpdatedAt  time.Time      `json:"updated_at"`
	DeletedAt  gorm.DeletedAt `gorm:"index"`
}
