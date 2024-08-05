package main

import (
	"rust-piscine/internal/tests/R00"

	"github.com/42-Short/shortinette"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
)

func main() {
	shortinette.Init(map[string]Module.Module{"00": *R00.R00()})
}
