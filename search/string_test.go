package search

import (
	"testing"

	"github.com/stretchr/testify/assert"
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
