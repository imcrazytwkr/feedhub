package next_preflight

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

const (
	rowID = iota
	rowTag
	rowLength
	rowByNewline
	rowByLength
)

func Parse(buf []byte) (*StreamParser, error) {
	s := &StreamParser{Rows: make(map[int]*Row)}
	if len(buf) == 0 {
		return s, nil
	}

	i := 0
	state := rowID
	id := 0
	tag := byte(0)
	length := 0

	for i < len(buf) {
		switch state {
		case rowID:
			if buf[i] == ':' {
				state = rowTag
				i++
				continue
			}

			nibble, ok := hexNibble(buf[i])
			if !ok {
				id = 0
				i++
				continue
			}

			id = (id << 4) | nibble
			i++
			continue
		case rowTag:
			b := buf[i]
			if b == 'T' {
				tag = b
				state = rowLength
				length = 0
				i++
				continue
			}

			if (b >= 'A' && b <= 'Z') || b == 'r' || b == 'x' {
				tag = b
				state = rowByNewline
				i++
				continue
			}

			tag = 0
			state = rowByNewline
		case rowLength:
			if buf[i] == ',' {
				state = rowByLength
				i++
				continue
			}

			nibble, ok := hexNibble(buf[i])
			if !ok {
				state = rowByNewline
				continue
			}

			length = (length << 4) | nibble
			i++
		case rowByNewline:
			end := bytesIndexByte(buf, i, '\n')
			if end < 0 {
				s.addRow(id, tag, buf[i:])
				return s, nil
			}

			s.addRow(id, tag, buf[i:end])
			i = end + 1
			id, tag, length = 0, 0, 0
			state = rowID
		case rowByLength:
			end := min(i+length, len(buf))
			s.addRow(id, tag, buf[i:end])
			i = end
			if i < len(buf) && buf[i] == '\n' {
				i++
			}

			id, tag, length = 0, 0, 0
			state = rowID
		}
	}

	return s, nil
}

func (s *StreamParser) addRow(id int, tag byte, payload []byte) {
	row := &Row{ID: id, Tag: tag}
	switch tag {
	case 0:
		// Works because jsontree.Parse returns `nil` and not a partially
		// constructed tree as the first value when erroring
		row.Value, _ = jsontree.Parse(payload)
	case 'T':
		row.Text = string(payload)
	default:
		v, err := jsontree.Parse(payload)
		if err == nil {
			row.Value = v
		} else {
			row.Text = string(payload)
		}
	}

	s.Rows[id] = row
	s.Order = append(s.Order, id)
}

func bytesIndexByte(buf []byte, from int, c byte) int {
	for i := from; i < len(buf); i++ {
		if buf[i] == c {
			return i
		}
	}
	return -1
}
