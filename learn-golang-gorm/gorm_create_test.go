package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestCreateDataGorm(t *testing.T) {
	user := User{
		Name: Name{
			FirstName: "John",
			LastName:  "Doe",
		},
		Email:    "john@example.com",
		Password: "password",
	}

	result := Db.Create(&user)

	assert.NoError(t, result.Error)
	assert.Equal(t, 1, int(result.RowsAffected))
}

func TestBatchCreateDataGorm(t *testing.T) {
	var users []User

	for i := 1; i <= 1000; i++ {
		users = append(users, User{
			Name: Name{
				FirstName: fmt.Sprintf("FirstName User %v", i),
				LastName:  fmt.Sprintf("LastName User %v", i),
			},
			Email:    fmt.Sprintf("user-%v@example.com", i),
			Password: "password",
		})
	}

	// result := Db.Create(&users)

	result := Db.CreateInBatches(&users, 100)

	assert.NoError(t, result.Error)
	assert.Equal(t, 1000, int(result.RowsAffected))
}
