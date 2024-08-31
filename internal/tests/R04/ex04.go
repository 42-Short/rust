package R04

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex04Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.InternalError("NOT IMPLEMENTED")
}

func ex04() Exercise.Exercise {
	return Exercise.NewExercise("04", "ex04", []string{"Cargo.toml", "src/main.rs"}, 10, ex04Test)
}
