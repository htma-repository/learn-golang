package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
)

func TestTransactionSuccess(t *testing.T) {

	err := Db.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(&User{Name: Name{FirstName: "John-1", LastName: "Doe-1"}, Email: "john@example.com", Password: "password"}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{Name: Name{FirstName: "John-2", LastName: "Doe-2"}, Email: "john@example.com", Password: "password"}).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.NoError(t, err)
}

func TestTransactionFailed(t *testing.T) {

	err := Db.Transaction(func(tx *gorm.DB) error {

		err := tx.Create(&User{Name: Name{FirstName: "John-1", LastName: "Doe-1"}, Email: "john@example.com", Password: "password"}).Error

		if err != nil {
			return err
		}

		err = tx.Create(&User{Model: gorm.Model{ID: 1}, Name: Name{FirstName: "John-2", LastName: "Doe-2"}, Email: "john@example.com", Password: "password"}).Error

		if err != nil {
			return err
		}

		return nil
	})

	assert.Error(t, err)
}

func TestManualTransactionSuccess(t *testing.T) {
	tx := Db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{Name: Name{FirstName: "John-1", LastName: "Doe-1"}, Email: "john@example.com", Password: "password"}).Error

	err = tx.Create(&User{Name: Name{FirstName: "John-2", LastName: "Doe-2"}, Email: "john@example.com", Password: "password"}).Error

	if err == nil {
		tx.Commit()
	}

	assert.NoError(t, err)
}

func TestManualTransactionFailed(t *testing.T) {
	tx := Db.Begin()
	defer tx.Rollback()

	err := tx.Create(&User{Name: Name{FirstName: "John-1", LastName: "Doe-1"}, Email: "john@example.com", Password: "password"}).Error

	err = tx.Create(&User{Model: gorm.Model{ID: 1}, Name: Name{FirstName: "John-2", LastName: "Doe-2"}, Email: "john@example.com", Password: "password"}).Error

	if err == nil {
		tx.Commit()
	}

	assert.Error(t, err)
}
