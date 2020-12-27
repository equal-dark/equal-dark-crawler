package config_test

import (
	"equal_dark_crawler/config"
	"io/ioutil"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_Init_WithoutEnvFile(t *testing.T) {
	recovery := func() {
		r := recover()
		if r != nil {
			assert.Error(t, r.(error))
		}
	}
	defer recovery()

	config.Init()
}

func Test_Init_WithEnvFile(t *testing.T) {
	ioutil.WriteFile(".env.test", []byte(`
SCYLLA_HOST=127.0.0.1
SCYLLA_PORT=9042`), os.FileMode(0644))

	config.Init(".env.test")

	assert.Equal(t, "127.0.0.1", config.Values.Scylla.Host)
	assert.Equal(t, 9042, config.Values.Scylla.Port)

	os.Remove(".env.test")
}
