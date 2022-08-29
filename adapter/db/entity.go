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
	Roles        []Role         `gorm:"many2many:role_user;"`
}

type ProductCategory struct {
	ID          uint   `gorm:"primarykey"`
	Name        string `gorm:"type:varchar(250)"`
	Description string
	CreatedAt   time.Time      `json:"created_at"`
	UpdatedAt   time.Time      `json:"updated_at"`
	DeletedAt   gorm.DeletedAt `gorm:"index"`
}

type Product struct {
	ID                uint   `gorm:"primarykey"`
	Name              string `gorm:"type:varchar(250)"`
	Description       string
	ProductCategoryID uint
	ProductCategory   ProductCategory
	ProductVariations []ProductVariation
	CreatedAt         time.Time      `json:"created_at"`
	UpdatedAt         time.Time      `json:"updated_at"`
	DeletedAt         gorm.DeletedAt `gorm:"index"`
}

type ProductVariation struct {
	ID        uint   `gorm:"primarykey"`
	Name      string `gorm:"type:varchar(250)"`
	ProductID uint
	Product   Product
	Price     uint
	Stock     uint
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

type PostCategory struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"type:varchar(250)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Post struct {
	ID        uint   `gorm:"primarykey"`
	Title     string `gorm:"type:varchar(250)"`
	content   string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Comment struct {
	ID        uint `gorm:"primarykey"`
	UserID    int
	User      User
	Comment   string `gorm:"type:text"`
	ModelID   int
	ModelType string
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type Role struct {
	ID          uint         `gorm:"primarykey"`
	Name        string       `gorm:"type:varchar(250)"`
	CreatedAt   time.Time    `json:"created_at"`
	UpdatedAt   time.Time    `json:"updated_at"`
	Permissions []Permission `gorm:"many2many:permission_role;"`
	Users       []User       `gorm:"many2many:role_user;"`
}

type Permission struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"type:varchar(250)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
	Roles     []Role    `gorm:"many2many:permission_role;"`
}

type Province struct {
	ID        uint      `gorm:"primarykey"`
	Name      string    `gorm:"type:varchar(250)"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

type City struct {
	ID         uint `gorm:"primarykey"`
	ProvinceID int
	Name       string    `gorm:"type:varchar(250)"`
	CreatedAt  time.Time `json:"created_at"`
	UpdatedAt  time.Time `json:"updated_at"`
	Province   Province
}

type Address struct {
	ID          uint `gorm:"primarykey"`
	CityID      int
	Name        string    `gorm:"type:varchar(250)"`
	Description string    `gorm:"type:text"`
	CreatedAt   time.Time `json:"created_at"`
	UpdatedAt   time.Time `json:"updated_at"`
	City        City
}

//
//type Transaction struct {
//	ID        uint   `gorm:"primarykey"`
//	Title     string `gorm:"type:varchar(250)"`
//	OrderID   uint
//	Order
//	CreatedAt time.Time `json:"created_at"`
//	UpdatedAt time.Time `json:"updated_at"`
//}

type AccessToken struct {
	ID         uint `gorm:"primarykey"`
	UserID     uint
	User       User
	Title      string `gorm:"type:varchar(250)"`
	Token      string `gorm:"type:varchar(250)"`
	ExpiryDate time.Time
	CreatedAt  time.Time
	UpdatedAt  time.Time
}
