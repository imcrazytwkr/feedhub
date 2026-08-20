package jsontree

import "strings"

func hasPrefixAt(b []byte, prefix string, position int) bool {
	end := position + len(prefix)
	return len(b) >= end && string(b[position:end]) == prefix
}

func hasPrefixFoldAt(b []byte, prefix string, position int) bool {
	end := position + len(prefix)
	return len(b) >= end && strings.EqualFold(string(b[position:end]), prefix)
}
