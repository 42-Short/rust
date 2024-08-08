package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex05Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("OK")
}

func ex05() Exercise.Exercise {
	return Exercise.NewExercise("05", "studentcode", "ex05", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::alloc::{alloc, dealloc, Layout}", "std::marker::Copy", "std::clone::Clone", "std::ops::{Deref, DerefMut}", "std::ptr::*", "std::mem::*", "cstr::cstr"}, map[string]int{"unsafe": 100}, 10, ex05Test)
}