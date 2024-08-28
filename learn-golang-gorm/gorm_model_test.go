package learn_golang_gorm

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

// use gorm tag & without conventions
type Product struct {
	ID           int       `gorm:"primaryKey;column:id"`
	Title        string    `gorm:"column:title"`
	Description  string    `gorm:"column:description"`
	Price        int64     `gorm:"column:price"`
	CreatedAt    time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt    time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	LikedByUsers []User    `gorm:"many2many:like_products;foreignKey:id;joinForeignKey:product_id;references:id;joinReferences:user_id"` // many to many
}

// with conventions
// type Product struct {
// 	ID          int
// 	Title       string
// 	Description string
// 	Price       int64
// 	CreatedAt   time.Time
// 	UpdatedAt   time.Time
// }

// change default table name, default is "products"
// func (user *Product) TableName() string {
// 	return "products"
// }

func TestModelGorm(t *testing.T) {
	product := Product{
		Title:       "Product-2",
		Description: "Description-2",
		Price:       250000,
	}

	err := Db.Create(&product).Error

	assert.NoError(t, err)
}
