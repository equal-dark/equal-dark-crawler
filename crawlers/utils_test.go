package crawlers_test

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"equal_dark_crawler/crawlers"
)

var _ = Describe("Utils", func() {
	Describe("GetCurrencyFromText", func() {
		Context("With currency symbol", func() {
			It("Should returns currency", func() {
				actual := crawlers.GetCurrencyFromText("£59.99")

				Expect(actual).To(Equal("GBP"))
			})
		})

		Context("Without currency symbol", func() {
			It("Should returns empty string", func() {
				actual := crawlers.GetCurrencyFromText("59.99")

				Expect(actual).To(Equal(""))
			})
		})
	})
})
