package main

import (
	"rust-piscine/internal/tests/R00"

	"github.com/42-Short/shortinette"
	Short "github.com/42-Short/shortinette/pkg/short"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
)

func main() {
	shortinette.Init()
	short := Short.NewShort("Rust Piscine 1.0", map[string]Module.Module{"00": *R00.R00()}, webhook.NewWebhookTestMode(modules["00"]))

}
