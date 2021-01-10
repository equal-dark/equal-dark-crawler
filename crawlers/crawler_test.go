package crawlers_test

import (
	"net/http"
	"net/http/httptest"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"equal_dark_crawler/crawlers"
)

var _ = Describe("Crawler", func() {
	var makeTestServer = func(body []byte, statusCode int) *httptest.Server {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.WriteHeader(statusCode)
			rw.Write(body)
		}))
		return server
	}

	Describe("GetProductsURL", func() {
		Context("When right page", func() {
			It("Should returns urls", func() {
				server := makeTestServer([]byte(killstarProductsPageDocument), 200)
				defer server.Close()

				actual, _ := crawlers.GetProductsURL(1, server.URL)
				Expect(actual[0]).To(Equal("https://www.killstar.com/collections/womens-dresses/products/wicked-world-dress"))
				Expect(actual[1]).To(Equal("https://www.killstar.com/collections/womens-dresses/products/dark-daydreams-dress"))
			})
		})

		Context("When wrong page", func() {
			It("Should returns invalid page error", func() {
				server := makeTestServer([]byte(disturbiaProductsPageDocument), 200)
				defer server.Close()

				_, err := crawlers.GetProductsURL(1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidPage))
			})
		})

		Context("When wrong brand id", func() {
			It("Should returns invalid brand id error", func() {
				server := makeTestServer([]byte(disturbiaProductsPageDocument), 200)
				defer server.Close()

				_, err := crawlers.GetProductsURL(-1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidBrandID))
			})
		})

		Context("When status not 200", func() {
			It("Should returns invalid page error", func() {
				server := makeTestServer(nil, 400)
				defer server.Close()

				_, err := crawlers.GetProductsURL(1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidPage))
			})
		})
	})
})
