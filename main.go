package main

import (
	"rust-piscine/internal/tests/R00"

	"github.com/42-Short/shortinette"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
)

func main() {
	// short := Short.NewShort("Rust Piscine 1.0", modules, webhook.NewWebhookTestMode(modules["00"]))

	shortinette.Init(map[string]Module.Module{"00": *R00.R00()})
}
