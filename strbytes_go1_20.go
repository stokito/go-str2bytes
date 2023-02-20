//go:build go1.20
// +build go1.20

// SPDX-License-Identifier: MIT
package strbytes

import (
	"unsafe"
)

// BytesToStr converts byte slice to a string without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func BytesToStr(b []byte) string {
	return unsafe.String(unsafe.SliceData(b), len(b))
}

// StrToBytes converts string to a byte slice without memory allocation.
//
// Note it may break if string and/or slice header will change
// in the future go versions.
func StrToBytes(s string) (b []byte) {
	return unsafe.Slice(unsafe.StringData(s), len(s))
}
