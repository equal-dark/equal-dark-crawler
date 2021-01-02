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

	var makeTestServer = func(body []byte) *httptest.Server {
		server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
			rw.Header().Set("Content-Type", "text/html")
			rw.Write(body)
		}))
		return server
	}

	Describe("IsValidProductsPage", func() {
		Context("When location is products page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(productsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When location is not products page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(mainPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductsURL", func() {
		It("should returns url array", func() {
			server := makeTestServer([]byte(productsPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductsURL(doc)

			Expect(actual[0]).To(Equal("https://www.killstar.com/collections/womens-dresses/products/wicked-world-dress"))
		})
	})

	Describe("IsValidProductPage", func() {
		Context("When location is product page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(saleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When location is not product page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(mainPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductName", func() {
		It("Should returns product name", func() {
			server := makeTestServer([]byte(saleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductName(doc)

			Expect(actual).To(Equal("Liliana Lace Dress"))
		})
	})

	Describe("GetProductCurrency", func() {
		It("Should returns product currency", func() {
			server := makeTestServer([]byte(saleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductCurrency(doc)

			Expect(actual).To(Equal("GBP"))
		})
	})

	Describe("GetProductPrice", func() {
		It("Should returns float price", func() {
			server := makeTestServer([]byte(saleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductPrice(doc)

			Expect(actual).To(Equal(59.99))
		})
	})

	Describe("GetProductSalePrice", func() {
		Context("With sale price", func() {
			It("Should returns float price", func() {
				server := makeTestServer([]byte(saleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				price := killstar.GetProductPrice(doc)
				salePrice := killstar.GetProductSalePrice(doc)

				Expect(price).NotTo(Equal(salePrice))
				Expect(salePrice).To(Equal(41.99))
			})
		})

		Context("Without sale price", func() {
			It("Should returns float price", func() {
				server := makeTestServer([]byte(notSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				price := killstar.GetProductPrice(doc)
				salePrice := killstar.GetProductSalePrice(doc)

				Expect(price).To(Equal(salePrice))
				Expect(salePrice).To(Equal(34.99))
			})
		})
	})

	Describe("GetProductImagesURL", func() {
		It("Should returns images url string array", func() {
			server := makeTestServer([]byte(saleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductImagesURL(doc)

			Expect(actual).To(Equal([]crawlers.ProductImage{
				{
					Thumbnail: "https://cdn.shopify.com/s/files/1/0228/2373/products/liliana-lace-dress_150x150_crop_center.jpg?v=1593094735",
					Src:       "https://cdn.shopify.com/s/files/1/0228/2373/products/liliana-lace-dress.jpg?v=1593094735",
				},
				{
					Thumbnail: "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-B_150x150_crop_center.jpg?v=1589376638",
					Src:       "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-B.jpg?v=1589376638",
				},
				{
					Thumbnail: "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-C_150x150_crop_center.jpg?v=1589376638",
					Src:       "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-C.jpg?v=1589376638",
				},
				{
					Thumbnail: "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS_150x150_crop_center.jpg?v=1589376638",
					Src:       "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS.jpg?v=1589376638",
				},
				{
					Thumbnail: "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-D_150x150_crop_center.jpg?v=1589376638",
					Src:       "https://cdn.shopify.com/s/files/1/0228/2373/products/LILIANA-DRESS-D.jpg?v=1589376638",
				},
			}))
		})
	})
})
