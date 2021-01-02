package crawlers

import "github.com/PuerkitoBio/goquery"

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
