package utils

import (
	"regexp"
	"strconv"
)

// ParseInt : string to int
func ParseInt(s string) int {
	r := regexp.
		MustCompile("[^0-9]").
		ReplaceAllString(s, "")

	i, err := strconv.ParseInt(r, 0, 64)
	if err != nil {
		return 0
	}
	return int(i)
}
