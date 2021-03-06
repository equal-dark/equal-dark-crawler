package disturbia

import (
	"equal_dark_crawler/models"
	"equal_dark_crawler/utils"
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

// IsValidProductPage checks page
func (disturbia *Disturbia) IsValidProductPage(doc *goquery.Document) bool {
	priceSelector := doc.Find(".product .detail .price")
	return priceSelector.Length() != 0
}

// IsSoldoutProduct checks soldout
func (disturbia *Disturbia) IsSoldoutProduct(doc *goquery.Document) bool {
	soldoutSelector := doc.Find(".product .detail .sold-out")
	return soldoutSelector.Length() != 0
}

// GetProductURL returns product original url
func (disturbia *Disturbia) GetProductURL(doc *goquery.Document) (productURL string) {
	urlSelector := doc.Find(`meta[property="og:url"]`)
	productURL, _ = urlSelector.Attr("content")
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
	currency = utils.GetCurrencyFromText(priceStr)
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
		index := utils.GetIntFromBool(sale)

		priceStr = strings.Split(priceSelectorText, "Now")[index]
	} else {
		priceStr = priceSelector.First().Text()
	}
	price = utils.GetFloatFromText(priceStr)
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
func (disturbia *Disturbia) GetProductImages(doc *goquery.Document) (images []models.ProductImage) {
	selector := doc.Find("ul.photos li img")

	thumbnails := selector.Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return "https://www.disturbia.co.uk" + src
	})

	images = funk.Map(thumbnails, func(src string) models.ProductImage {
		return models.ProductImage{
			Thumbnail: src,
			Src:       src,
		}
	}).([]models.ProductImage)
	return
}

// GetProductSizes returns product size name and in stock
func (disturbia *Disturbia) GetProductSizes(doc *goquery.Document) (sizes []models.ProductSize) {
	selector := doc.Find("select.stock option")

	existSizes := selector.
		FilterFunction(func(i int, s *goquery.Selection) bool {
			name := s.Text()
			return name != ""
		}).
		Map(func(i int, s *goquery.Selection) string {
			name := s.Text()
			return name
		})

	sizes = funk.Map(existSizes, func(name string) models.ProductSize {
		return models.ProductSize{
			Name:    name,
			InStock: true,
		}
	}).([]models.ProductSize)
	return
}

// GetProductDescription returns description
func (disturbia *Disturbia) GetProductDescription(doc *goquery.Document) (description string) {
	metaSelector := doc.Find(`meta[property="og:description"]`)
	description, _ = metaSelector.Attr("content")
	return
}
