package R04

import (
	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"
)

func R04() *Module.Module {
	exercises := map[string]Exercise.Exercise{
		"00": ex00(),
		"01": ex01(),
		"02": ex02(),
		"03": ex03(),
		"04": ex04(),
		"05": ex05(),
		"06": ex06(),
		"07": ex07(),
	}
	r00 := Module.NewModule("04", 50, exercises, "subject.md") // TODO: add actual subject
	return &r00
}
