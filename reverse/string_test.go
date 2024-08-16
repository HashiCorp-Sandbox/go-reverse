// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
)

// TestReverseString tests the ReverseString function.
func TestReverseString(t *testing.T) {
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
		result := reverse.ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
