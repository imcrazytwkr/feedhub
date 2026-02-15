package stringutil

import (
	"strings"
	"unicode"
	"unicode/utf8"
)

// @NOTE: borrowed from `strings` module since it doesn't export it
var asciiSpace = [256]uint8{'\t': 1, '\n': 1, '\v': 1, '\f': 1, '\r': 1, ' ': 1}

// This function had been extracted because we need fancy checks on `xmlquery.{TextNode,CharDataNode}`
// so as not to waste CPU and RAM on writing useless strings
func TrimTrailingSpace(str string) string {
	stop := len(str)
	for stop > 0 {
		char := str[stop-1]

		if char >= utf8.RuneSelf {
			// If we run into a non-ASCII byte, fall back to the
			// slower unicode-aware method on the remaining bytes
			return strings.TrimRightFunc(str[:stop], unicode.IsSpace)
		}

		if asciiSpace[char] == 0 {
			break
		}

		stop--
	}

	return str[:stop]
}
