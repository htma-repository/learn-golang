package learn_golang_viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperJSON(t *testing.T) {
	config := viper.New()
	config.SetConfigName("config")
	config.SetConfigType("json")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "learn-golang-viper", config.GetString("app.name"))
	assert.Equal(t, 3000, config.GetInt("app.port"))
	assert.True(t, config.GetBool("database.ssl"))
	fmt.Println(config.GetStringMap("app"))
}
