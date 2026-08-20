package next_preflight

import (
	"bytes"
	"encoding/base64"
	"encoding/json"
	"strings"

	"github.com/antchfx/htmlquery"
	"golang.org/x/net/html"
)

// Collect concatenates Next.js inline Flight slices the way the browser does:
// tags 0 and 2 are skipped, tag 1 appends UTF-8, tag 3 appends decoded binary.
func Collect(root *html.Node) []byte {
	if root == nil {
		return nil
	}

	var buf bytes.Buffer
	for _, script := range htmlquery.Find(root, "//script") {
		text := scriptText(script)
		if len(text) == 0 {
			continue
		}

		data, ok := strings.CutPrefix(text, "self.__next_f.push(")
		if !ok {
			continue
		}

		data, ok = strings.CutSuffix(data, ")")
		if !ok {
			continue
		}

		var payload []json.RawMessage
		if json.Unmarshal([]byte(data), &payload) != nil {
			continue
		}

		if len(payload) == 0 {
			continue
		}

		var tag InlineFlightTag
		if json.Unmarshal(payload[0], &tag) != nil {
			continue
		}

		if len(payload) < 2 {
			continue
		}

		var chunk string
		if json.Unmarshal(payload[1], &chunk) != nil || len(chunk) == 0 {
			continue
		}

		switch tag {
		case InlineFlightTagData:
			buf.WriteString(chunk)
		case InlineFlightTagBinary:
			decoded, err := base64.StdEncoding.DecodeString(chunk)
			if err != nil {
				continue
			}

			buf.Write(decoded)
		}
	}

	if buf.Len() == 0 {
		return nil
	}

	return buf.Bytes()
}

// Valid non-empty `<script>` elements can only have a single child of type
// TextNode, everything else is invalid
func scriptText(n *html.Node) string {
	child := n.FirstChild
	if child == nil || n.LastChild != child || child.Type != html.TextNode {
		return ""
	}

	return strings.TrimSpace(child.Data)
}
