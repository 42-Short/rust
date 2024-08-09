package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex03Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("NOT IMPLEMENTED")
}

func ex03() Exercise.Exercise {
	return Exercise.NewExercise("03", "studentcode", "ex03", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::clone::Clone  std::marker::Copy", "std::cell::UnsafeCell", "std::ptr::*", "std::mem::*"}, map[string]int{"unsafe": 100}, 10, ex03Test)
}