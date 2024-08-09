package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex02Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("NOT IMPLEMENTED")
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "studentcode", "ex02", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::alloc::{alloc, dealloc, Layout, handle_alloc_error}", "std::ops::{Deref, DerefMut}", "std::clone::Clone", "std::marker::PhantomData", "std::ptr::{NonNull, drop_in_place}"}, map[string]int{"unsafe": 100}, 10, ex02Test)
}
