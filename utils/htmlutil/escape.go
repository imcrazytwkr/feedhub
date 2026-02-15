package htmlutil

import (
	"fmt"
	"io"
	"strings"
)

const escapedChars = "&'<>\"\r"

var replacements = map[byte]string{
	'&': "&amp",
	// "&#39;" is shorter than "&apos;" and apos was not in HTML until HTML5
	'\'': "&#39;",
	// mostly for the sake of consistency with "'", maybe consider "&quot;"
	'"': "&#34;",
	'<': "&lt;",
	'>': "&gt;",
	// exclusively for the sake of consistency with `html` package
	'\r': "&#13;",
}

func WriteEscaped(w io.StringWriter, str string) (err error) {
	i := strings.IndexAny(str, escapedChars)
	for i > -1 {
		_, err = w.WriteString(str[:i])
		if err != nil {
			return
		}

		replacement, exists := replacements[str[i]]
		if !exists {
			err = fmt.Errorf("unrecognized escape character %q", str[i])

			// @NOTE: this error means we have typos in code
			panic(err)
		}

		_, err = w.WriteString(replacement)
		if err != nil {
			return
		}

		str = str[i+1:]
		i = strings.IndexAny(str, escapedChars)
	}

	_, err = w.WriteString(str)
	return err
}
