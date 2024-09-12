package R02

import (
	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var cargoTestModAsString07 = ``

var clippyTomlAsString07 = ``

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	//TODO: add logic for asserting 
    return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "ex07", []string{"src/main.rs", "Cargo.toml"}, 25, ex07Test) //TODO: add actual grading points
}