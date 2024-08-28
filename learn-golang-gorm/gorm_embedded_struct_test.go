package learn_golang_gorm

import (
	"time"

	"gorm.io/gorm"
)

type User struct {
	gorm.Model        // model struct
	Name         Name `gorm:"embedded"`
	Email        string
	Password     string
	Wallet       Wallet    `gorm:"foreignKey:user_id; references:id"`                                                                     // one to one, has one
	Address      []Address `gorm:"foreignKey:user_id;references:id"`                                                                      // one to many, has many
	LikeProducts []Product `gorm:"many2many:like_products;foreignKey:id;joinForeignKey:user_id; references:id;joinReferences:product_id"` // many to many
}

type Name struct {
	FirstName string
	LastName  string
}

type DateAt struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	DeletedAt gorm.DeletedAt
}
