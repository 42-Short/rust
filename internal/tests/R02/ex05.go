package R02

import (
	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var cargoTestModAsString05 = ``

var clippyTomlAsString05 = ``

func ex05Test(exercise *Exercise.Exercise) Exercise.Result {
    return runDefaultTest(exercise, cargoTestModAsString05, clippyTomlAsString05)
}

func ex05() Exercise.Exercise {
	return Exercise.NewExercise("05", "ex05", []string{"src/lib.rs", "Cargo.toml"}, 25, ex05Test) //TODO: add actual grading points
}