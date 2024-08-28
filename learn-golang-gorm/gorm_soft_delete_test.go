package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSoftDelete(t *testing.T) {
	/// Option 1
	var user User

	err := Db.Take(&user, "id = ?", 105).Error

	assert.NoError(t, err)

	err = Db.Delete(&user).Error

	assert.NoError(t, err)

	/// Option 2

	err = Db.Delete(&User{}, "id = ?", 106).Error

	assert.NoError(t, err)

	/// Option 3

	err = Db.Where("id = ?", 107).Delete(&User{}).Error

	assert.NoError(t, err)
}

func TestQuerySoftDelete(t *testing.T) {
	var users []User

	err := Db.Find(&users, "id IN (?)", []int{105, 106, 107}).Error

	assert.NoError(t, err)
	assert.Equal(t, 0, len(users))

	fmt.Println(users)

	err = Db.Unscoped().Find(&users, "id IN (?)", []int{105, 106, 107}).Error

	assert.NoError(t, err)
	assert.Equal(t, 3, len(users))

	fmt.Println(users)

	var user User

	err = Db.Unscoped().Take(&user, "id = ?", 105).Error

	isDeleted := user.DeletedAt.Valid

	assert.NoError(t, err)
	assert.Equal(t, true, isDeleted)

	fmt.Println(isDeleted)
}
