package brands

import "github.com/PuerkitoBio/goquery"

type brand interface {
	GetProductsURL(url string) (products []string, err error)
	GetName(doc *goquery.Document) (name string, err error)
	GetPrice(doc *goquery.Document) (price float64, err error)
	GetImages(doc *goquery.Document) (images []string, err error)
	GetSizes(doc *goquery.Document) (sizes []string, err error)
	GetDescription(doc *goquery.Document) (description string, err error)
}
