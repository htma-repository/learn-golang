package learn_golang_gorm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type Order struct {
	ID        uint   `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int    `gorm:"column:user_id"`
	ProductID int    `gorm:"column:product_id"`
	CreatedAt uint64 `gorm:"column:created_at;autoCreateTime:milli"`
	UpdatedAt uint64 `gorm:"column:updated_at;autoCreateTime:milli;autoUpdateTime:milli"`
}

func TestTimestampTracking(t *testing.T) {
	newOrder := Order{
		UserID:    1,
		ProductID: 3,
	}

	err := Db.Create(&newOrder).Error

	assert.NoError(t, err)

	fmt.Println(newOrder)

	time := time.UnixMilli(int64(newOrder.CreatedAt))

	fmt.Println(time)
}

func TestUpdateTimestampTracking(t *testing.T) {
	updateOrder := Order{
		ProductID: 11,
	}

	var updatedOrder Order

	err := Db.Where("id = ?", 2).Updates(&updateOrder).Take(&updatedOrder).Error

	assert.NoError(t, err)

	assert.Equal(t, 11, updatedOrder.ProductID)

	fmt.Println(updatedOrder)

	time := time.UnixMilli(int64(updatedOrder.UpdatedAt))

	fmt.Println(time)
}
