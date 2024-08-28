package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestDelete(t *testing.T) {
	/// Option 1
	var user User

	err := Db.Take(&user, "id = ?", 102).Error

	assert.NoError(t, err)

	err = Db.Delete(&user).Error

	assert.NoError(t, err)

	/// Option 2

	err = Db.Delete(&User{}, "id = ?", 103).Error

	assert.NoError(t, err)

	/// Option 3

	err = Db.Where("id = ?", 104).Delete(&User{}).Error

	assert.NoError(t, err)
}
