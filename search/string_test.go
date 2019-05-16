package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/xrash/smetrics"
)

func TestSanitizeLine(t *testing.T) {
	testCases := []struct {
		Input string
		Want  string
	}{
		{
			Input: "RESOURCE:       envFrom <[]Object>",
			Want:  "envFrom",
		},
		{
			Input: "FIELD:        prefix      <string>",
			Want:  "prefix",
		},
		{
			Input: "configMapRef <Object>",
			Want:  "configMapRef",
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Input, func(t *testing.T) {
			output := sanitizeLine(tc.Input)
			assert.Equal(t, output, tc.Want)
		})
	}
}

func TestMatchingScore(t *testing.T) {
	testCases := []struct {
		Name   string
		Given  string
		Target string
		Score  float64
	}{
		{
			Name:   "Exact match",
			Given:  "livenessProbe",
			Target: "livenessProbe",
			Score:  1.0,
		},
		{
			Name:   "Partial match",
			Given:  "liveProb",
			Target: "livenessProbe",
			Score:  smetrics.Jaro("liveProb", "livenessProbe"),
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			score := matchingScore(tc.Given, tc.Target)
			assert.Equal(t, score, tc.Score)
		})
	}
}
