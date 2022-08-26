package db

import (
	"gorm.io/gorm"
	"time"
)

type User struct {
	ID           uint           `gorm:"primarykey" json:"id"`
	Name         string         `json:"name" gorm:"type:varchar(250);"`
	Email        string         `json:"email" gorm:"type:varchar(250);uniqueIndex"`
	NationalCode string         `json:"national_code" gorm:"type:varchar(20);uniqueIndex"`
	CellNumber   string         `json:"cell_number" gorm:"type:varchar(25);uniqueIndex"`
	Password     string         `json:"password" gorm:"type:varchar(250);"`
	CreatedAt    time.Time      `json:"created_at"`
	UpdatedAt    time.Time      `json:"updated_at"`
	DeletedAt    gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(250)"`
	Description string
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type ProductVariation struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(250)"`
	Price     int
	Stock     int
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index"`
}

type ShoppingCart struct {
	ID        uint `gorm:"primarykey"`
	UserID    int
	User      User
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type ShoppingCartItem struct {
	ID             uint `gorm:"primarykey"`
	ShoppingCartID int
	ShoppingCart   ShoppingCart
	Quantity       int
	CreatedAt      time.Time `json:"created_at"`
	UpdatedAt      time.Time `json:"updated_at"`
}

type Order struct {
	ID         uint `gorm:"primarykey"`
	TotalPrice int
	UserID     int
	User       User
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type OrderItem struct {
	ID         uint `gorm:"primarykey"`
	Quantity   int
	TotalPrice int
	OrderID    int
	Order      Order
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type Refund struct {
	ID         uint `gorm:"primarykey"`
	TotalPrice int
	OrderID    int
	Order      Order
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}

type RefundItem struct {
	ID         uint `gorm:"primarykey"`
	Quantity   int
	TotalPrice int
	RefundID   int
	Refund     Refund
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
}
