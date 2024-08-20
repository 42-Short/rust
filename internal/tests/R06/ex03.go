package R06

import (
	"path/filepath"
	"rust-piscine/internal/alloweditems"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var tests03 = `
#[cfg(test)]
mod shortinette_tests_rust_0603 {
    use super::*;

    #[test]
    fn test_basic() {
        let mut cell = Cellule::new(42);

        assert_eq!(cell.get(), 42);

        cell.set(100);
        assert_eq!(cell.get(), 100);

        let old_value = cell.replace(200);
        assert_eq!(old_value, 100);
        assert_eq!(cell.get(), 200);

        let value = cell.into_inner();
        assert_eq!(value, 200);
    }

    #[test]
    fn test_set() {
        let mut cell = Cellule::new(10);
        assert_eq!(cell.get(), 10);

        cell.set(20);
        assert_eq!(cell.get(), 20);
    }

    #[test]
    fn test_replace() {
        let mut cell = Cellule::new(5);
        assert_eq!(cell.get(), 5);

        let old_value = cell.replace(15);
        assert_eq!(old_value, 5);
        assert_eq!(cell.get(), 15);

        let old_value = cell.replace(30);
        assert_eq!(old_value, 15);
        assert_eq!(cell.get(), 30);
    }

    #[test]
    fn test_get_mut() {
        let mut cell = Cellule::new(7);
        assert_eq!(cell.get(), 7);

        {
            let mut_ref = cell.get_mut();
            *mut_ref = 14;
        }

        assert_eq!(cell.get(), 14);
    }

    #[test]
    fn test_into_inner() {
        let cell = Cellule::new(99);
        let value = cell.into_inner();
        assert_eq!(value, 99);
    }

    #[test]
    fn test_multiple_cells() {
        let mut cell1 = Cellule::new(1);
        let mut cell2 = Cellule::new(2);

        assert_eq!(cell1.get(), 1);
        assert_eq!(cell2.get(), 2);

        cell1.set(10);
        cell2.set(20);

        assert_eq!(cell1.get(), 10);
        assert_eq!(cell2.get(), 20);
    }

    #[test]
    fn test_copy_value() {
        let cell1 = Cellule::new(5);

        let value = cell1.get();
        let cell2 = Cellule::new(value);

        assert_eq!(cell1.get(), 5);
        assert_eq!(cell2.get(), 5);

        let value = cell2.into_inner();
        assert_eq!(value, 5);
    }

    #[test]
    fn test_copy_value_with_set() {
        let mut cell1 = Cellule::new(12);
        
        let mut value = cell1.get();
        value += 10;

        // Set the modified value back into cell1
        cell1.set(value);
        assert_eq!(cell1.get(), 22);
    }
}
`

var clippyTomlAsString03 = `
disallowed-types = ["std::cell::Cell", "std::cell::RefCell", "std::sync::Mutex", "std::mem::ManuallyDrop", "std::rc::Rc", "std::sync::Arc", "std::sync::RwLock"]
disallowed-methods = ["std::ptr::null", "std::ptr::null_mut", "std::mem::transmute"]
`

func ex03Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, clippyTomlAsString03, []string{"#![allow(clippy::doc_lazy_continuation)]", "#![allow(dead_code)]"}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.RepoDirectory, exercise.TurnInDirectory)
	if err := testutils.AppendStringToFile(tests03, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"valgrind", "test"})
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	return Exercise.Passed("OK")
}

func ex03() Exercise.Exercise {
	return Exercise.NewExercise("03", "studentcode", "ex03", []string{"src/lib.rs", "Cargo.toml"}, map[string]int{"unsafe": 100}, 10, ex03Test)
}
