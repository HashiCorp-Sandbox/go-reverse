// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse_test

import (
	"testing"

	"github.com/hashicorp-sandbox/go-reverse/reverse"
)

// TestReverseSlice tests the ReverseSlice function.
func TestReverseSlice(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    []int
		expected []int
	}{
		{[]int{1, 2, 3, 4, 5}, []int{5, 4, 3, 2, 1}},
		{[]int{10, 20, 30}, []int{30, 20, 10}},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
	}

	for _, test := range tests {
		result := reverse.ReverseSlice(test.input)
		for i, v := range result {
			if v != test.expected[i] {
				t.Errorf("ReverseSlice(%v) = %v; expected %v", test.input, result, test.expected)
				break
			}
		}
	}
}
