package learn_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperENV(t *testing.T) {
	config := viper.New()
	config.SetConfigFile("development.env")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "learn-golang-viper", config.GetString("APP_NAME"))
	assert.Equal(t, 8080, config.GetInt("APP_PORT"))
	assert.True(t, config.GetBool("DB_SSL"))
}
