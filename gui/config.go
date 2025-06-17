//go:build cgo

package gui

import (
	"fmt"
	"strconv"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/proxy"
)

var (
	protocols         = []string{"http", "https"}
	host              string
	protocolSelection int32
	port              int32

	master  *g.MasterWindow
	windows map[string]bool
)

func configWindow(cfgPath string) {
	g.SingleWindow().Layout(
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
						g.Label("Host:"),
						g.
							InputText(&host).
							Size(310),
					),
					g.Row(
						g.Label("Port:"),
						g.
							InputInt(&port).
							Size(310),
					),
					g.Align(g.AlignCenter).To(
						g.Button("Save").
							OnClick(func() {
								p := &proxy.Proxy{Server: fmt.Sprintf(
									"%s://%s:%d",
									protocols[protocolSelection],
									host,
									port,
								)}
								proxy.SaveProxy(cfgPath, p)
							}),
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
							master.SetSize(485, 485)
						}),
				),
			),
		)...,
	)
	editConfig(cfgPath, "[NEW CONFIG]")
}

func readProxyParts(url string) {
	protocol, h, p := proxy.SplitProxyUrl(url)
	host = h

	pInt, _ := strconv.ParseInt(p, 10, 32)
	port = int32(pInt)

	if protocol == "https" {
		protocolSelection = 1
	}
}

func Config(cfgPath string) {
	p, err := proxy.ReadProxy(cfgPath)
	if err != nil || p.Server == "" {
		goto guiWindow
	}

	readProxyParts(p.Server)

guiWindow:
	windows = make(map[string]bool)
	master = g.NewMasterWindow("Proxy Switcher - Config", 440, 345, 0)
	master.Run(func() {
		configWindow(cfgPath)
	})
}
