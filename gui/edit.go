//go:build cgo

package gui

import (
	"fmt"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/util"
)

func editConfig(cfgPath string, name string) {
	// Skip if window is not being displayed
	if !windows[name] {
		return
	}

	config := getConfigByName(cfgPath, name)
	if config == nil {
		config = &util.Config{Name: name}
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
							),
						},
						setViews...,
					),
					g.Row(
						g.Label("Unsetters:"),
						g.Label(fmt.Sprintf("[%d entries]", len(config.Unset))),
					),
				),
				unsetViews...,
			),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("Save").OnClick(func() {
						util.SaveConfig(cfgPath, *config)
					}),
					g.Button("Cancel").OnClick(func() {
						windows[name] = false
						master.SetSize(440, 340)
					}),
				)),
		)...,
	)
}
