package R02

import (
	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var cargoTestModAsString06 = ``

var clippyTomlAsString06 = ``

func ex06Test(exercise *Exercise.Exercise) Exercise.Result {
	//TODO: requires different logic with carog fuzz
    return runDefaultTest(exercise, cargoTestModAsString06, clippyTomlAsString06)
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "ex06", []string{"src/main.rs", "Cargo.toml"}, 25, ex06Test) //TODO: add actual grading points
}