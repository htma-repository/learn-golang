package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestAutoUpsertRelation(t *testing.T) {
	user := User{
		Model: gorm.Model{
			ID: 5006,
		},
		Name: Name{
			FirstName: "Chris",
			LastName:  "Cole",
		},
		Email:    "chris@example.com",
		Password: "Password",
		Wallet: Wallet{
			UserID:  5006,
			Balance: "100000",
		},
	}

	err := Db.Create(&user).Error

	assert.NoError(t, err)
}

func TestSkipAutoUpsertRelation(t *testing.T) {
	user := User{
		Model: gorm.Model{
			ID: 5007,
		},
		Name: Name{
			FirstName: "Chris",
			LastName:  "Cole",
		},
		Email:    "chris@example.com",
		Password: "Password",
		Wallet: Wallet{
			UserID:  5007,
			Balance: "100000",
		},
	}

	err := Db.Omit(clause.Associations).Create(&user).Error

	assert.NoError(t, err)
}
