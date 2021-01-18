package rogueandwolf_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/PuerkitoBio/goquery"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	rogueandwolf "equal_dark_crawler/crawlers/rogue_and_wolf"
)

var _ = Describe("RogueAndWolf", func() {
	var rogueAndWolf rogueandwolf.RogueAndWolf

	var makeTestServer = func(body []byte) *httptest.Server {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.Write(body)
		}))
		return server
	}

	Describe("IsValidProductsPage", func() {
		Context("When products page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(productsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := rogueAndWolf.IsValidProductsPage(doc)

				Expect(actual).To(BeTrue())
			})
		})
	})

	Describe("GetProductsURL", func() {
		Context("When products page", func() {
			It("Should returns urls", func() {
				server := makeTestServer([]byte(productsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := rogueAndWolf.GetProductsURL(doc)

				Expect(actual[0]).To(Equal("https://rogueandwolf.com/products/affection-black-silver-midi-ring-us3-to-us5"))
			})
		})
	})

	Describe("IsValidProductPage", func() {
		Context("When products page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(productsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := rogueAndWolf.IsValidProductPage(doc)

				Expect(actual).To(BeFalse())
			})
		})

		Context("When product page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(noSaleProductDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := rogueAndWolf.IsValidProductPage(doc)

				Expect(actual).To(BeTrue())
			})
		})
	})
})
