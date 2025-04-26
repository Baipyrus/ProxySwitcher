//go:build cgo

package gui

import (
	g "github.com/AllenDang/giu"
)

func Config() {
	title := "Proxy Switcher - Config"
	wnd := g.NewMasterWindow(title, 400, 200, g.MasterWindowFlagsNotResizable)
	wnd.Run(func() {
		g.SingleWindow().Layout(
			g.Label("TODO: Add basic proxy URL configuration"),
		)
	})
}
