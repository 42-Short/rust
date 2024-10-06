package R02

import (
	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var cargoTestModAsString02=`

#[cfg(test)]
mod shortinette_tests_rust_0202 {
    use super::*;

    fn test_from_delivery_time(start: u32, end: u32, expected_status: PizzaStatus) {
        for day in start..end {
            assert_eq!(PizzaStatus::from_delivery_time(day), expected_status, "Day: {}", day);
        }
    }
    
    #[test]
    fn test_from_delivery_time_ranges() {
        test_from_delivery_time(0, 2, PizzaStatus::Ordered);
        test_from_delivery_time(2, 7, PizzaStatus::Cooking);
        test_from_delivery_time(7, 10, PizzaStatus::Cooked);
        test_from_delivery_time(10, 17, PizzaStatus::Delivering);
        test_from_delivery_time(17, 31, PizzaStatus::Delivered);
    }

    #[test]
    fn test_get_delivery_time_in_days() {
        let test_cases = [
            (PizzaStatus::Ordered, 17),
            (PizzaStatus::Cooking, 15),
            (PizzaStatus::Cooked, 10),
            (PizzaStatus::Delivering, 7),
            (PizzaStatus::Delivered, 0),
        ];
        for (status, expected_days) in test_cases {
            assert_eq!(status.get_delivery_time_in_days(), expected_days, "Status: {:?}", status);
        }
    }
}

`

var clippyTomlAsString02 = ``

func ex02Test(exercise *Exercise.Exercise) Exercise.Result {
    return runDefaultTest(exercise, cargoTestModAsString02, clippyTomlAsString02, map[string]int{"unsafe": 0})
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "ex02", []string{"src/lib.rs", "Cargo.toml"}, 10, ex02Test)
}