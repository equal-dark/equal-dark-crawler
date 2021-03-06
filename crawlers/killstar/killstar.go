package killstar

import (
	"encoding/json"
	"errors"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/thoas/go-funk"

	"equal_dark_crawler/models"
	"equal_dark_crawler/utils"
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

// IsSoldoutProduct checks soldout
func (killstar *Killstar) IsSoldoutProduct(doc *goquery.Document) bool {
	soldoutSelector := doc.Find(".mp-product-when-available > button > span")
	findNotString := regexp.MustCompile("[^a-zA-Z]")
	selectorText := findNotString.ReplaceAllString(soldoutSelector.Text(), "")
	return strings.ToLower(selectorText) == "soldout"
}

// GetProductURL returns product original url
func (killstar *Killstar) GetProductURL(doc *goquery.Document) (productURL string) {
	urlSelector := doc.Find(`link[rel="canonical"]`)
	productURL, _ = urlSelector.Attr("href")
	return
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
	currency = utils.GetCurrencyFromText(priceStr)
	return
}

// GetProductPrice returns float price
func (killstar *Killstar) GetProductPrice(doc *goquery.Document) (price float64) {
	priceSelector := doc.Find("[data-price-wrapper] .money")
	priceStr := priceSelector.First().Text()
	price = utils.GetFloatFromText(priceStr)
	return
}

// GetProductSalePrice returns float price
func (killstar *Killstar) GetProductSalePrice(doc *goquery.Document) (salePrice float64) {
	priceSelector := doc.Find("[data-price-wrapper] .money")
	priceStr := priceSelector.Last().Text()
	salePrice = utils.GetFloatFromText(priceStr)
	return
}

// GetProductImages returns product images thumbnail and big image url
func (killstar *Killstar) GetProductImages(doc *goquery.Document) (images []models.ProductImage) {
	selector := doc.Find("li[data-mp-slider-thumb] img")

	thumbnails := selector.Map(func(i int, s *goquery.Selection) string {
		src, _ := s.Attr("src")
		return "https:" + src
	})

	images = funk.Map(thumbnails, func(thumbnail string) models.ProductImage {
		src := strings.Replace(thumbnail, "_150x150_crop_center", "", -1)
		return models.ProductImage{
			Thumbnail: thumbnail,
			Src:       src,
		}
	}).([]models.ProductImage)
	return
}

// GetProductSizes returns product size name and in stock
func (killstar *Killstar) GetProductSizes(doc *goquery.Document) (sizes []models.ProductSize) {
	sizes = []models.ProductSize{}
	selector := doc.Find(`select[name="id"] option`)

	selector.Each(func(i int, s *goquery.Selection) {
		innerText := s.Text()
		rawName := strings.Split(innerText, " / ")[0]
		name := strings.Trim(
			strings.Trim(rawName, "\n"),
			" ",
		)

		_, disabled := s.Attr("disabled")
		inStock := !disabled

		sizes = append(sizes, models.ProductSize{
			Name:    name,
			InStock: inStock,
		})
	})
	return
}

func (killstar *Killstar) getProductDescriptionFromJSON(doc *goquery.Document) (description string, err error) {
	title := doc.Find(".mp-product-description strong").First().Text()

	selector := doc.Find("span[data-sheets-value]")
	descJSON, exist := selector.First().Attr("data-sheets-value")
	if !exist {
		err = errors.New("not found data-sheet-value")
		return
	}

	var dataSheet map[string]interface{}
	if err = json.Unmarshal([]byte(descJSON), &dataSheet); err != nil {
		return
	}

	desc, exist := dataSheet["2"].(string)
	if !exist {
		err = errors.New("not found data-sheet-value")
		return
	}

	description = title + "\n\n" + desc
	return
}

func (killstar *Killstar) getProductDescriptionFromSelector(doc *goquery.Document) (description string) {
	selector := doc.Find(".mp-product-description p")
	desc := selector.Map(func(i int, s *goquery.Selection) string {
		innerText := s.Contents().Map(func(i int, ss *goquery.Selection) string {
			return ss.Text()
		})
		return strings.ReplaceAll(strings.Join(innerText, "\n"), "\n\n", "\n") + "\n"
	})
	description = strings.Join(desc, "\n")
	return
}

// GetProductDescription returns description
func (killstar *Killstar) GetProductDescription(doc *goquery.Document) (description string) {
	description, err := killstar.getProductDescriptionFromJSON(doc)
	if err == nil {
		return
	}

	description = killstar.getProductDescriptionFromSelector(doc)
	return
}
