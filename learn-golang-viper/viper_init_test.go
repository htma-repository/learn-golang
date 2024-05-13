package learn_golang_viper

import (
	"testing"

	"github.com/spf13/viper"
	"github.com/stretchr/testify/assert"
)

func TestViperInit(t *testing.T) {
	config := viper.New()

	assert.NotNil(t, config)
}
