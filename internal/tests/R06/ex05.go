package R06

import (
	"path/filepath"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var tests05 = `
#[cfg(test)]
mod shortinette_tests_rust_0605 {
    use super::*;

    #[test]
    fn test_new_tableau() {
        let tableau: Tableau<i32> = Tableau::new();
        assert_eq!(tableau.len(), 0);
        assert!(tableau.is_empty());
    }

    #[test]
    fn test_push() {
        let mut tableau = Tableau::new();
        tableau.push(10);
        assert_eq!(tableau.len(), 1);
        assert!(!tableau.is_empty());
        assert_eq!(tableau[0], 10);

        tableau.push(20);
        assert_eq!(tableau.len(), 2);
        assert_eq!(tableau[1], 20);
    }

    #[test]
    fn test_pop() {
        let mut tableau = Tableau::new();
        assert_eq!(tableau.pop(), None);

        tableau.push(5);
        tableau.push(15);
        assert_eq!(tableau.len(), 2);

        let popped = tableau.pop();
        assert_eq!(popped, Some(15));
        assert_eq!(tableau.len(), 1);

        let popped = tableau.pop();
        assert_eq!(popped, Some(5));
        assert_eq!(tableau.len(), 0);

        assert!(tableau.is_empty());
        assert_eq!(tableau.pop(), None);
    }

    #[test]
    fn test_clear() {
        let mut tableau = Tableau::new();
        tableau.push(1);
        tableau.push(2);
        tableau.push(3);
        assert_eq!(tableau.len(), 3);

        tableau.clear();
        assert_eq!(tableau.len(), 0);
        assert!(tableau.is_empty());
        assert_eq!(tableau.pop(), None);
    }

    #[test]
    fn test_deref() {
        let mut tableau = Tableau::new();
        tableau.push(100);
        tableau.push(200);
        tableau.push(300);

        let slice: &[i32] = &*tableau;
        assert_eq!(slice, &[100, 200, 300]);

        let slice_mut: &mut [i32] = &mut *tableau;
        slice_mut[1] = 250;
        assert_eq!(tableau[1], 250);
    }

    #[test]
    fn test_into_iterator() {
        let mut tableau = Tableau::new();
        tableau.push("a");
        tableau.push("b");
        tableau.push("c");

        let mut iter = tableau.into_iter();
        assert_eq!(iter.next(), Some("a"));
        assert_eq!(iter.next(), Some("b"));
        assert_eq!(iter.next(), Some("c"));
        assert_eq!(iter.next(), None);
    }

    #[test]
    fn test_iterate_over_references() {
        let mut tableau = Tableau::new();
        tableau.push(1);
        tableau.push(2);
        tableau.push(3);

        for (i, val) in tableau.iter().enumerate() {
            assert_eq!(*val, (i + 1) as i32);
        }

        for val in tableau.iter_mut() {
            *val *= 2;
        }

        let expected = [2, 4, 6];
        for (i, val) in tableau.iter().enumerate() {
            assert_eq!(*val, expected[i]);
        }
    }

    #[test]
    fn test_clone() {
        let mut tableau = Tableau::new();
        tableau.push(String::from("Hello"));
        tableau.push(String::from("World"));

        let tableau_clone = tableau.clone();
        assert_eq!(tableau.len(), tableau_clone.len());

        for (orig, clone) in tableau.iter().zip(tableau_clone.iter()) {
            assert_eq!(orig, clone);
        }

        tableau.push(String::from("!"));
        assert_eq!(tableau.len(), 3);
        assert_eq!(tableau_clone.len(), 2);
    }

    #[test]
    fn test_with_different_types() {
        let mut int_tableau = Tableau::new();
        int_tableau.push(1);
        int_tableau.push(2);
        int_tableau.push(3);
        assert_eq!(&*int_tableau, &[1, 2, 3]);

        let mut string_tableau = Tableau::new();
        string_tableau.push("foo");
        string_tableau.push("bar");
        assert_eq!(&*string_tableau, &["foo", "bar"]);

        #[derive(Debug, PartialEq)]
        struct Point {
            x: i32,
            y: i32,
        }

        let mut point_tableau = Tableau::new();
        point_tableau.push(Point { x: 1, y: 2 });
        point_tableau.push(Point { x: 3, y: 4 });

        let expected_points = [Point { x: 1, y: 2 }, Point { x: 3, y: 4 }];
        assert_eq!(&*point_tableau, &expected_points);
    }

    #[test]
    fn test_indexing() {
        let mut tableau = Tableau::new();
        tableau.push(10);
        tableau.push(20);
        tableau.push(30);

        assert_eq!(tableau[0], 10);
        assert_eq!(tableau[1], 20);
        assert_eq!(tableau[2], 30);

        tableau[1] = 25;
        assert_eq!(tableau[1], 25);
    }

    #[test]
    #[should_panic(expected = "index out of bounds")]
    fn test_index_out_of_bounds() {
        let tableau: Tableau<i32> = Tableau::new();
        let _ = tableau[0];
    }
}
`

func ex05Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.RepoDirectory, exercise.TurnInDirectory)
	if err := testutils.AppendStringToFile(tests05, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"valgrind", "test"})
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	return Exercise.Passed("OK")
}

func ex05() Exercise.Exercise {
	return Exercise.NewExercise("05", "studentcode", "ex05", []string{"src/lib.rs", "Cargo.toml"}, []string{"std::alloc::{alloc, dealloc, Layout}", "std::marker::Copy", "std::clone::Clone", "std::ops::{Deref, DerefMut}", "std::ptr::*", "std::mem::*", "cstr::cstr"}, map[string]int{"unsafe": 100}, 10, ex05Test)
}
