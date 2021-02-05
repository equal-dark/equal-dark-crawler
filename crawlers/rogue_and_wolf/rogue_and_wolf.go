package rogueandwolf

import (
	"equal_dark_crawler/utils"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

// RogueAndWolf : rogue + wolf crawler
type RogueAndWolf struct{}

// IsValidProductsPage checks page
func (rogueAndWolf *RogueAndWolf) IsValidProductsPage(doc *goquery.Document) bool {
	selector := doc.Find("#CollectionAjaxContent")
	return selector.Length() != 0
}

// GetProductsURL returns product url
func (rogueAndWolf *RogueAndWolf) GetProductsURL(doc *goquery.Document) (productsURL []string) {
	products := doc.Find("a.grid-product__link")
	productsURL = products.Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Attr("href")
		return "https://rogueandwolf.com" + strings.ReplaceAll(href, "collections/new/", "")
	})
	return
}

// IsValidProductPage checks page
func (rogueAndWolf *RogueAndWolf) IsValidProductPage(doc *goquery.Document) bool {
	priceSelector := doc.Find(".product__price")
	return priceSelector.Length() != 0
}

// IsSoldoutProduct checks soldout
func (rogueAndWolf *RogueAndWolf) IsSoldoutProduct(doc *goquery.Document) bool {
	soldOutText := doc.Find(".add-to-cart").First().Text()
	return strings.Contains(soldOutText, "Sold Out")
}

// GetProductName returns product name
func (rogueAndWolf *RogueAndWolf) GetProductName(doc *goquery.Document) (name string) {
	nameSelector := doc.Find("meta[property=\"og:title\"]")
	name, _ = nameSelector.First().Attr("content")
	return
}

// GetProductPrice returns float price
func (rogueAndWolf *RogueAndWolf) GetProductPrice(doc *goquery.Document) (price float64) {
	priceStr := doc.Find(".product-single__meta div.product__price .money").First().Text()
	price = utils.GetFloatFromText(priceStr)
	return
}
