package R04

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex05Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.InternalError("NOT IMPLEMENTED")
}

func ex05() Exercise.Exercise {
	return Exercise.NewExercise("05", "ex05", []string{"Cargo.toml", "src/main.rs"}, 10, ex05Test)
}
