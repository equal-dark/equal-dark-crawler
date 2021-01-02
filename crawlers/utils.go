package crawlers

import "strings"

var symbols = []string{"£"}
var currencies = map[string]string{
	"£": "GBP",
}

// GetCurrencyFromText find text in currency
func GetCurrencyFromText(s string) (currency string) {
	for key, value := range currencies {
		if strings.Contains(s, key) {
			return value
		}
	}
	return
}
