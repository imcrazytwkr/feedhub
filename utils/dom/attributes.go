package dom

import (
	"strings"

	"golang.org/x/net/html"
)

func ParseAttributes(attrs []html.Attribute) map[string]string {
	if len(attrs) == 0 {
		return nil
	}

	result := make(map[string]string, len(attrs))
	for _, attr := range attrs {
		value := strings.TrimSpace(attr.Val)
		if len(value) > 0 {
			result[attr.Key] = value
		}
	}

	return result
}

func SerializeAttributes(attrs map[string]string) []html.Attribute {
	if len(attrs) == 0 {
		return nil
	}

	result := make([]html.Attribute, len(attrs))
	i := 0

	for key, val := range attrs {
		result[i].Key = key
		result[i].Val = val
		i++
	}

	if i < len(attrs) {
		result = result[:i]
	}

	return result
}
