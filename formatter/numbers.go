package formatter

import "regexp"

// Expression to match only numbers.
var onlyNumbersRegex = regexp.MustCompile("[^0-9]*")

//OnlyNumbers - Removes non-numeric characters, returning only numerics digits.
func OnlyNumbers(value string) string {
	return onlyNumbersRegex.ReplaceAllString(value, "")
}
