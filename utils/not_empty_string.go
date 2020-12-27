package utils

import "github.com/thoas/go-funk"

// NotEmptyString : returns input value is not empty
func NotEmptyString(s string) bool {
	return funk.NotEmpty(s)
}
