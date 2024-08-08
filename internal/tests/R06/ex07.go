package R06

import Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "studentcode", "ex07", []string{"ft_putchar.rs"}, []string{"std::mem::MaybeUninit", "std::ffi::CStr", "std::ffi::{c_int, c_char}"}, map[string]int{"unsafe": 100}, 10, ex07Test)
}
