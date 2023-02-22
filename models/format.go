package models

type Format int

const (
	FormatUndefined Format = iota
	FormatAtom
	FormatRss
)

const (
	formatAtom = "atom"
	formatRss  = "rss"
)

func (f Format) String() string {
	switch f {
	case FormatAtom:
		return formatAtom
	case FormatRss:
		return formatRss
	default:
		return formatRss
	}
}

func ParseFormat(value string) Format {
	switch value {
	case formatAtom:
		return FormatAtom
	case formatRss:
		return FormatRss
	default:
		return FormatUndefined
	}
}
