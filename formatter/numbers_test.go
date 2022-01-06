package formatter

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestOnlyNumbers(t *testing.T) {
	formatTests := []struct {
		name     string
		in       string
		expected string
	}{
		{"Document CPF formated", "702.745.280-45", "70274528045"},
		{"Document CPF formated", "Test scenario with text and digits 702.745.280-45", "70274528045"},
	}

	for _, tt := range formatTests {
		t.Run(tt.name, func(t *testing.T) {
			out := OnlyNumbers(tt.in)
			assert.Equal(t, tt.expected, out)
		})
	}
}
