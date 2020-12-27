package fn

import (
	"fmt"
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

// GetDocumentFromURL : returns document
func GetDocumentFromURL(url string) (doc *goquery.Document, err error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, fmt.Errorf("status code error: %d %s", res.StatusCode, res.Status)
	}

	doc, err = goquery.NewDocumentFromReader(res.Body)
	return doc, err
}
