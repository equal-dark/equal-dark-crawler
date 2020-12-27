package brands

import (
	"encoding/json"
	"equal_dark_crawler/crawlers/fn"
	"equal_dark_crawler/utils"
	"errors"

	"github.com/PuerkitoBio/goquery"
)

// Killstar : brand
type Killstar struct{}

// GetProductsURL : returns products url
func (killstar *Killstar) GetProductsURL(url string) (products []string, err error) {
	protocol := "https"
	domain := "www.killstar.com"
	siteURL := fn.GetSiteURL(protocol, domain)
	doc, err := fn.GetDocumentFromURL(url)
	if err != nil {
		return
	}

	products = doc.Find("#mp-collection-grid > div").Map(func(i int, product *goquery.Selection) string {
		href, _ := product.Find("a").First().Attr("href")
		return siteURL + href
	})
	return
}

// GetName : returns product name
func (killstar *Killstar) GetName(doc *goquery.Document) (name string, err error) {
	name = doc.Find("[uk-grid] h2").First().Text()
	return
}

// GetPrice : returns price
func (killstar *Killstar) GetPrice(doc *goquery.Document) (price float64, err error) {
	selector := doc.Find("[data-price-wrapper] .money")

	if selector.Length() == 0 {
		err = errors.New("not found price")
		return
	}

	var priceStr string

	priceStr = selector.Last().Text()
	price = utils.ParseFloat(priceStr)
	return
}

// GetImages : returns image array
func (killstar *Killstar) GetImages(doc *goquery.Document) (images []string, err error) {
	selector := doc.Find("li[data-mp-slider-thumb] img")

	images = selector.Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return "https:" + src
	})
	return
}

// GetSizes : returns size array
func (killstar *Killstar) GetSizes(doc *goquery.Document) (sizes []string, err error) {
	selector := doc.Find(`button[data-mp-option="Size"]`)

	sizes = selector.Map(func(i int, s *goquery.Selection) string {
		size, _ := s.Attr("data-mp-value")
		return size
	})
	return
}

// GetDescription : returns description
func (killstar *Killstar) GetDescription(doc *goquery.Document) (description string, err error) {
	title := doc.Find(".mp-product-description strong").First().Text()

	selector := doc.Find("span[data-sheets-value]")
	desc, exist := selector.First().Attr("data-sheets-value")

	if !exist {
		err = errors.New("not found data-sheet-value")
		return
	}

	var dataSheet map[string]interface{}
	json.Unmarshal([]byte(desc), &dataSheet)

	description = title + "\n\n" + dataSheet["2"].(string)
	return
}
