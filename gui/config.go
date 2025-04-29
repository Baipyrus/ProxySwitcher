//go:build cgo

package gui

import (
	"fmt"
	"slices"
	"strings"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/util"
)

func getConfigByName(cfgPath string, name string) *util.Config {
	configs, _ := util.ReadConfigs(cfgPath)

	idx := slices.IndexFunc(
		configs,
		func(other *util.Config) bool {
			ownName :=
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ToLower(name),
						" ", ""),
					"-", "_")

			otherName :=
				strings.ReplaceAll(
					strings.ReplaceAll(
						strings.ToLower(other.Name),
						" ", ""),
					"-", "_")

			return ownName == otherName
		},
	)

	return configs[idx]
}

func editConfigModal(cfgPath string, name string) g.Widget {
	config := getConfigByName(cfgPath, name)

	return g.PopupModal(fmt.Sprintf("Editing: %s", name)).Layout(
		g.Row(
			g.Label("Name:     "),
			g.InputText(&config.Name),
		),
		g.Row(
			g.Label("Command:  "),
			g.InputText(&config.Cmd),
		),
		g.Row(
			g.Label("Setters:  "),
			g.Label(fmt.Sprintf("[%d entries]", len(config.Set))),
		),
		g.Row(
			g.Label("Unsetters:"),
			g.Label(fmt.Sprintf("[%d entries]", len(config.Unset))),
		),
		g.Align(g.AlignCenter).To(
			g.Row(
				g.Button("Save").OnClick(func() {
					util.SaveConfig(cfgPath, *config)
				}),
				g.Button("Cancel").OnClick(func() { g.CloseCurrentPopup() }),
			)),
	).Flags(g.WindowFlagsAlwaysAutoResize)
}

func configWindow(cfgPath string) {
	var (
		protocols = []string{"http", "https"}
		selection int32
		host      string
		port      int32
	)

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

func buildNode(cfgPath string, cats map[string][]string, nodes map[string]g.Widget, path string) g.Widget {
	// Base case: Node already linked correctly
	// SHOULD never occur for Proxy Switcher
	if node, ok := nodes[path]; ok {
		return node
	}

	// Use last part as node name
	parts := strings.Split(path, " - ")
	name := parts[len(parts)-1]

	// If this is a category, set frame and arrow
	if children, ok := cats[path]; ok {
		node := g.
			TreeNode(name).
			Flags(g.TreeNodeFlagsFramed)
		nodes[path] = node

		// Gather child nodes recursively
		var childNodes []g.Widget

		for _, child := range children {
			// Prepend path to child name for recursion
			childNodes = append(childNodes, buildNode(
				cfgPath,
				cats,
				nodes,
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
					g.OpenPopup(fmt.Sprintf("Editing: %s", path))
				}),
			),
		editConfigModal(cfgPath, path),
	)
	nodes[path] = node

	return node
}

func generateTree(cfgPath string) (tree []g.Widget) {
	cats := parseCategories(cfgPath)

	// Map tree by recursively getting nodes from categories
	nodes := make(map[string]g.Widget)
	for _, path := range cats[""] {
		tree = append(tree, buildNode(cfgPath, cats, nodes, path))
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
