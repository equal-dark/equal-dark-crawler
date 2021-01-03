package crawlers_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestCrawlers(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Crawlers Suite")
}
