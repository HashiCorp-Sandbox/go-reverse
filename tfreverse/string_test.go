// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfreverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/tfreverse"
	"github.com/hashicorp/terraform-plugin-framework/types"
)

// TestReverseString tests the ReverseString function.
func TestReverseString(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    types.String
		expected types.String
	}{
		{types.StringValue("hello"), types.StringValue("olleh")},
		{types.StringValue("world"), types.StringValue("dlrow")},
		{types.StringValue("Go"), types.StringValue("oG")},
		{types.StringValue(""), types.StringValue("")},
		{types.StringValue("a"), types.StringValue("a")},
	}

	for _, test := range tests {
		result := tfreverse.ReverseString(test.input)
		if result != test.expected {
			t.Errorf("ReverseString(%q) = %q; expected %q", test.input.ValueString(), result.ValueString(), test.expected.ValueString())
		}
	}
}
