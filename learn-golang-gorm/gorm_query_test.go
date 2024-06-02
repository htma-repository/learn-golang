package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestQueryFirstGorm(t *testing.T) {
	var user User

	result := Db.First(&user)

	assert.NoError(t, result.Error)

	fmt.Println(user)
}

func TestQueryLastGorm(t *testing.T) {
	var user User

	result := Db.Last(&user)

	assert.NoError(t, result.Error)

	fmt.Println(user)
}

func TestQueryTakeGorm(t *testing.T) {
	var user User

	result := Db.Take(&user, "id = ?", 2)

	assert.NoError(t, result.Error)

	fmt.Println(user)
}

func TestQueryFindGorm(t *testing.T) {
	var user []User

	result := Db.Find(&user)

	assert.NoError(t, result.Error)
	assert.Equal(t, 2004, len(user))

	fmt.Println(user)
}

func TestQueryFindInGorm(t *testing.T) {
	var user []User

	result := Db.Find(&user, "id IN (?) ORDER BY id DESC", []int{100, 200, 300, 400, 500})

	assert.NoError(t, result.Error)

	fmt.Println(user)
}
