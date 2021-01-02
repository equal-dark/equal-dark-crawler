package crawlers_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/PuerkitoBio/goquery"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"equal_dark_crawler/crawlers"
)

var _ = Describe("Killstar", func() {
	var killstar crawlers.Killstar

	var makeTestServer = func(statusCode int, body []byte) *httptest.Server {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.WriteHeader(statusCode)
			rw.Write(body)
		}))
		return server
	}

	Describe("IsValidProductsPage", func() {
		Context("When location is products page", func() {
			It("Should returns true", func() {
				server := makeTestServer(200, []byte(productsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When location is not products page", func() {
			It("Should returns false", func() {
				server := makeTestServer(200, []byte(mainPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductsURL", func() {

	})
})
