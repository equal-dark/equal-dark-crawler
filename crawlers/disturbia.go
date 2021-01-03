package crawlers

import "github.com/PuerkitoBio/goquery"

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
