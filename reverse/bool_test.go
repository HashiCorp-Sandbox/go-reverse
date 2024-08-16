// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
)

// TestInvertBoolean tests the InvertBoolean function.
func TestInvertBoolean(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    bool
		expected bool
	}{
		{true, false},
		{false, true},
	}

	for _, test := range tests {
		result := reverse.InvertBoolean(test.input)
		if result != test.expected {
			t.Errorf("InvertBoolean(%v) = %v; expected %v", test.input, result, test.expected)
		}
	}
}
