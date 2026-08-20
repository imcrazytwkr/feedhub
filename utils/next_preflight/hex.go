package next_preflight

func parseHexPrefix(s string) (int, int) {
	n := 0
	id := 0
	for n < len(s) {
		nibble, ok := hexNibble(s[n])
		if !ok {
			break
		}
		id = (id << 4) | nibble
		n++
	}
	return id, n
}

func hexNibble(b byte) (int, bool) {
	switch {
	case b >= '0' && b <= '9':
		return int(b - '0'), true
	case b >= 'a' && b <= 'f':
		return int(b - 'a' + 10), true
	case b >= 'A' && b <= 'F':
		return int(b - 'A' + 10), true
	default:
		return 0, false
	}
}
