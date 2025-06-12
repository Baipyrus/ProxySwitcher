//go:build cgo

package gui

import (
	g "github.com/AllenDang/giu"
)

var (
	protocols         = []string{"http", "https"}
	host              string
	protocolSelection int32
	port              int32

	windows map[string]bool
)

func configWindow(cfgPath string) {
	g.Window("Config").Layout(
		append(
			append(
				[]g.Widget{
					g.Align(g.AlignCenter).To(
						g.Label("Proxy Server"),
					),
					g.Row(
						g.Label("Protocol:"),
						g.
							Combo("", protocols[protocolSelection], protocols, &protocolSelection).
							Size(310),
					),
					g.Row(
						g.Label("Host:    "),
						g.
							InputText(&host).
							Size(310),
					),
					g.Row(
						g.Label("Port:    "),
						g.
							InputInt(&port).
							Size(310),
					),
					g.Align(g.AlignCenter).To(
						g.Label("Configurations"),
					),
				},
				generateTree(cfgPath)...,
			),
			g.Align(g.AlignCenter).To(
				g.Row(
					g.Button("New Config").
						OnClick(func() {
							windows["[NEW CONFIG]"] = true
						}),
				),
			),
		)...,
	)
	editConfig(cfgPath, "[NEW CONFIG]")
}

func Config(cfgPath string) {
	windows = make(map[string]bool)

	wnd := g.NewMasterWindow("Proxy Switcher", 400, 300, g.MasterWindowFlagsHidden)
	wnd.Run(func() {
		configWindow(cfgPath)
	})
}
