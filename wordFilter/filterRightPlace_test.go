package wordfilter

import (
	"reflect"
	"testing"
)

func TestFilterRightPlace(t *testing.T) {
	tests := []struct {
		name               string
		input              Words
		runesRightInPlace  string
		expected           Words
	}{
		{
			name:              "Exact match with two words",
			input:             Words{"apple", "angle", "bmble"},
			runesRightInPlace: "a___e",
			expected:          Words{"apple", "angle"},
		},
		{
			name:              "All underscores - no filtering",
			input:             Words{"bravo", "scuba", "light"},
			runesRightInPlace: "_____",
			expected:          Words{"bravo", "scuba", "light"},
		},
		{
			name:              "Only one match",
			input:             Words{"scuba", "slate", "shame"},
			runesRightInPlace: "s____",
			expected:          Words{"scuba", "slate", "shame"},
		},
		{
			name:              "Mixed match",
			input:             Words{"scuba", "scope", "score"},
			runesRightInPlace: "s__r_",
			expected:          Words{"score"},
		},
		{
			name:              "No match",
			input:             Words{"bravo", "alpha", "gamma"},
			runesRightInPlace: "zzzzz",
			expected:          nil,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.FilterRightPlace(tt.runesRightInPlace)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, tt.input)
			}
		})
	}
}
