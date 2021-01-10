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

	Describe("GetFloatFromText", func() {
		Context("With number in text", func() {
			It("Should returns only float", func() {
				actual := crawlers.GetFloatFromText("£59.99")

				Expect(actual).To(Equal(59.99))
			})
		})

		Context("Without number in text", func() {
			It("Should returns 0", func() {
				actual := crawlers.GetFloatFromText("£")

				Expect(actual).To(Equal(float64(0)))
			})
		})
	})

	Describe("GetIntFromBool", func() {
		Context("When true", func() {
			It("Should returns 1", func() {
				actual := crawlers.GetIntFromBool(true)

				Expect(actual).To(Equal(1))
			})
		})

		Context("When false", func() {
			It("Should returns 0", func() {
				actual := crawlers.GetIntFromBool(false)

				Expect(actual).To(Equal(0))
			})
		})
	})
})
