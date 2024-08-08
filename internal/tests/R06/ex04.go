package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex04Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("OK")
}

func ex04() Exercise.Exercise {
	return Exercise.NewExercise("04", "studentcode", "ex04", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::copy::Copy", "std::clone::Clone", "std::str::from_utf8_unchecked", "libc::__errno_location", "libc::strerror", "libc::{write, read, open, close}", "cstr::cstr", "std::cmp::{PartialEq, Eq, PartialOrd, Ord}", "std::fmt::{Debug, Display}"}, map[string]int{"unsafe": 100}, 10, ex04Test)
}