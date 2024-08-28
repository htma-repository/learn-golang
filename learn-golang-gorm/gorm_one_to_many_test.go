package learn_golang_gorm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

type Address struct {
	ID        uint      `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	Address   string    `gorm:"column:address"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      User      `gorm:"foreignKey:user_id;references:id"` // many to one, belongs to
}

func (a *Address) TableName() string {
	return "addresses"
}

func TestCreateOneToMany(t *testing.T) {

	user := User{
		Model: gorm.Model{
			ID: 5010,
		},
		Name: Name{
			FirstName: "Ginwoo",
			LastName:  "Onodera",
		},
		Email:    "ginwoo@example.com",
		Password: "password",
		Wallet: Wallet{
			UserID:  5010,
			Balance: "1000000",
		},
		Address: []Address{
			{
				UserID:  5010,
				Address: "Tokyo",
			},
			{
				UserID:  5010,
				Address: "Kyoto",
			},
		},
	}

	err := Db.Create(&user).Error

	assert.NoError(t, err)
}

func TestFindOneToMany(t *testing.T) {
	var users []User

	err := Db.Model(&User{}).Preload("Address").Joins("Wallet").Find(&users).Error

	assert.NoError(t, err)

	fmt.Println(users)
}

func TestTakeOneToMany(t *testing.T) {
	var user User

	err := Db.Model(&User{}).Preload("Address").Joins("Wallet").Take(&user, "users.id = ?", 5010).Error

	assert.NoError(t, err)

	assert.NotNil(t, user.ID)
	assert.NotNil(t, user.Wallet.ID)
	assert.Len(t, user.Address, 2)

	fmt.Println(user)
}
