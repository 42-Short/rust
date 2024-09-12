package R01

import (
	"rust-piscine/internal/alloweditems"
	"rust-piscine/internal/cargo"
	"time"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var cargoTest04 = `
#[cfg(test)]
mod shortinette_tests_rust_0104 {
	use super::*;

	#[test]
	fn test_empty() {
		let mut boxes = [];
		sort_boxes(&mut boxes);
		assert_eq!(boxes.len(), 0);
	}

	#[test]
	fn test_datatype() {
		let mut boxes = [[1u32, 1u32], [2u32, 2u32]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[2u32, 2u32], [1u32, 1u32]]);
	}

	#[test]
	fn test_0() {
		let mut boxes = [[2, 2], [1, 1], [3, 3]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[3, 3], [2, 2], [1, 1]]);
	}

	#[test]
	fn test_1() {
		let mut boxes = [[0, 0], [1, 1], [1, 1], [0, 0]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[1, 1], [1, 1], [0, 0], [0, 0]]);
	}

	#[test]
	fn test_2() {
		let mut boxes = [[0, 1], [1, 1]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[1, 1], [0, 1]]);
	}

	#[test]
	fn test_3() {
		let mut boxes = [[1, 0], [1, 1]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[1, 1], [1, 0]]);
	}

	#[test]
	fn test_4() {
		let mut boxes = [[1, 1], [2, 1]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[2, 1], [1, 1]]);
	}

	#[test]
	fn test_5() {
		let mut boxes = [[1, 1], [1, 2]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[1, 2], [1, 1]]);
	}

	#[test]
	fn test_6() {
		let mut boxes = [[5, 3], [5, 2], [8, 5], [2, 2], [1, 1], [2, 1]];
		sort_boxes(&mut boxes);
		assert_eq!(boxes, [[8, 5], [5, 3], [5, 2], [2, 2], [2, 1], [1, 1]]);
	}

	#[test]
	#[should_panic]
	fn test_7() {
		let mut boxes = [[5, 3], [5, 2], [8, 5], [2, 2], [1, 2], [2, 1]];
		sort_boxes(&mut boxes);
	}

	#[test]
	#[should_panic]
	fn test_8() {
		let mut boxes = [[2, 1], [5, 3], [5, 2], [1, 2], [8, 5], [2, 2]];
		sort_boxes(&mut boxes);
	}
}
`

func ex04Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, "", map[string]int{"unsafe": 0}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	if err := testutils.AppendStringToFile(cargoTest04, exercise.TurnInFiles[0]); err != nil {
		return Exercise.InternalError(err.Error())
	}
	return cargo.CargoTest(exercise, 500*time.Millisecond, []string{})
}

func ex04() Exercise.Exercise {
	return Exercise.NewExercise("04", "ex04", []string{"src/lib.rs", "Cargo.toml"}, 10, ex04Test)
}
