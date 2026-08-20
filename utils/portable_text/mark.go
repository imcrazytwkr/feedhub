package portable_text

import (
	"iter"
	"slices"
	"sort"
	"strings"
)

type markDef struct {
	typ  string
	href string
}

var knownDecorators = []string{"strong", "em", "code", "underline", "strike-through"}

func nestMarks(block *node) []*node {
	children := block.children
	if len(children) == 0 {
		return nil
	}

	sorted := make([][]string, len(children))
	for i, child := range children {
		sorted[i] = sortMarks(child, i, children)
	}

	root := &node{kind: kindMark, typ: " "}
	stack := []*node{root}

	for i, span := range children {
		if span == nil {
			continue
		}

		needed := slices.Clone(sorted[i])
		pos := 1
		if len(stack) > 1 {
			for pos < len(stack) {
				idx := slices.Index(needed, stack[pos].markKey)
				if idx < 0 {
					break
				}

				needed = append(needed[:idx], needed[idx+1:]...)
				pos++
			}
		}

		stack = stack[:pos]
		current := stack[len(stack)-1]

		for _, markKey := range needed {
			def := block.defs[markKey]
			markType := markKey
			if def != nil {
				markType = def.typ
			}

			n := &node{kind: kindMark, typ: markType, markKey: markKey, def: def}
			current.children = append(current.children, n)
			stack = append(stack, n)
			current = n
		}

		if span.kind == kindSpan {
			for part := range detachNewLines(span.text) {
				current.children = append(current.children, &node{kind: kindText, text: part})
			}

			continue
		}

		current.children = append(current.children, span)
	}

	return root.children
}

func sortMarks(span *node, index int, siblings []*node) []string {
	if span == nil || span.kind != kindSpan || len(span.marks) == 0 {
		return nil
	}

	marks := slices.Clone(span.marks)
	occ := make(map[string]int, len(marks))
	for _, mark := range marks {
		occ[mark] = 1
		for j := index + 1; j < len(siblings); j++ {
			sib := siblings[j]
			if sib != nil && sib.kind == kindSpan && slices.Contains(sib.marks, mark) {
				occ[mark]++
				continue
			}

			break
		}
	}

	sort.SliceStable(marks, func(i, j int) bool {
		a, b := marks[i], marks[j]
		if occ[a] != occ[b] {
			return occ[a] > occ[b]
		}

		ia := slices.Index(knownDecorators, a)
		ib := slices.Index(knownDecorators, b)
		if ia != ib {
			return ia < ib
		}

		return a < b
	})

	return marks
}

/**
 * This helper detaches newline characters from their respective lines
 * so that they are correctly rendered as hard-line-breaks according
 * to [Portable Text](https://www.portabletext.org) specification
 */
func detachNewLines(s string) iter.Seq[string] {
	return func(yield func(string) bool) {
		i := strings.IndexByte(s, '\n')
		if i < 0 {
			yield(s)
			return
		}

		line := s[:i]
		if !yield(line) {
			return
		}

		s = s[i+1:]
		for len(s) > 0 {
			i = strings.IndexByte(s, '\n')
			if i < 0 {
				line, s = s, ""
			} else {
				line, s = s[:i], s[i+1:]
			}

			if !yield("\n") {
				return
			}

			if !yield(line) {
				return
			}
		}
	}
}
