package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

type UserResponse struct {
	ID        int
	FirstName string
	LastName  string
}

func TestQueryNonModel(t *testing.T) {
	var users []UserResponse

	err := Db.Model(&User{}).Select("id, first_name, last_name").Where("first_name LIKE ?", "%john%").Find(&users).Error

	assert.NoError(t, err)

	fmt.Println(users)
}
