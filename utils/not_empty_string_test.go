package utils_test

import (
	"equal_dark_crawler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_NotEmptyString_WithEmptyString(t *testing.T) {
	actual := utils.NotEmptyString("")

	assert.Equal(t, false, actual)
}

func Test_NotEmptyString_WithString(t *testing.T) {
	actual := utils.NotEmptyString("value")

	assert.Equal(t, true, actual)
}
