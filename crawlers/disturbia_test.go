package crawlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/PuerkitoBio/goquery"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"equal_dark_crawler/crawlers"
)

var _ = Describe("Disturbia", func() {
	var disturbia crawlers.Disturbia

	var makeTestServer = func(body []byte) *httptest.Server {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.Write(body)
		}))
		return server
	}

	Describe("IsValidProductsPage", func() {
		Context("When is products page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(disturbiaProductsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsValidProductsPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When is not products page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsValidProductsPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductsURL", func() {
		It("Should returns url array", func() {
			server := makeTestServer([]byte(disturbiaProductsPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := disturbia.GetProductsURL(doc)

			Expect(actual[0]).To(Equal("https://www.disturbia.co.uk/products/womens-all-tops/Blaze-Jumper"))
		})
	})

	Describe("GetProductName", func() {
		It("Should returns product name", func() {
			server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := disturbia.GetProductName(doc)

			Expect(actual).To(Equal("Infernal Eternity Lace Up Vest"))
		})
	})
})
