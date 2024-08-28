package learn_golang_gorm

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

type UserLog struct {
	ID        int       `gorm:"primaryKey;column:id;autoIncrement"`
	UserID    int       `gorm:"column:user_id"`
	Action    string    `gorm:"column:action"`
	CreatedAt time.Time `gorm:"column:created_at;autoCreateTime;<-:create"`
	UpdatedAt time.Time `gorm:"column:updated_at;autoCreateTime;autoUpdateTime"`
}

func TestAutoIncrement(t *testing.T) {
	userLog := UserLog{
		UserID: 1,
		Action: "Action Test",
	}

	err := Db.Create(&userLog).Error

	assert.NoError(t, err)

	fmt.Println(userLog, userLog.ID)
}
