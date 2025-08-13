/*
Package tocebab

Returned string in cebabe-case style
*/
package tocebab

import (
	"strings"
)

func Convert(s string) string {
	result := strings.ToLower(s)
	result = strings.ReplaceAll(result, " ", "-")
	return result
}
