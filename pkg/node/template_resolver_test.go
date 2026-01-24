package node

import (
	"testing"
)

func TestTemplateResolver_Resolve_WithSpaces(t *testing.T) {
	variables := map[string]interface{}{
		"foo.bar": "baz",
	}
	resolver := NewTemplateResolver(variables)

	tests := []struct {
		name     string
		input    string
		expected string
	}{
		{
			name:     "No spaces",
			input:    "Value is {{foo.bar}}",
			expected: "Value is baz",
		},
		{
			name:     "With spaces",
			input:    "Value is {{ foo.bar }}",
			expected: "Value is baz",
		},
		{
			name:     "With leading space",
			input:    "Value is {{ foo.bar}}",
			expected: "Value is baz",
		},
		{
			name:     "With trailing space",
			input:    "Value is {{foo.bar }}",
			expected: "Value is baz",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result, err := resolver.Resolve(tt.input)
			if err != nil {
				t.Fatalf("Resolve() error = %v", err)
			}
			if result != tt.expected {
				t.Errorf("Resolve() = %v, want %v", result, tt.expected)
			}
		})
	}
}
