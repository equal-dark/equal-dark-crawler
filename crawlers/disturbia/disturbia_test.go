package disturbia_test

import (
	"net/http"
	"net/http/httptest"

	"github.com/PuerkitoBio/goquery"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"

	"equal_dark_crawler/crawlers/disturbia"
	"equal_dark_crawler/models"
)

var _ = Describe("Disturbia", func() {
	var disturbia disturbia.Disturbia

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

	Describe("IsValidProductPage", func() {
		Context("When it is product page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsValidProductPage(doc)

				Expect(actual).To(Equal(true))
			})
		})

		Context("When it is not product page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(disturbiaProductsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsValidProductPage(doc)

				Expect(actual).To(Equal(false))
			})
		})
	})

	Describe("IsSoldoutProduct", func() {
		Context("When it is soldout product page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(disturbiaSoldoutProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsSoldoutProduct(doc)

				Expect(actual).To(Equal(true))
			})
		})

		Context("When it is not soldout product page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.IsSoldoutProduct(doc)

				Expect(actual).To(Equal(false))
			})
		})
	})

	Describe("GetProductURL", func() {
		It("Should returns product url", func() {
			server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := disturbia.GetProductURL(doc)

			Expect(actual).To(Equal("https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest"))
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

	Describe("GetProductCurrency", func() {
		Context("When is not sale product", func() {
			It("Should returns product currency", func() {
				server := makeTestServer([]byte(disturbiaNotSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.GetProductCurrency(doc)

				Expect(actual).To(Equal("GBP"))
			})
		})

		Context("When is sale product", func() {
			It("Should returns product currency", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.GetProductCurrency(doc)

				Expect(actual).To(Equal("GBP"))
			})
		})
	})

	Describe("GetProductPrice & GetProductSalePrice", func() {
		Context("When is not sale product", func() {
			It("Should returns equal price and sale price", func() {
				server := makeTestServer([]byte(disturbiaNotSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				price := disturbia.GetProductPrice(doc)
				salePrice := disturbia.GetProductSalePrice(doc)

				Expect(price).To(Equal(float64(46)))
				Expect(salePrice).To(Equal(float64(46)))
			})
		})

		Context("When is sale product", func() {
			It("Should returns different price and sale price", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				price := disturbia.GetProductPrice(doc)
				salePrice := disturbia.GetProductSalePrice(doc)

				Expect(price).To(Equal(float64(28)))
				Expect(salePrice).To(Equal(float64(16)))
			})
		})
	})

	Describe("GetProductImages", func() {
		It("Should returns product images", func() {
			server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := disturbia.GetProductImages(doc)

			Expect(actual).To(Equal([]models.ProductImage{
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14743.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14743.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14497.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14497.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14739.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14739.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14741.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14741.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14740.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14740.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14496.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14496.jpeg",
				},
				{
					Thumbnail: "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14742.jpeg",
					Src:       "https://www.disturbia.co.uk/products/womens-all-tops/infernal-eternity-lace-up-vest/image/14742.jpeg",
				},
			}))
		})
	})

	Describe("GetProductSizes", func() {
		Context("With sizes product", func() {
			It("returns product sizes", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.GetProductSizes(doc)

				Expect(actual).To(Equal([]models.ProductSize{
					{Name: "UK 6", InStock: true},
					{Name: "UK 8", InStock: true},
					{Name: "UK 10", InStock: true},
					{Name: "UK 12", InStock: true},
					{Name: "UK 16 ", InStock: true},
				}))
			})
		})

		Context("Without sizes product", func() {
			It("returns empty product sizes", func() {
				server := makeTestServer([]byte(disturbiaNoSizeProductDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := disturbia.GetProductSizes(doc)

				Expect(actual).To(Equal([]models.ProductSize{}))
			})
		})
	})

	Describe("GetProductDescription", func() {
		It("Should returns url array", func() {
			server := makeTestServer([]byte(disturbiaSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := disturbia.GetProductDescription(doc)

			Expect(actual).To(Equal("Longline graphic vest with lace up side seam details.\n\nLarge soft-touch front screen print.&nbsp;\n\nArtwork by Mark Riddick.\n\nMetal eyelets.\n\nRegular fit.\n\n100% Cotton.&nbsp;\n"))
		})
	})
})
