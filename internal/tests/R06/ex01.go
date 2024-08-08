package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex01Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("OK")
}

func ex01() Exercise.Exercise {
	return Exercise.NewExercise("01", "studentcode", "ex01", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::mem::transmute", "std::slice::from_raw_parts"}, map[string]int{"unsafe": 101}, 10, ex01Test)
}
