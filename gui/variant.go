//go:build cgo

package gui

import (
	"strings"

	g "github.com/AllenDang/giu"
	"github.com/Baipyrus/ProxySwitcher/util"
)

func editVariant(variant *util.Variant, isSetter bool) g.Widget {
	var (
		variantArgs              = strings.Join(variant.Arguments, " ")
		variantTypes             = []string{"text", "variable"}
		discardOpts              = []string{"no", "yes"}
		height           float32 = 115
		discardProxy     g.Widget
		discardSelection int32
		typeSelection    int32
	)

	if variant.DiscardProxy {
		discardSelection = 1
	}

	if isSetter {
		discardProxy = g.Row(
			g.Label("Discard Proxy?"),
			g.Combo("", discardOpts[discardSelection], discardOpts, &discardSelection).
				OnChange(func() {
					variant.DiscardProxy = discardSelection == 1
				}).
				Size(180),
		)
		height += 25
	}

	return g.Child().Layout(
		g.Row(
			g.Label("Arguments:    "),
			g.InputText(&variantArgs).
				OnChange(func() {
					variant.Arguments = strings.Split(variantArgs, " ")
				}).
				Size(180),
		),
		g.Row(
			g.Label("Type:         "),
			g.Combo("", variantTypes[typeSelection], variantTypes, &typeSelection).
				OnChange(func() {
					variant.Type = util.VariantType(variantTypes[typeSelection])
				}).
				Size(180),
		),
		g.Row(
			g.Label("Equator:      "),
			g.InputText(&variant.Equator).Size(180),
		),
		g.Row(
			g.Label("Surround:     "),
			g.InputText(&variant.Surround).Size(180),
		),
		discardProxy,
	).Size(-1, height)
}
