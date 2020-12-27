package utils_test

import (
	"equal_dark_crawler/utils"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_ParseFloat_WithoutError(t *testing.T) {
	actual := utils.ParseFloat("£30.09")

	assert.Equal(t, 30.09, actual)
}

func Test_ParseFloat_WithError(t *testing.T) {
	actual := utils.ParseFloat("£")

	assert.Zero(t, actual)
}
