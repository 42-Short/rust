package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex06Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("NOT IMPLEMENTED")
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "studentcode", "ex06", []string{"src/lib.rs", "Cargo.toml", "awesome.c", "build.rs"}, []string{"std::mem::MaybeUninit", "std::ffi::CStr", "std::ffi::{c_int, c_char}"}, map[string]int{"unsafe": 100}, 20, ex06Test)
}