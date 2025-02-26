// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package secret_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/secret"
)

// TestGetByID tests the GetByID function.
func TestGetByID(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    string
		expected string
	}{
		{"my/secret", "super secret"},
		{"1", "super secret"},
		{"2", "super secret"},
		{"foo", "super secret"},
		{"bar/foo", "super secret"},
	}

	for _, test := range tests {
		result := secret.GetByID(test.input)
		if result != test.expected {
			t.Errorf("GetByID(%q) = %q; expected %q", test.input, result, test.expected)
		}
	}
}
