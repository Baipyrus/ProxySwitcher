//go:build cgo

package gui

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/util"
)

var configs map[string]*util.Config

func editConfig(cfgPath string, name string) {
	if configs == nil {
		configs = make(map[string]*util.Config)
	}

	// Skip if window is not being displayed
	if !windows[name] {
		return
	}

	// Cache configs during runtime
	config := configs[name]
	if config == nil {
		config = getConfigByName(cfgPath, name)
		if config == nil {
			config = &util.Config{Name: name}
		}

		configs[name] = config
	}

	var setViews, unsetViews []g.Widget
	for _, set := range config.Set {
		setViews = append(setViews, editVariant(set, true))
	}
	for _, unset := range config.Unset {
		unsetViews = append(unsetViews, editVariant(unset, false))
	}

	g.SingleWindow().Layout(
		append(
			append(
				append(
					append(
						[]g.Widget{
							g.Row(
								g.Label("Name:"),
								g.InputText(&config.Name),
							),
							g.Row(
								g.Label("Command:"),
								g.InputText(&config.Cmd),
							),
							g.Row(
								g.Label("Setters:"),
								g.Label(fmt.Sprintf("[%d entries]", len(config.Set))),
								g.Button("New").OnClick(func() {
									config.Set = append(config.Set, &util.Variant{})
								}),
							),
						},
						setViews...,
					),
					g.Row(
						g.Label("Unsetters:"),
						g.Label(fmt.Sprintf("[%d entries]", len(config.Unset))),
						g.Button("New").OnClick(func() {
							config.Unset = append(config.Unset, &util.Variant{})
						}),
					),
				),
				unsetViews...,
			),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("Save").OnClick(func() {
						util.SaveConfig(cfgPath, *config)
					}),
					g.Button("Close").OnClick(func() {
						windows[name] = false
						master.SetSize(440, 340)
					}),
				)),
		)...,
	)
}
