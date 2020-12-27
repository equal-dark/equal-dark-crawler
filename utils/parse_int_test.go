package utils_test

import (
	"equal_dark_crawler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseIntWithError(t *testing.T) {
	actual := utils.ParseInt("")

	assert.Zero(t, actual)
}

func Test_ParseIntWithoutError(t *testing.T) {
	actual := utils.ParseInt("1,200")

	assert.Equal(t, 1200, actual)
}
