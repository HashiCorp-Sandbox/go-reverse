// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse

// ReverseSlice reverses the elements of a slice.
func ReverseSlice[T any](s []T) []T {
	result := make([]T, len(s))
	copy(result, s)
	for i, j := 0, len(result)-1; i < j; i, j = i+1, j-1 {
		result[i], result[j] = result[j], result[i]
	}
	return result
}
