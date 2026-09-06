// Copyright (c) 2026, Farzan Hajian
// All rights reserved.
//
// This source code is licensed under the BSD 3-Clause license found in the
// LICENSE file in the root directory of this source tree.

package domain

import "fmt"

var NotebookColors = map[string]string{
	"blue":   "#3B82F6",
	"red":    "#EF4444",
	"green":  "#22C55E",
	"yellow": "#EAB308",
	"purple": "#A855F7",
	"orange": "#F97316",
	"pink":   "#EC4899",
	"gray":   "#6B7280",
	"white":  "#FFFFFF",
	"beige":  "#F5F5F5",
}

const DefaultNotebookColor = "#3B82F6" // "blue"

func IsValidNotebookColor(hex string) bool {
	for _, v := range NotebookColors {
		if v == hex {
			return true
		}
	}
	return false
}

func ValidateNotebookColor(hex string) error {
	if !IsValidNotebookColor(hex) {
		return fmt.Errorf("invalid notebook color: %q", hex)
	}
	return nil
}
