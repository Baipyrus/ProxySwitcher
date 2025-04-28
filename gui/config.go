//go:build cgo

package gui

import (
	g "github.com/AllenDang/giu"
)

var (
	protocols = []string{"http", "https"}
	selection int32
	host      string
	port      int32
)

func config_window() {
	g.SingleWindow().Layout(
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
	)
}

func Config() {
	title := "Proxy Switcher - Config"
	wnd := g.NewMasterWindow(title, 400, 200, g.MasterWindowFlagsNotResizable)

	wnd.Run(config_window)
}
