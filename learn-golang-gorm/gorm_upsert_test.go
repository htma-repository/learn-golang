package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestSaveOrUpdate(t *testing.T) {
	newUser := User{
		Name: Name{
			FirstName: "Joko",
			LastName:  "Prasetyo",
		},
		Email:    "test@example.com",
		Password: "password",
	}

	err := Db.Save(&newUser).Error //create

	assert.NoError(t, err)

	fmt.Println(newUser)

	newUser.Name.FirstName = "Jack"

	err = Db.Save(&newUser).Error // update

	assert.NoError(t, err)

	fmt.Println(newUser)
}

func TestSaveOrUpdateNonIncrement(t *testing.T) {
	newUser := User{
		Model: gorm.Model{ID: uint(5000)},
		Name: Name{
			FirstName: "Joko",
			LastName:  "Prasetyo",
		},
		Email:    "test@example.com",
		Password: "password",
	}

	err := Db.Save(&newUser).Error //create

	assert.NoError(t, err)

	fmt.Println(newUser)

	newUser.Name.FirstName = "Jack"

	err = Db.Save(&newUser).Error // update

	assert.NoError(t, err)

	fmt.Println(newUser)
}

func TestOnConflict(t *testing.T) {
	newUser := User{
		Model: gorm.Model{ID: uint(5001)},
		Name: Name{
			FirstName: "Joko",
			LastName:  "Prasetyo",
		},
		Email:    "test@example.com",
		Password: "password",
	}

	err := Db.Clauses(clause.OnConflict{UpdateAll: true}).Create(&newUser).Error

	assert.NoError(t, err)

	fmt.Println(newUser)
}
