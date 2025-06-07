package wordfilter

import (
	"reflect"
	"testing"
)

func TestFilterWrongPlace(t *testing.T) {
	tests := []struct {
		name              string
		input             Words
		runesWrongPlace   string
		expected          Words
	}{
		{
			name:            "Simple wrong place",
			input:           Words{"paalt", "plate", "slate", "panel"},
			runesWrongPlace: "_l___",
			expected:        Words{"paalt", "panel"}, // 'l' not in position 1, but must exist
		},
		{
			name:            "Multiple constraints",
			input:           Words{"ptanl", "lpate", "slate", "panel"},
			runesWrongPlace: "_l__t",
			expected:        Words{"ptanl", "lpate"}, // 'l' not at 1, 't' not at 4, but both must exist
		},
		{
			name:            "Letter in correct place should fail",
			input:           Words{"plant", "plate", "slate"},
			runesWrongPlace: "__a__",
			expected:        nil, // only 'slate' has 'a' NOT at position 2
		},
		{
			name:            "Single letter wrong place",
			input:           Words{"apple", "gpppp"},
			runesWrongPlace: "p____",
			expected:        Words{"apple", "gpppp"}, // 'g' is in first position in "grape", invalid
		},
		{
			name:            "All underscores - no filtering",
			input:           Words{"apple", "grape"},
			runesWrongPlace: "_____",
			expected:        Words{"apple", "grape"},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.input.FilterWrongPlace(tt.runesWrongPlace)
			if !reflect.DeepEqual(tt.input, tt.expected) {
				t.Errorf("Expected %v, got %v", tt.expected, tt.input)
			}
		})
	}
}
