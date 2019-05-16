package search

import (
	"regexp"

	"github.com/xrash/smetrics"
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

func matchingScore(given, target string) float64 {
	return smetrics.Jaro(given, target)
}
