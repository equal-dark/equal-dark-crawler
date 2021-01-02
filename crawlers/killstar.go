package crawlers

import "github.com/PuerkitoBio/goquery"

// Killstar crawler
type Killstar struct{}

// GetProductsURL returns product url
func (killstar *Killstar) GetProductsURL(doc *goquery.Document) (productsURL []string) {
	products := doc.Find("#mp-collection-grid > div")
	productsURL = products.Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Find("a").First().Attr("href")
		return "https://www.killstar.com" + href
	})
	return
}

// IsValidProductsPage checks page
func (killstar *Killstar) IsValidProductsPage(doc *goquery.Document) bool {
	products := doc.Find("#mp-collection-grid > div")
	if products.Length() == 0 {
		return false
	}
	return true
}
