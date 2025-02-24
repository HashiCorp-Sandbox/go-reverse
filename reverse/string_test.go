// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
)

// TestString tests the String function.
func TestString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{"hello", "olleh"},
		{"world", "dlrow"},
		{"Go", "oG"},
		{"", ""},
		{"a", "a"},
	}

	for _, test := range tests {
		result := reverse.String(test.input)
		if result != test.expected {
			t.Errorf("String(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
