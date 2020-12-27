package brands_test

import (
	"equal_dark_crawler/crawlers/brands"
	"equal_dark_crawler/crawlers/fn"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func Test_GetProducts_WithHttpError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.WriteHeader(http.StatusInternalServerError)
	}))
	defer server.Close()

	killstar := new(brands.Killstar)
	_, err := killstar.GetProductsURL(server.URL)
	assert.NotNil(t, err)
}

func Test_GetProducts_WithProductsURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div id="mp-collection-grid">
				<div><a href="/1"></a></div>
				<div><a href="/2"></a></div>
			</div>
		`))
	}))
	defer server.Close()

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetProductsURL(server.URL)
	assert.Equal(
		t,
		[]string{"https://www.killstar.com/1", "https://www.killstar.com/2"},
		actual,
	)
}

func Test_GetProducts_WithoutProductsURL(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write(nil)
	}))
	defer server.Close()

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetProductsURL(server.URL)
	assert.Empty(t, actual)
}

func Test_GetPrice_WithDiscount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div data-price-wrapper="">
				<p class="uk-text-small" style="margin-bottom: 5px;"><span class="rrp-price uk-hidden">WAS: £</span>
					WAS: <span class="money">£79.99</span>
				</p>
				<div class="uk-text-center uk-text-left@m mp-text-large">                  
					<span data-product-price="" class=" mp-color-sale mp-font-weight-600 ">
						<span class="money">£55.99</span> <span class="mp-font-weight-400">(30% OFF)</span>
					</span>
				</div>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	originPrice, _ := killstar.GetPrice(doc)

	assert.Equal(t, 55.99, originPrice)
}

func Test_GetPrice_WithoutDiscount(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div data-price-wrapper="">
				<div class="uk-text-center uk-text-left@m mp-text-large">                  
					<span data-product-price="" class="">
						<span class="money">£99.99</span> 
					</span>
				</div>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	originPrice, _ := killstar.GetPrice(doc)

	assert.Equal(t, 99.99, originPrice)
}

func Test_GetPrice_WithError(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div data-price-wrapper="">
				<div class="uk-text-center uk-text-left@m mp-text-large">                  
					<span data-product-price="" class="">
					</span>
				</div>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	_, err := killstar.GetPrice(doc)

	assert.NotNil(t, err)
}

func Test_GetName(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div uk-grid>
				<div class="uk-first-column">    
					<h2 class="uk-margin-remove-bottom">Electrowerk Skirt</h2>
				</div>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetName(doc)

	assert.Equal(t, "Electrowerk Skirt", actual)
}

func Test_GetImages(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<ul id="mp-product-slider-thumbs" class="uk-thumbnav uk-thumbnav-vertical">
				<li data-mp-slider-thumb="" uk-slider-item="0" class="uk-active">
					<a href="#" class="uk-width-1-1">
						<img src="//cdn.shopify.com/s/files/1/0228/2373/products/ELECTRO-WERK-SKIRT-D_150x150_crop_center.jpg?v=1608049273" alt="" class="uk-border-rounded uk-width-1-1">
					</a>
				</li>
				<li data-mp-slider-thumb="" uk-slider-item="1">
					<a href="#" class="uk-width-1-1">
						<img src="//cdn.shopify.com/s/files/1/0228/2373/products/ELEKTRO-WERK-SKIRT-C_150x150_crop_center.jpg?v=1607514899" alt="" class="uk-border-rounded uk-width-1-1">
					</a>
				</li>
				<li data-mp-slider-thumb="" uk-slider-item="2">
					<a href="#" class="uk-width-1-1">
						<img src="//cdn.shopify.com/s/files/1/0228/2373/products/ELEKTRO-WERK-SKIRT-B_150x150_crop_center.jpg?v=1607514899" alt="" class="uk-border-rounded uk-width-1-1">
					</a>
				</li>
				<li data-mp-slider-thumb="" uk-slider-item="3">
					<a href="#" class="uk-width-1-1">
						<img src="//cdn.shopify.com/s/files/1/0228/2373/products/ELECTROWERK-SKIRT_150x150_crop_center.jpg?v=1607514899" alt="" class="uk-border-rounded uk-width-1-1">
					</a>
				</li>
			</ul>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetImages(doc)

	assert.Equal(t, []string{
		"https://cdn.shopify.com/s/files/1/0228/2373/products/ELECTRO-WERK-SKIRT-D_150x150_crop_center.jpg?v=1608049273",
		"https://cdn.shopify.com/s/files/1/0228/2373/products/ELEKTRO-WERK-SKIRT-C_150x150_crop_center.jpg?v=1607514899",
		"https://cdn.shopify.com/s/files/1/0228/2373/products/ELEKTRO-WERK-SKIRT-B_150x150_crop_center.jpg?v=1607514899",
		"https://cdn.shopify.com/s/files/1/0228/2373/products/ELECTROWERK-SKIRT_150x150_crop_center.jpg?v=1607514899",
	}, actual)
}

func Test_GetSizes(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div class="uk-width-expand uk-flex uk-flex-wrap uk-flex-between uk-flex-left@m uk-grid-margin uk-first-column" uk-margin="">
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default uk-first-column mp-variant-selected" data-mp-option-index="1" data-mp-option="Size" data-mp-value="XS" type="button">
						XS
				</button>
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default" data-mp-option-index="1" data-mp-option="Size" data-mp-value="S" type="button">
						S
				</button>
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default" data-mp-option-index="1" data-mp-option="Size" data-mp-value="M" type="button">
						M
				</button>
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default" data-mp-option-index="1" data-mp-option="Size" data-mp-value="L" type="button">
						L
				</button>
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default" data-mp-option-index="1" data-mp-option="Size" data-mp-value="XL" type="button">
						XL
				</button>
				<button class="uk-position-relative mp-variant-button uk-button uk-button-default" data-mp-option-index="1" data-mp-option="Size" data-mp-value="XXL" type="button">
						XXL
				</button>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetSizes(doc)

	assert.Equal(t, []string{"XS", "S", "M", "L", "XL", "XXL"}, actual)
}

func Test_GetDescription_WithDataSheetsValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div class="mp-product-description uk-position-relative">
				<p><strong>ELECTROWERK.</strong></p>
				<style type="text/css"><!--
				td {border: 1px solid #ccc;}br {mso-data-placement:same-cell;}
				--></style>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets-value="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">Welcome to our alternative world.</span></p>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets-value="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">- Super-soft Jersey. <br>- Wide Elastic Waist. <br>- Suspender Flares. <br>- Contrast Stitching. <br>- Fitted.<br><br>I read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.</span></p>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets-value="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">Match with crop top, vests, tunics or yer fav knits - as your heart desires!</span></p>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets-value="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">Wash Cold - Gentle Cycle. </span></p>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets-value="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">With KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. <br></span></p>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	actual, _ := killstar.GetDescription(doc)

	assert.Equal(t, `ELECTROWERK.

Welcome to our alternative world.

- Super-soft Jersey. 
- Wide Elastic Waist. 
- Suspender Flares. 
- Contrast Stitching. 
- Fitted.

I read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.

Match with crop top, vests, tunics or yer fav knits - as your heart desires!

Wash Cold - Gentle Cycle. 

With KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. 
`, actual)
}

func Test_GetDescription_WithoutDataSheetsValue(t *testing.T) {
	server := httptest.NewServer(http.HandlerFunc(func(rw http.ResponseWriter, req *http.Request) {
		rw.Write([]byte(`
			<div class="mp-product-description uk-position-relative">
				<p><strong>ELECTROWERK.</strong></p>
				<style type="text/css"><!--
				td {border: 1px solid #ccc;}br {mso-data-placement:same-cell;}
				--></style>
				<p><span data-sheets-userformat="{&quot;2&quot;:769,&quot;3&quot;:{&quot;1&quot;:0},&quot;11&quot;:4,&quot;12&quot;:0}" data-sheets="{&quot;1&quot;:2,&quot;2&quot;:&quot;Welcome to our alternative world.\n\n- Super-soft Jersey. \n- Wide Elastic Waist. \n- Suspender Flares. \n- Contrast Stitching. \n- Fitted.\n\nI read the rules before I break them; the 'Electrowerk' skirt is made of super-soft jersey - complete with a fitted waistline, statement suspender flares, and contrast stitching. A perfect piece to infuse some power and texture into your life; night or day.\n\nMatch with crop top, vests, tunics or yer fav knits - as your heart desires!\n\nWash Cold - Gentle Cycle. \n\nWith KILLSTAR Branding, 43% Cotton, 43% Viscose, 14% Elastane. \n&quot;}">Welcome to our alternative world.</span></p>
			</div>
		`))
	}))
	defer server.Close()
	doc, _ := fn.GetDocumentFromURL(server.URL)

	killstar := new(brands.Killstar)
	_, err := killstar.GetDescription(doc)

	assert.NotNil(t, err)
}
