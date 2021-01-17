package rogueandwolf_test

import (
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

func TestRogueAndWolf(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "RogueAndWolf Suite")
}
