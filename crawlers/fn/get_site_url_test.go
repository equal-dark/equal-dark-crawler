package fn_test

import (
	"equal_dark_crawler/crawlers/fn"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetSiteURL(t *testing.T) {
	protocol := "https"
	domain := "www.killstar.com"
	actual := fn.GetSiteURL(protocol, domain)

	assert.Equal(t, protocol+"://"+domain, actual)
}
