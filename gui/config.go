//go:build cgo

package gui

import (
	"slices"
	"strings"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/util"
)

var (
	protocols = []string{"http", "https"}
	selection int32
	host      string
	port      int32
)

func configWindow(cfgPath string) {
	g.SingleWindow().Layout(
		append(
			[]g.Widget{
				g.Align(g.AlignCenter).To(
					g.Label("Proxy Server"),
				),
				g.Row(
					g.Label("Protocol:"),
					g.Combo("", protocols[selection], protocols, &selection),
				),
				g.Row(
					g.Label("Host:    "),
					g.InputText(&host),
				),
				g.Row(
					g.Label("Port:    "),
					g.InputInt(&port),
				),
				g.Align(g.AlignCenter).To(
					g.Label("Configurations"),
				),
			},
			generateTree(cfgPath)...,
		)...,
	)
}

func splitPath(path string) [][]string {
	parts := strings.Split(path, "-")
	for i := range parts {
		parts[i] = strings.TrimSpace(parts[i])
	}

	var result [][]string
	for i := range parts {
		// Concatenate all paths up to current
		result = append(result, parts[:i+1])
	}

	return result
}

func parseCategories(cfgPath string) map[string][]string {
	// Read configs, split into categories based on name
	cats := make(map[string][]string)
	configs, _ := util.ReadConfigs(cfgPath)

	for _, cfg := range configs {
		for _, levels := range splitPath(cfg.Name) {
			// Extract hierarchy and item name
			parent := strings.Join(levels[:len(levels)-1], " - ")
			item := levels[len(levels)-1]

			// Initialize category slice if empty
			if _, ok := cats[parent]; !ok {
				cats[parent] = []string{item}
				continue
			}

			// Ignore duplicate entries
			if slices.Contains(cats[parent], item) {
				continue
			}

			cats[parent] = append(cats[parent], item)
		}
	}

	return cats
}

func buildNode(cats map[string][]string, nodes map[string]*g.TreeNodeWidget, path string) *g.TreeNodeWidget {
	// Base case: Node already linked correctly
	// SHOULD never occur for Proxy Switcher
	if node, ok := nodes[path]; ok {
		return node
	}

	// Use last part as node name
	parts := strings.Split(path, " - ")
	name := parts[len(parts)-1]
	node := g.TreeNode(name)
	nodes[path] = node

	if children, ok := cats[path]; ok {
		// Gather child nodes recursively
		var childNodes []g.Widget

		for _, child := range children {
			// Prepend path to child name for recursion
			childNodes = append(childNodes, buildNode(
				cats,
				nodes,
				path+" - "+child,
			))
		}

		node.Layout(childNodes...)
	}

	return node
}

func generateTree(cfgPath string) (tree []g.Widget) {
	cats := parseCategories(cfgPath)

	// Map tree by recursively getting nodes from categories
	nodes := make(map[string]*g.TreeNodeWidget)
	for _, path := range cats[""] {
		tree = append(tree, buildNode(cats, nodes, path))
	}

	return tree
}

func Config(cfgPath string) {
	title := "Proxy Switcher - Config"
	wnd := g.NewMasterWindow(title, 600, 400, 0)
	wnd.Run(func() {
		configWindow(cfgPath)
	})
}
