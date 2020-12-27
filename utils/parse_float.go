package utils

import (
	"regexp"
	"strconv"
)

// ParseFloat : string to float
func ParseFloat(s string) float64 {
	i := regexp.
		MustCompile("[^0-9|.]").
		ReplaceAllString(s, "")

	f, err := strconv.ParseFloat(i, 64)
	if err != nil {
		return 0
	}
	return f
}
