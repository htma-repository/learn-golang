package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

func TestLock(t *testing.T) {

	err := Db.Transaction(func(tx *gorm.DB) error {
		var user User

		err := tx.Clauses(clause.Locking{Strength: clause.LockingStrengthUpdate}).Take(&user, "id = ?", 50).Error

		if err != nil {
			return err
		}

		user.Name.FirstName = "Joko"
		user.Name.LastName = "Prasetyo"
		err = tx.Save(&user).Error
		return err
	})

	assert.NoError(t, err)
}
