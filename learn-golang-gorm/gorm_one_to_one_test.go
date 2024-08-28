package learn_golang_gorm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Wallet struct {
	ID        uint      `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	Balance   string    `gorm:"column:balance"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
	User      *User     `gorm:"foreignKey:user_id;references:id"` // one to one, belongs to
}

func (w *Wallet) TableName() string {
	return "wallets"
}

func TestOneToOneHasOne(t *testing.T) {
	wallet := Wallet{
		UserID:  59,
		Balance: "200000",
	}

	err := Db.Create(&wallet).Error

	assert.NoError(t, err)

	fmt.Println(wallet)

	var user User

	err = Db.Model(&User{}).Preload("Wallet").Take(&user, "id = ?", 59).Error // if use preload, sql query 2 times, query user & query wallet

	assert.NoError(t, err)

	fmt.Println(user)
}

func TestOneToOneHasOneJoins(t *testing.T) {
	wallet := Wallet{
		UserID:  154,
		Balance: "200000",
	}

	err := Db.Create(&wallet).Error

	assert.NoError(t, err)

	fmt.Println(wallet)

	var user User

	err = Db.Model(&User{}).Joins("Wallet").Take(&user, "users.id = ?", 154).Error // if use join, sql only query 1 time

	assert.NoError(t, err)

	fmt.Println(user.Wallet.Balance)
}

func TestOneToOneBelongsToPreload(t *testing.T) {
	var wallets []Wallet

	err := Db.Preload("User").Find(&wallets).Error

	assert.NoError(t, err)

	fmt.Println(wallets)
}

func TestOneToOneBelongsToJoins(t *testing.T) {
	var wallets []Wallet

	err := Db.Joins("User").Find(&wallets).Error

	assert.NoError(t, err)

	fmt.Println(wallets)

	fmt.Println(wallets[0].User.Name.FirstName)
}
