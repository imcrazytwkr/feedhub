package portable_text

func nestLists(blocks []*node) []*node {
	tree := make([]*node, 0, len(blocks))
	var current *node

	for _, block := range blocks {
		if block == nil {
			continue
		}
		if block.kind != kindBlock || block.listItem == "" {
			tree = append(tree, block)
			current = nil
			continue
		}

		if current == nil {
			root, cur := createNestedLists(block, 0)
			current = cur
			tree = append(tree, root)
			continue
		}

		if block.ListLevel() == current.level && block.listItem == current.listItem {
			current.children = append(current.children, block)
			continue
		}

		if block.ListLevel() > current.level {
			root, cur := createNestedLists(block, current.level)
			appendNestedList(current, root)
			current = cur
			continue
		}

		if block.ListLevel() < current.level {
			if match := findListMatching(tree[len(tree)-1], block.listItem, block.ListLevel(), true); match != nil {
				current = match
				current.children = append(current.children, block)
				continue
			}
			root, cur := createNestedLists(block, 0)
			current = cur
			tree = append(tree, root)
			continue
		}

		if block.listItem != current.listItem {
			if match := findListMatching(tree[len(tree)-1], "", block.ListLevel(), false); match != nil && match.listItem == block.listItem {
				current = match
				current.children = append(current.children, block)
				continue
			}
			root, cur := createNestedLists(block, 0)
			current = cur
			tree = append(tree, root)
			continue
		}

		tree = append(tree, block)
	}

	return tree
}

func createNestedLists(block *node, startLevel int) (root, current *node) {
	level := block.ListLevel()
	first := startLevel + 1
	root = newList(block, first, listChildren(block, first, level))
	current = root
	for lv := first + 1; lv <= level; lv++ {
		list := newList(block, lv, listChildren(block, lv, level))
		appendNestedList(current, list)
		current = list
	}
	return root, current
}

func newList(block *node, level int, children []*node) *node {
	return &node{
		kind:     kindList,
		listItem: block.listItem,
		level:    level,
		children: children,
	}
}

func listChildren(block *node, listLevel, target int) []*node {
	if listLevel == target {
		return []*node{block}
	}
	item := block.Clone()
	item.children = nil
	item.level = listLevel
	return []*node{item}
}

func appendNestedList(parent, child *node) {
	if len(parent.children) == 0 {
		return
	}
	last := parent.children[len(parent.children)-1].Clone()
	last.children = append(last.children, child)
	parent.children[len(parent.children)-1] = last
}

func findListMatching(root *node, listItem string, level int, filterOnType bool) *node {
	if root == nil {
		return nil
	}

	style := "normal"
	if len(listItem) > 0 {
		style = listItem
	}

	if root.kind == kindList && root.ListLevel() == level && filterOnType && root.ListStyle() == style {
		return root
	}

	if len(root.children) == 0 {
		return nil
	}

	last := root.children[len(root.children)-1]
	if last == nil || last.kind == kindSpan {
		return nil
	}

	return findListMatching(last, listItem, level, filterOnType)
}
