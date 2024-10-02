package main

import (
	"fmt"
	"time"

	"rust-piscine/internal/scheduler"
	"rust-piscine/internal/tests/R00"
	"rust-piscine/internal/tests/R01"
	"rust-piscine/internal/tests/R02"
	"rust-piscine/internal/tests/R03"
	"rust-piscine/internal/tests/R04"
	"rust-piscine/internal/tests/R05"
	"rust-piscine/internal/tests/R06"

	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
	Short "github.com/42-Short/shortinette/pkg/short"
	"github.com/42-Short/shortinette/pkg/short/testmodes/webhook"
)

func main() {
	modules := map[string]Module.Module{
		"00": *R00.R00(),
		"01": *R01.R01(),
		"02": *R02.R02(),
		"03": *R03.R03(),
		"04": *R04.R04(),
		"05": *R05.R05(),
		"06": *R06.R06(),
	}
	short := Short.NewShort("Rust Piscine 1.0", modules, webhook.NewWebhookTestMode(modules, "/webhook", "8080"))
	short.Start()
	if err := scheduler.Schedule(short, time.Now(), time.Second*20); err != nil {
		fmt.Println(err)
	}
}
