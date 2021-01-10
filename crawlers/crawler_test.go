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

	Describe("GetProduct", func() {
		Context("When right product page", func() {
			It("Should returns product", func() {
				server := makeTestServer([]byte(killstarSaleProductPageDocument), 200)
				defer server.Close()

				actual, _ := crawlers.GetProduct(1, server.URL)

				Expect(actual.SoldOut).To(BeFalse())
				Expect(actual.URL).To(Equal("https://www.killstar.com/products/liliana-lace-dress-b"))
				Expect(actual.Name).To(Equal("Liliana Lace Dress"))
				Expect(actual.Currency).To(Equal("GBP"))
				Expect(actual.Price).To(Equal(59.99))
				Expect(actual.SalePrice).To(Equal(41.99))
				Expect(actual.Images).To(Equal([]crawlers.ProductImage{
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
				Expect(actual.Sizes).To(Equal([]crawlers.ProductSize{
					{Name: "XS", InStock: false},
					{Name: "S", InStock: false},
					{Name: "M", InStock: false},
					{Name: "L", InStock: true},
					{Name: "XL", InStock: false},
					{Name: "XXL", InStock: true},
				}))
				Expect(actual.Description).To(Equal("LILIANA.\n\nHe saw the darkness in her beauty. She saw beauty in his darkness.\n\n- Lavish Victorian Lace.\n-\u00a0High Neck & Ruffle.\n- Detachable Cross.\n- Bell Sleeves\n- Fully Lined.\n- Fitted.\n\nModern vampire vibes with the enchanted 'Liliana' maiden dress - in a luxe Victorian-style lace,\u00a0\nfully lined with a high neckline, ruffle detail, detachable hardware cross and long delicate bell-sleeves\u00a0- all finished with lace-trimmed\u00a0hems and satin bow\u00a0front.\u00a0\nTransport yourself to a dreamy candle-lit world of magic - and channel your inner goddess.Living the doll-life; match with stockings, statement jewellery and a beautiful handbag.\n\nThe model is 5’3 (160cm) and wears a size XS.\n\nWash Cold - Gentle Cycle.\n\nwith KILLSTAR Branding, 65% Nylon, 35% Cotton.\n"))
			})
		})

		Context("When wrong page", func() {
			It("Should returns invalid page error", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument), 200)
				defer server.Close()

				_, err := crawlers.GetProduct(1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidPage))
			})
		})

		Context("When wrong brand id", func() {
			It("Should returns invalid brand id error", func() {
				server := makeTestServer([]byte(disturbiaSaleProductPageDocument), 200)
				defer server.Close()

				_, err := crawlers.GetProduct(-1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidBrandID))
			})
		})

		Context("When status not 200", func() {
			It("Should returns invalid page error", func() {
				server := makeTestServer([]byte("Internal server error"), 500)
				defer server.Close()

				_, err := crawlers.GetProduct(1, server.URL)
				Expect(err).To(Equal(crawlers.ErrorInvalidPage))
			})
		})
	})
})
