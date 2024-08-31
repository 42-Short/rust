package R04

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex06Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.InternalError("NOT IMPLEMENTED")
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "ex06", []string{"Cargo.toml", "src/main.rs"}, 10, ex06Test)
}
