package search

import (
	"regexp"
)

//
// https://rubular.com/r/pYxq5Bc9zVfjyW
//
func sanitizeLine(line string) string {
	regex := regexp.MustCompile(`[A-Z]*:?\s*(\w*)\s*<.*>`)
	occurrences := regex.FindStringSubmatch(line)

	if len(occurrences) == 0 {
		return ""
	}

	return occurrences[len(occurrences)-1]
}
