package portable_text

type kind int

const (
	kindUnknown kind = iota
	kindBlock
	kindList
	kindSpan
	kindText
	kindMark
)

type node struct {
	kind     kind
	typ      string
	style    string
	listItem string
	level    int
	children []*node
	marks    []string
	text     string
	defs     map[string]*markDef
	def      *markDef
	markKey  string
}

func (n *node) ListLevel() int {
	if n == nil {
		return 1
	}

	return max(n.level, 1)
}

func (n *node) ListStyle() string {
	if n == nil || len(n.listItem) == 0 {
		return "normal"
	}

	return n.listItem
}

func (n *node) Clone() *node {
	if n == nil {
		return nil
	}

	c := *n
	if n.children != nil {
		c.children = append([]*node(nil), n.children...)
	}

	if n.marks != nil {
		c.marks = append([]string(nil), n.marks...)
	}

	return &c
}
