package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManyToOnePreload(t *testing.T) {
	var address []Address

	err := Db.Preload("User").Find(&address).Error

	assert.NoError(t, err)

	fmt.Println(address)
}

func TestManyToOneJoins(t *testing.T) {
	var address []Address

	err := Db.Joins("User").Find(&address).Error

	assert.NoError(t, err)

	fmt.Println(address)
}
