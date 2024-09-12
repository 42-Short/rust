package R02

import (
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

func R02() *Module.Module {
	exercises := map[string]Exercise.Exercise{
		"00": ex00(),
	}
	r00 := Module.NewModule("00", 70, exercises, "subject.md") // TODO: add actual subject, add actual minimum grade
	return &r00
}
