package learn_golang_gorm

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func Connection() *gorm.DB {
	dialect := "hutamatr:Rahmanto123~@tcp(127.0.0.1:3306)/learn_golang_gorm?charset=utf8mb4&parseTime=True&loc=Local"
	db, err := gorm.Open(mysql.Open(dialect), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		panic(err)
	}

	return db
}

var Db = Connection()

func TestConnectionGorm(t *testing.T) {
	assert.NotNil(t, Db)
}
