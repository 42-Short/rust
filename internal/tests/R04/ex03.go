package R04

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex03Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.InternalError("NOT IMPLEMENTED")
}

func ex03() Exercise.Exercise {
	return Exercise.NewExercise("03", "ex03", []string{"Cargo.toml", "src/main.rs"}, 10, ex03Test)
}
