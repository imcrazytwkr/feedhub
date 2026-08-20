package mappers

import (
	"bytes"
	"encoding/json"
	"iter"

	m "github.com/imcrazytwkr/feedhub/providers/anthropic/models"
)

func publicationSectionsSeq(body []byte) iter.Seq[*m.PublicationSection] {
	return func(yield func(*m.PublicationSection) bool) {
		for line := range bytes.Lines(body) {
			if !bytes.Contains(line, []byte("publicationList")) {
				continue
			}

			s := bytes.IndexByte(line, '[')
			if s < 0 {
				continue
			}

			var row []json.RawMessage
			if json.Unmarshal(line[s:], &row) != nil {
				continue
			}

			for _, item := range row {
				// Valid preflight container is an object
				if len(item) == 0 || item[0] != '{' || item[len(item)-1] != '}' {
					continue
				}

				var container m.Container[m.PublicationSection]
				if json.Unmarshal(item, &container) != nil {
					continue
				}

				if container.Page == nil || len(container.Page.Sections) == 0 {
					continue
				}

				for _, section := range container.Page.Sections {
					if !yield(&section) {
						return
					}
				}
			}
		}
	}
}
