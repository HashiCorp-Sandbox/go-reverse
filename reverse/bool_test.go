// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
)

// TestBoolean tests the Bool function.
func TestBoolean(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		result := reverse.Bool(test.input)
		if result != test.expected {
			t.Errorf("Boolean(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
