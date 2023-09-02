package stringutils

import (
	"errors"
	"strings"
	"unicode/utf8"
)

var ErrInvalidUTF8 = errors.New("input is not a valid UTF-8")
var ErrInvalidSubstring = errors.New("base string does not contains target string")

// Reverse reverses given string!
func Reverse(s string) (string, error) {
	if !utf8.ValidString(s) {
		return s, ErrInvalidUTF8
	}
	r := []rune(s)
	lr := len(r)
	ss := make([]rune, lr)

	for i := 0; i < lr; i++ {
		ss[lr-1-i] = r[i]
	}

	return string(ss), nil
}

// Replace replaces occurrence of a substring in a given string.
func Replace(base string, target string, replacement string) (string, error) {
	if !utf8.ValidString(base) || !utf8.ValidString(target) || !utf8.ValidString(replacement) {
		return base, ErrInvalidUTF8
	}

	if !strings.Contains(base, target) {
		return base, ErrInvalidSubstring
	}

	var result string

	targetLen := len(target)
	baseLen := len(base)
	for i := 0; i < baseLen; {
		if i+targetLen <= baseLen && base[i:i+targetLen] == target {
			result += replacement
			i += targetLen
		} else if i >= baseLen {
			break
		} else {
			result += string(base[i])
			i++
		}
	}
	return result, nil
}
