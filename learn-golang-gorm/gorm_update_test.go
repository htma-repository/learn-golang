package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUpdateSaveGorm(t *testing.T) {
	var user User

	err := Db.Take(&user, "id = ?", 1).Error

	assert.NoError(t, err)

	user.Name.FirstName = "Hutama"
	user.Name.LastName = "Trirahmanto"

	err = Db.Save(&user).Error

	assert.NoError(t, err)
}

func TestUpdatesStructGorm(t *testing.T) {
	err := Db.Where("id = ?", 1).Updates(User{Name: Name{FirstName: "Joko", LastName: "Prasetyo"}}).Error

	assert.NoError(t, err)
}

func TestUpdatesMapGorm(t *testing.T) {
	err := Db.Model(&User{}).Where("id = ?", 1).Updates(map[string]interface{}{
		"first_name": "",
		"last_name":  "Trirahmanto",
	}).Error

	assert.NoError(t, err)
}

func TestUpdateGorm(t *testing.T) {
	err := Db.Model(&User{}).Where("id = ?", 1).Update("first_name", "Tama").Error

	assert.NoError(t, err)
}
