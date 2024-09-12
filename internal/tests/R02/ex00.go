package R02

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var CargoTestModAsString00 = `

#[cfg(test)]
mod shortinette_tests_rust_0200 {
    use super::*;

    #[test]
    fn test_seconds_to_minutes_types() {
        let s: Seconds = 120.0;
        let m: Minutes = seconds_to_minutes(s);
        assert_eq!(m, 2.0);
    }

    #[test]
    fn test_seconds_to_minutes_zero() {
        let seconds: Seconds = 0.0;
        let minutes = seconds_to_minutes(seconds);
        assert_eq!(minutes, 0.0);
    }

    #[test]
    fn test_seconds_to_minutes_negative() {
        let seconds: Seconds = -120.0;
        let minutes = seconds_to_minutes(seconds);
        assert_eq!(minutes, -2.0);
    }

    #[test]
    fn test_seconds_to_minutes_large_input() {
        let seconds: Seconds = 3600.0;
        let minutes = seconds_to_minutes(seconds);
        assert_eq!(minutes, 60.0);
    }

    #[test]
    fn test_seconds_to_minutes_fractional() {
        let seconds: Seconds = 30.5;
        let minutes = seconds_to_minutes(seconds);
        assert_eq!(minutes, 0.5083333333333333);
    }

}`

var clippyTomlAsString00 = ``

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(CargoTestModAsString00, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"valgrind test"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "ex00", []string{"src/main.rs", "src/lib.rs", "Cargo.toml"}, 25, ex00Test) //TODO: add actual grading points
}