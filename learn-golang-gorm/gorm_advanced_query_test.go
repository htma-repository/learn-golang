package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWhereOperator(t *testing.T) {
	var users = new([]User)

	err := Db.Where("first_name LIKE ?", "%john%").Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)

	var users2 = new([]User)

	err = Db.Where("first_name LIKE ?", "%john%").Where("last_name LIKE ?", "%doe%").Order("id ASC").Find(users2).Error

	assert.NoError(t, err)

	fmt.Println("users2", users2)
}

func TestOrOperator(t *testing.T) {
	var users = new([]User)

	err := Db.Where("first_name LIKE ?", "%john%").Or("password = ?", "password").Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}

func TestSelectOperator(t *testing.T) {
	var users = new([]User)

	err := Db.Select("id, first_name").Where("first_name LIKE ?", "%john%").Where("last_name LIKE ?", "%doe%").Order("id ASC").Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}

func TestNotOperator(t *testing.T) {
	var users = new([]User)
	err := Db.Not("first_name LIKE ?", "%user%").Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}

func TestStructCondition(t *testing.T) {
	userCondition := User{
		Name: Name{
			FirstName: "john",
			LastName:  "doe",
		},
	}

	var users = new([]User)

	err := Db.Where(userCondition).Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}

func TestMapCondition(t *testing.T) {
	// if want use default value like empty string,0, etc use map
	userCondition := map[string]any{
		"last_name": "",
	}

	var users = new([]User)

	err := Db.Where(userCondition).Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}

func TestOrderLimitOffset(t *testing.T) {
	var users = new([]User)

	err := Db.Select("id, first_name").Where("first_name LIKE ?", "%john%").Order("id ASC").Limit(5).Offset(2).Find(users).Error

	assert.NoError(t, err)

	fmt.Println("users", users)
}
