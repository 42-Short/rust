package R06

import (
	"path/filepath"
	"rust-piscine/internal/alloweditems"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var tests00 = `
#[cfg(test)]
mod shortinette_tests_rust_0600 {
    use super::*;

    #[test]
    fn swap_basic_type() {
        let mut a = 69;
        let mut b = 420;
        ft_swap(&mut a, &mut b);
        assert_eq!(a, 420);
        assert_eq!(b, 69);
    }

    #[test]
    fn swap_complex_type() {
        let mut a = String::from("Hello, World!");
        let mut b = String::from("Goodbye, World!");
        ft_swap(&mut a, &mut b);
        assert_eq!(a, "Goodbye, World!");
        assert_eq!(b, "Hello, World!");
    }

    #[test]
    fn swap_zero_sized_type() {
        let mut a = ();
        let mut b = ();
        ft_swap(&mut a, &mut b);
        assert_eq!(a, ());
        assert_eq!(b, ());
    }

    #[test]
    fn strlen() {
        let s = b"Hello, World!\0";
        let len = unsafe { ft_strlen(s.as_ptr()) };
        assert_eq!(len, 13);
    }

    #[test]
    fn strcpy() {
        let s = b"Hello, World!\0";
        let mut dst = [0u8; 14];
        unsafe { ft_strcpy(dst.as_mut_ptr(), s.as_ptr()) };
        assert_eq!(&dst, s);
    }

    #[test]
    fn strcpy_insufficient_buffer() {
        let s = b"Hello, World!\0";
        let mut dst = [0u8; 5];
        unsafe { ft_strcpy(dst.as_mut_ptr(), s.as_ptr()) };
        let dst_str = std::str::from_utf8(&dst).expect("Invalid UTF-8 sequence");
        assert_eq!(dst_str, "Hello");
    }
}
`

var clippyTomlAsString00 = `
disallowed-methods = ["std::mem::replace", "std::mem::take", "std::mem::swap", "std::ptr::swap_nonoverlapping", "std::mem::size_of", "std::mem::align_of", "std::ptr::copy", "std::ptr::copy_nonoverlapping", "core::mem::replace", "core::mem::take", "core::mem::swap", "core::ptr::swap_nonoverlapping", "core::mem::size_of", "core::mem::align_of", "core::ptr::copy", "core::ptr::copy_nonoverlapping"]
`

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, clippyTomlAsString00, nil); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	if err := testutils.AppendStringToFile(tests00, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"valgrind", "test"})
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "ex00", []string{"src/lib.rs", "Cargo.toml"}, map[string]int{"unsafe": 100}, 10, ex00Test)
}
