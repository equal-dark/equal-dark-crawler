package crawlers

import (
	"regexp"
	"strconv"
	"strings"
)

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

// GetFloatFromText find not (number and dot) and remove, return parse float
func GetFloatFromText(s string) float64 {
	i := regexp.
		MustCompile("[^0-9|.]").
		ReplaceAllString(s, "")

	f, err := strconv.ParseFloat(i, 64)
	if err != nil {
		return 0
	}
	return f
}
