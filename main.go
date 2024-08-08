package main

import (
	"rust-piscine/internal/tests/R00"
	"rust-piscine/internal/tests/R06"

	"github.com/42-Short/shortinette"
	"github.com/42-Short/shortinette/pkg/short/testmodes/webhooktestmode"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
	Short "github.com/42-Short/shortinette/pkg/short"
)

func main() {
	shortinette.Init()
	modules := map[string]Module.Module{
		"00": *R00.R00(),
		// TODO: "01": *R01.R01(),
		// TODO: "02": *R02.R02(),
		// TODO: "03": *R03.R03(),
		// TODO: "04": *R04.R04(),
		// TODO: "05": *R05.R05(),
		"06": *R06.R06(),
	}
	short := Short.NewShort("Rust Piscine 1.0", modules, webhook.NewWebhookTestMode(modules))
	short.Start("00")
}
