package crawlers

// Product is product object
type Product struct {
	BrandID     int
	URL         string
	SoldOut     bool
	Name        string
	Currency    string
	Price       float64
	SalePrice   float64
	Images      []ProductImage
	Sizes       []ProductSize
	Description string
}

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
