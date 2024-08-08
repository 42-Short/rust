package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "studentcode", "ex00", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::ptr::{write, read, add}"}, map[string]int{"unsafe": 100}, 10, ex00Test)
}
