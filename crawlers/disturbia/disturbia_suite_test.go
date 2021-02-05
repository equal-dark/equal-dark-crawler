package disturbia_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestDisturbia(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Disturbia Suite")
}
