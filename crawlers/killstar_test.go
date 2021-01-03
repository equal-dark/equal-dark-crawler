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
				server := makeTestServer([]byte(killstarProductsPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When location is not products page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(killstarMainPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductsPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductsURL", func() {
		It("should returns url array", func() {
			server := makeTestServer([]byte(killstarProductsPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductsURL(doc)

			Expect(actual[0]).To(Equal("https://www.killstar.com/collections/womens-dresses/products/wicked-world-dress"))
		})
	})

	Describe("IsValidProductPage", func() {
		Context("When location is product page", func() {
			It("Should returns true", func() {
				server := makeTestServer([]byte(killstarSaleProductPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductPage(doc)

				Expect(actual).To(BeTrue())
			})
		})

		Context("When location is not product page", func() {
			It("Should returns false", func() {
				server := makeTestServer([]byte(killstarMainPageDocument))
				defer server.Close()

				doc, _ := goquery.NewDocument(server.URL)
				actual := killstar.IsValidProductPage(doc)

				Expect(actual).To(BeFalse())
			})
		})
	})

	Describe("GetProductName", func() {
		It("Should returns product name", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductName(doc)

			Expect(actual).To(Equal("Liliana Lace Dress"))
		})
	})

	Describe("GetProductCurrency", func() {
		It("Should returns product currency", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductCurrency(doc)

			Expect(actual).To(Equal("GBP"))
		})
	})

	Describe("GetProductPrice", func() {
		It("Should returns float price", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductPrice(doc)

			Expect(actual).To(Equal(59.99))
		})
	})

	Describe("GetProductSalePrice", func() {
		Context("With sale price", func() {
			It("Should returns float price", func() {
				server := makeTestServer([]byte(killstarSaleProductPageDocument))
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
				server := makeTestServer([]byte(killstarNotSaleProductPageDocument))
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
		It("Should returns images url array", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductImages(doc)

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

	Describe("GetProductSizes", func() {
		It("Should returns size name and in stock array", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductSizes(doc)

			Expect(actual).To(Equal([]crawlers.ProductSize{
				{Name: "XS", InStock: false},
				{Name: "S", InStock: false},
				{Name: "M", InStock: false},
				{Name: "L", InStock: true},
				{Name: "XL", InStock: false},
				{Name: "XXL", InStock: true},
			}))
		})

		It("Should returns size name and in stock array", func() {
			server := makeTestServer([]byte(killstarNotSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductSizes(doc)

			Expect(actual).To(Equal([]crawlers.ProductSize{
				{Name: "XS", InStock: true},
				{Name: "S", InStock: true},
				{Name: "M", InStock: true},
				{Name: "L", InStock: true},
				{Name: "XL", InStock: true},
				{Name: "XXL", InStock: true},
			}))
		})
	})

	Describe("GetProductDescription", func() {
		It("Should returns description", func() {
			server := makeTestServer([]byte(killstarSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductDescription(doc)

			Expect(actual).To(Equal("LILIANA.\n\nHe saw the darkness in her beauty. She saw beauty in his darkness.\n\n- Lavish Victorian Lace.\n-\u00a0High Neck & Ruffle.\n- Detachable Cross.\n- Bell Sleeves\n- Fully Lined.\n- Fitted.\n\nModern vampire vibes with the enchanted 'Liliana' maiden dress - in a luxe Victorian-style lace,\u00a0\nfully lined with a high neckline, ruffle detail, detachable hardware cross and long delicate bell-sleeves\u00a0- all finished with lace-trimmed\u00a0hems and satin bow\u00a0front.\u00a0\nTransport yourself to a dreamy candle-lit world of magic - and channel your inner goddess.Living the doll-life; match with stockings, statement jewellery and a beautiful handbag.\n\nThe model is 5’3 (160cm) and wears a size XS.\n\nWash Cold - Gentle Cycle.\n\nwith KILLSTAR Branding, 65% Nylon, 35% Cotton.\n"))
		})

		It("Should returns description", func() {
			server := makeTestServer([]byte(killstarNotSaleProductPageDocument))
			defer server.Close()

			doc, _ := goquery.NewDocument(server.URL)
			actual := killstar.GetProductDescription(doc)

			Expect(actual).To(Equal("DEITY.\u00a0\n\nShe's been through hell and came out stronger. Nothing or no-one has the power to break her magic.\n\n- Super Stretch Jersey. \n- Oversized Hood. \n- Low-cut Neckline.\n- Flared Sleeve. \n- Fitted. \n\nLive your life with unreasonable passion and see the magic that you can create. Channel yer inner goddess and look divine in the Deity Hood Dress. Conjured using a luxe soft-touch stretch jersey for a flattering and seductive fit. No detail has been spared; featuring sultry low-cut neckline, oversized hood and dramatic flare sleeve perfect for a date with the devil or night out with yer coven. Oh-so versatile -  perfect from dusk til dawn. \n\nLooks spellbinding with yer fav statement heels + a choker - mix and match as yer dark heart desires!\n\nWash Cold - Gentle Cycle. \n\nwith KILLSTAR branding, 95% Rayon, 5% Elastane. \n"))
		})
	})
})
