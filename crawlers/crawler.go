package crawlers

import (
	"errors"

	"github.com/PuerkitoBio/goquery"
)

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

// GetProduct is get product's info
func GetProduct(brandID int, targetURL string) (product *Product, err error) {
	brand, exist := brands[brandID]
	if !exist {
		err = ErrorInvalidBrandID
		return
	}

	doc, err := goquery.NewDocument(targetURL)
	if err != nil {
		return
	}

	if !brand.IsValidProductPage(doc) {
		err = ErrorInvalidPage
		return
	}

	product = &Product{
		URL:         targetURL,
		BrandID:     brandID,
		SoldOut:     brand.IsSoldoutProduct(doc),
		Name:        brand.GetProductName(doc),
		Currency:    brand.GetProductCurrency(doc),
		Price:       brand.GetProductPrice(doc),
		SalePrice:   brand.GetProductSalePrice(doc),
		Images:      brand.GetProductImages(doc),
		Sizes:       brand.GetProductSizes(doc),
		Description: brand.GetProductDescription(doc),
	}
	return
}
