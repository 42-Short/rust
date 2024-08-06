package main

import (
	"rust-piscine/internal/tests/R00"

	"github.com/42-Short/shortinette"
	"github.com/42-Short/shortinette/pkg/short/testmodes/webhooktestmode"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
	Short "github.com/42-Short/shortinette/pkg/short"
)

func main() {
	shortinette.Init()
	modules := map[string]Module.Module{
		"00": *R00.R00(),
	}
	short := Short.NewShort("Rust Piscine 1.0", modules, webhook.NewWebhookTestMode(modules))
	short.Start("00")
}
