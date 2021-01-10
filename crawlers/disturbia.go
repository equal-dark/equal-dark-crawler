package crawlers

import (
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/thoas/go-funk"
)

// Disturbia crawler
type Disturbia struct{}

// IsValidProductsPage checks page
func (disturbia *Disturbia) IsValidProductsPage(doc *goquery.Document) bool {
	productsClass := doc.Find("body").First().HasClass("products")
	productsLength := doc.Find("ul.products > li").Length() != 0
	categorySelector := doc.Find(".category").Length() != 0

	return productsClass && productsLength && categorySelector
}

// GetProductsURL returns product url
func (disturbia *Disturbia) GetProductsURL(doc *goquery.Document) (productsURL []string) {
	products := doc.Find("ul.products > li")
	productsURL = products.Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Find("a").First().Attr("href")
		return "https://www.disturbia.co.uk" + href
	})
	return
}

// GetProductName returns product name
func (disturbia *Disturbia) GetProductName(doc *goquery.Document) (name string) {
	nameSelector := doc.Find("h1")
	name = nameSelector.First().Text()
	return
}

// GetProductCurrency returns currency
func (disturbia *Disturbia) GetProductCurrency(doc *goquery.Document) (currency string) {
	priceSelector := doc.Find(".product .detail .price")
	priceStr := priceSelector.First().Text()
	currency = GetCurrencyFromText(priceStr)
	return
}

func (disturbia *Disturbia) isSaleProduct(priceSelector *goquery.Selection) bool {
	return priceSelector.HasClass("reduced")
}

func (disturbia *Disturbia) getProductPrice(doc *goquery.Document, sale bool) (price float64) {
	priceSelector := doc.Find(".product .detail .price")
	var priceStr string

	if disturbia.isSaleProduct(priceSelector) {
		priceSelectorText := priceSelector.First().Text()
		index := GetIntFromBool(sale)

		priceStr = strings.Split(priceSelectorText, "Now")[index]
	} else {
		priceStr = priceSelector.First().Text()
	}
	price = GetFloatFromText(priceStr)
	return
}

// GetProductPrice returns float price
func (disturbia *Disturbia) GetProductPrice(doc *goquery.Document) (price float64) {
	price = disturbia.getProductPrice(doc, false)
	return
}

// GetProductSalePrice returns float price
func (disturbia *Disturbia) GetProductSalePrice(doc *goquery.Document) (salePrice float64) {
	salePrice = disturbia.getProductPrice(doc, true)
	return
}

// GetProductImages returns product images thumbnail and big image url
func (disturbia *Disturbia) GetProductImages(doc *goquery.Document) (images []ProductImage) {
	selector := doc.Find("ul.photos li img")

	thumbnails := selector.Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return "https://www.disturbia.co.uk" + src
	})

	images = funk.Map(thumbnails, func(src string) ProductImage {
		return ProductImage{
			Thumbnail: src,
			Src:       src,
		}
	}).([]ProductImage)
	return
}

// GetProductSizes returns product size name and in stock
func (disturbia *Disturbia) GetProductSizes(doc *goquery.Document) (sizes []ProductSize) {
	selector := doc.Find("select.stock option")

	allSizes := selector.Map(func(i int, s *goquery.Selection) string {
		name := s.Text()
		return name
	})
	existSizes := funk.Filter(allSizes, func(name string) bool {
		return name != ""
	}).([]string)

	sizes = funk.Map(existSizes, func(name string) ProductSize {
		return ProductSize{
			Name:    name,
			InStock: true,
		}
	}).([]ProductSize)
	return
}
