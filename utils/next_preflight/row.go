package next_preflight

import "github.com/imcrazytwkr/feedhub/utils/jsontree"

type Row struct {
	ID    int
	Tag   byte // 0 means JSON model
	Value *jsontree.Value
	Text  string
}
