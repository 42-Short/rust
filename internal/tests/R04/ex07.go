package R04

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.InternalError("NOT IMPLEMENTED")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "ex07", []string{"Cargo.toml", "src/main.rs"}, 10, ex07Test)
}
