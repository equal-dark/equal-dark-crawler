package killstar_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestKillstar(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "Killstar Suite")
}
