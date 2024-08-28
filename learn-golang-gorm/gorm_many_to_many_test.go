package learn_golang_gorm

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManyToManyCreate(t *testing.T) {

	Db.Table("like_products").Create(map[string]interface{}{
		"user_id":    202,
		"product_id": 1,
	})

	Db.Table("like_products").Create(map[string]interface{}{
		"user_id":    203,
		"product_id": 1,
	})
}

func TestManyToManyTakeProduct(t *testing.T) {
	var product Product

	err := Db.Preload("LikedByUsers").Take(&product, "id = ?", 1).Error

	assert.NoError(t, err)
	assert.Equal(t, 5, len(product.LikedByUsers))

	fmt.Println(product.LikedByUsers)

}

func TestManyToManyTakeUser(t *testing.T) {
	var user User

	err := Db.Preload("LikeProducts").Take(&user, "id = ?", 201).Error

	assert.NoError(t, err)

	fmt.Println(user.LikeProducts)
}

func TestMany2ManyAssociationFind(t *testing.T) {
	// find like by users product
	var product Product

	err := Db.Take(&product, "id = ?", "1").Error
	assert.NoError(t, err)

	var users []User

	err = Db.Model(&product).Association("LikedByUsers").Find(&users)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(users))

	fmt.Println(users)

	err = Db.Model(&product).Where("users.first_name LIKE ?", "FirstName%").Association("LikedByUsers").Find(&users)
	assert.NoError(t, err)
	assert.Equal(t, 5, len(users))

	// find user like products
	var user User
	err = Db.Take(&user, "id = ?", 201).Error
	assert.NoError(t, err)

	var products []Product
	err = Db.Model(&user).Association("LikeProducts").Find(&products)
	assert.NoError(t, err)
	assert.Equal(t, 1, len(products))

	fmt.Println(products)
}

func TestMany2ManyAssociationAdd(t *testing.T) {
	var user User
	err := Db.Take(&user, "id = ?", 204).Error
	assert.NoError(t, err)

	var product Product
	err = Db.Take(&product, "id = ?", 2).Error
	assert.NoError(t, err)

	err = Db.Model(&product).Association("LikedByUsers").Append(&user)
	assert.NoError(t, err)
}
