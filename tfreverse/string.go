// Copyright (c) HashiCorp, Inc.
// SPDX-License-Identifier: MPL-2.0

package tfreverse

import "github.com/hashicorp/terraform-plugin-framework/types"

// String the individual characters in types.String.
func String(s types.String) types.String {
	runes := []rune(s.ValueString())
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return types.StringValue(string(runes))
}
