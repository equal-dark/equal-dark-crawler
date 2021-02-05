package main

import (
	"equal_dark_crawler/crawlers/crawler"
	"fmt"
)

func main() {
	productsURL, err := crawler.GetProductsURL(2, "https://www.disturbia.co.uk/products/womens-all-tops/page1")
	if err != nil {
		panic(err)
	}

	fmt.Println(productsURL)
}
