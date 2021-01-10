package crawlers

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

// ProductImage type of product image
type ProductImage struct {
	Thumbnail string
	Src       string
}

// ProductSize type of product size name and in stock
type ProductSize struct {
	Name    string
	InStock bool
}

// Crawler interface of crawlers
type Crawler interface {
	IsValidProductsPage(doc *goquery.Document) bool
	GetProductsURL(doc *goquery.Document) []string

	IsValidProductPage(doc *goquery.Document) bool
	IsSoldoutProduct(doc *goquery.Document) bool
	GetProductName(doc *goquery.Document) string
	GetProductCurrency(doc *goquery.Document) string
	GetProductPrice(doc *goquery.Document) float64
	GetProductSalePrice(doc *goquery.Document) float64
	GetProductImages(doc *goquery.Document) []ProductImage
	GetProductSizes(doc *goquery.Document) []ProductSize
	GetProductDescription(doc *goquery.Document) string
}

var brands = map[int]Crawler{
	1: new(Killstar),
	2: new(Disturbia),
}

// ErrorInvalidBrandID is not registered brand id
var ErrorInvalidBrandID = errors.New("invalid brandID")

// ErrorInvalidPage is catched error page
var ErrorInvalidPage = errors.New("invalid products page")

// GetProductsURL get brand id and target url and returns page's products urls
func GetProductsURL(brandID int, targetURL string) (productsURL []string, err error) {
	brand, exist := brands[brandID]
	if !exist {
		err = ErrorInvalidBrandID
		return
	}

	doc, err := goquery.NewDocument(targetURL)
	if err != nil {
		return
	}

	if !brand.IsValidProductsPage(doc) {
		err = errors.New("invalid products page")
		return
	}

	productsURL = brand.GetProductsURL(doc)
	return
}
