package crawlers

import "github.com/PuerkitoBio/goquery"

// Killstar crawler
type Killstar struct{}

// GetProductsURL returns product url
func (killstar *Killstar) GetProductsURL(doc *goquery.Document) (productsURL []string) {
	return nil
}

// IsValidProductsPage checks page
func (killstar *Killstar) IsValidProductsPage(doc *goquery.Document) bool {
	return false
}
