package learn_golang_viper

import (
	"fmt"
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperYAML(t *testing.T) {
	config := viper.New()
	// config.SetConfigName("config")
	// config.SetConfigType("yaml")
	config.SetConfigFile("config.yaml")
	config.AddConfigPath(".")

	if err := config.ReadInConfig(); err != nil {
		t.Error(err.Error())
	}

	assert.Equal(t, "learn-golang-viper", config.GetString("app.name"))
	assert.Equal(t, 8080, config.GetInt("app.port"))
	assert.True(t, config.GetBool("database.ssl"))
	fmt.Println(config.GetStringMap("app"))
}
