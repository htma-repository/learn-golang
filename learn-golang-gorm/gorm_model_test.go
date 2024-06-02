package learn_golang_gorm

import (
	"testing"
	"time"
)

// use gorm tag & without conventions
// type Product struct {
// 	ID          int       `gorm:"not_null;autoIncrement;<-:create"`
// 	Title       string    `gorm:"not_null;size:255"`
// 	Description string    `gorm:"not_null"`
// 	Price       int       `gorm:"not_null"`
// 	CreatedAt   time.Time `gorm:"not_null;autoCreateTime;<-:create"`
// 	UpdatedAt   time.Time `gorm:"not_null;autoCreateTime;autoUpdateTime"`
// }

// with conventions
type Product struct {
	ID          int
	Title       string
	Description string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// change default table name, default is "products"
// func (user *Product) TableName() string {
// 	return "products"
// }

func TestModelGorm(t *testing.T) {

}
