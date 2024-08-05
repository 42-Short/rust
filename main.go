package main

import (
	"rust-piscine/internal/tests/R00"

	"github.com/42-Short/shortinette"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
	"github.com/42-Short/shortinette/pkg/short/testmodes/webhooktestmode"
	Short "github.com/42-Short/shortinette/pkg/short"
)

func main() {
	shortinette.Init()
	short := Short.NewShort("Rust Piscine 1.0", map[string]Module.Module{"00": *R00.R00()}, nil)
	short.TestMode = webhook.NewWebhookTestMode(short.Modules["00"])
	shortinette.Start(short, "00")
}
