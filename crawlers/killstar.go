package crawlers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/thoas/go-funk"
)

// Killstar crawler
type Killstar struct{}

// IsValidProductsPage checks page
func (killstar *Killstar) IsValidProductsPage(doc *goquery.Document) bool {
	products := doc.Find("#mp-collection-grid > div")
	return products.Length() != 0
}

// GetProductsURL returns product url
func (killstar *Killstar) GetProductsURL(doc *goquery.Document) (productsURL []string) {
	products := doc.Find("#mp-collection-grid > div")
	productsURL = products.Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Find("a").First().Attr("href")
		return "https://www.killstar.com" + href
	})
	return
}

// IsValidProductPage checks page
func (killstar *Killstar) IsValidProductPage(doc *goquery.Document) bool {
	nameSelector := doc.Find("[uk-grid] h2")
	return nameSelector.Length() != 0
}

// GetProductName returns product name
func (killstar *Killstar) GetProductName(doc *goquery.Document) (name string) {
	nameSelector := doc.Find("[uk-grid] h2")
	name = nameSelector.First().Text()
	return
}

// GetProductCurrency returns currency
func (killstar *Killstar) GetProductCurrency(doc *goquery.Document) (currency string) {
	priceSelector := doc.Find("[data-price-wrapper] .money")
	priceStr := priceSelector.First().Text()
	currency = GetCurrencyFromText(priceStr)
	return
}

// GetProductPrice returns float price
func (killstar *Killstar) GetProductPrice(doc *goquery.Document) (price float64) {
	priceSelector := doc.Find("[data-price-wrapper] .money")
	priceStr := priceSelector.First().Text()
	price = GetFloatFromText(priceStr)
	return
}

// GetProductSalePrice returns float price
func (killstar *Killstar) GetProductSalePrice(doc *goquery.Document) (salePrice float64) {
	priceSelector := doc.Find("[data-price-wrapper] .money")
	priceStr := priceSelector.Last().Text()
	salePrice = GetFloatFromText(priceStr)
	return
}

// GetProductImagesURL returns product images url
func (killstar *Killstar) GetProductImagesURL(doc *goquery.Document) (images []ProductImage) {
	selector := doc.Find("li[data-mp-slider-thumb] img")

	thumbnails := selector.Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return "https:" + src
	})

	images = funk.Map(thumbnails, func(thumbnail string) ProductImage {
		src := strings.Replace(thumbnail, "_150x150_crop_center", "", -1)
		return ProductImage{
			Thumbnail: thumbnail,
			Src:       src,
		}
	}).([]ProductImage)
	return
}
