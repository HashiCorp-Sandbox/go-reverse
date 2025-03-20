// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package reverse

// Reverses the characters in the given string.
func String(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}
