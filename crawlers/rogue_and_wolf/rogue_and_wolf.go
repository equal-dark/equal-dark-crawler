package rogueandwolf

import (
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
