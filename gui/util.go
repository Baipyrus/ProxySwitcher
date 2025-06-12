package gui

import (
	"strings"

	g "github.com/AllenDang/giu"
)

func buildNode(cfgPath string, cats map[string][]string, path string) g.Widget {
	// Use last part as node name
	parts := strings.Split(path, " - ")
	name := parts[len(parts)-1]

	// If this is a category, set frame and arrow
	if children, ok := cats[path]; ok {
		node := g.
			TreeNode(name).
			Flags(g.TreeNodeFlagsFramed)

		// Gather child nodes recursively
		var childNodes []g.Widget

		for _, child := range children {
			// Prepend path to child name for recursion
			childNodes = append(childNodes, buildNode(
				cfgPath,
				cats,
				path+" - "+child,
			))
		}

		node.Layout(childNodes...)
		return node
	}

	// If this is a leaf, create selectable element
	node := g.Row(
		g.
			Tooltip("Double click to edit config").
			To(g.
				Selectable(name).
				OnDClick(func() {
					windows[path] = true
				}),
			),
	)
	editConfig(cfgPath, path)

	return node
}

func generateTree(cfgPath string) (tree []g.Widget) {
	cats := parseCategories(cfgPath)

	// Map tree by recursively getting nodes from categories
	for _, path := range cats[""] {
		tree = append(tree, buildNode(cfgPath, cats, path))
	}

	return tree
}
