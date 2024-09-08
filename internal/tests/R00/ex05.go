package R00

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"rust-piscine/internal/alloweditems"
	"strconv"
	"strings"
	"time"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

const CargoTestEx05 = `
#[cfg(test)]
mod shortinette_tests_rust_0005 {
    use super::*;

    #[test]
    fn test_is_leap_year() {
        assert!(is_leap_year(1600));
        assert!(!is_leap_year(1500));
        assert!(is_leap_year(2004));
        assert!(!is_leap_year(2003));
    }

    #[test]
    fn test_num_days_in_month() {
		assert_eq!(num_days_in_month(2001, 1), 31);
		assert_eq!(num_days_in_month(2000, 1), 31);
        assert_eq!(num_days_in_month(2000, 2), 29);
        assert_eq!(num_days_in_month(2001, 2), 28);
        assert_eq!(num_days_in_month(2004, 2), 29);
		assert_eq!(num_days_in_month(2100, 2), 28);
		assert_eq!(num_days_in_month(2001, 3), 31);
		assert_eq!(num_days_in_month(2000, 3), 31);
		assert_eq!(num_days_in_month(2001, 4), 30);
		assert_eq!(num_days_in_month(2000, 4), 30);
		assert_eq!(num_days_in_month(2001, 5), 31);
		assert_eq!(num_days_in_month(2000, 5), 31);
		assert_eq!(num_days_in_month(2001, 6), 30);
		assert_eq!(num_days_in_month(2000, 6), 30);
		assert_eq!(num_days_in_month(2001, 7), 31);
		assert_eq!(num_days_in_month(2000, 7), 31);
		assert_eq!(num_days_in_month(2001, 8), 31);
		assert_eq!(num_days_in_month(2000, 8), 31);
		assert_eq!(num_days_in_month(2001, 9), 30);
		assert_eq!(num_days_in_month(2000, 9), 30);
		assert_eq!(num_days_in_month(2001, 10), 31);
		assert_eq!(num_days_in_month(2000, 10), 31);
		assert_eq!(num_days_in_month(2001, 11), 30);
		assert_eq!(num_days_in_month(2000, 11), 30);
		assert_eq!(num_days_in_month(2001, 12), 31);
		assert_eq!(num_days_in_month(2000, 12), 31);
    }

    #[test]
    #[should_panic]
    fn test_invalid_year() {
        is_leap_year(0);
    }

    #[test]
    #[should_panic]
    fn test_invalid_month() {
        num_days_in_month(2000, 13);
    }

	#[test]
    #[should_panic]
    fn test_invalid_month_2() {
        num_days_in_month(0, 2);
    }
}
`

func missingTests(message string) Exercise.Result {
	return Exercise.Result{
		Passed: false,
		Output: fmt.Sprintf("missing test: %s", message),
	}
}

func monthTestExistenceTest(content string) Exercise.Result {
	leapYearsFound := make(map[uint64]bool)
	commonYearsFound := make(map[uint64]bool)
	invalidFound := false

	pattern := `num_days_in_month\((\d+),\s*(\d+)\)`
	regex := regexp.MustCompile(pattern)
	matches := regex.FindAllStringSubmatch(content, -1)
	for _, match := range matches {
		if len(match) == 3 {
			year, err := strconv.ParseUint(match[1], 10, 32)
			if err != nil {
				return Exercise.InternalError(err.Error())
			}
			month, err := strconv.ParseUint(match[2], 10, 32)
			if err != nil {
				return Exercise.InternalError(err.Error())
			}
			if month < 1 || month > 12 || year == 0 {
				invalidFound = true
			} else {
				if year%4 != 0 || (year%100 == 0 && year%400 != 0) {
					commonYearsFound[month] = true
				} else {
					leapYearsFound[month] = true
				}
			}
		}
	}
	if !invalidFound {
		return missingTests("no test for invalid month")
	}
	for _, found := range leapYearsFound {
		if !found {
			return missingTests("missing test(s) for num_days_in_month function (leap years)")
		}
	}
	for _, found := range commonYearsFound {
		if !found {
			return missingTests("missing test(s) for num_days_in_month function (common years)")
		}
	}
	return Exercise.Passed("")
}

func testExistenceTest(exercise *Exercise.Exercise) Exercise.Result {
	content, err := os.ReadFile(exercise.TurnInFiles[0])
	if err != nil {
		return Exercise.InternalError(fmt.Sprintf("error reading file: %v", err.Error()))
	}
	if !strings.Contains(string(content), "is_leap_year(1600)") || !strings.Contains(string(content), "is_leap_year(1500)") {
		return missingTests("missing test(s) for is_leap_year function")
	}
	if !strings.Contains(string(content), "is_leap_year(2004)") || !strings.Contains(string(content), "is_leap_year(2003)") {
		return missingTests("missing test(s) for is_leap_year function")
	}
	if !strings.Contains(string(content), "is_leap_year(0)") {
		return missingTests("missing test(s) for is_leap_year function")
	}
	return monthTestExistenceTest(string(content))
}

func expectedMainOutput() string {
	currentDate := time.Now()
	year, month, day := currentDate.Date()

	var result strings.Builder
	for i := 1; i <= year; i++ {
		maxMonth := time.December
		if i == year {
			maxMonth = month
			if day < 13 {
				maxMonth--
			}
		}
		for j := time.January; j <= maxMonth; j++ {
			date := time.Date(i, j, 13, 0, 0, 0, 0, time.UTC)
			if date.Weekday() == time.Friday {
				result.WriteString(fmt.Sprintf("Friday, %s 13, %d\n", j, i))
			}
		}
	}
	return result.String()
}

func ex05Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, "", map[string]int{"unsafe": 0}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	/*
		if result := testExistenceTest(exercise); !result.Passed {
			return result
		}
	*/
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test"})
	if err != nil {
		logger.Exercise.Printf("%v", err)
		return Exercise.Result{Passed: false, Output: err.Error()}
	}
	if err := testutils.AppendStringToFile(CargoTestEx05, exercise.TurnInFiles[0]); err != nil {
		return Exercise.InternalError(err.Error())
	}
	_, err = testutils.RunCommandLine(workingDirectory, "cargo", []string{"test"})
	if err != nil {
		logger.Exercise.Printf("%v", err)
		return Exercise.Result{Passed: false, Output: err.Error()}
	}
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run"})
	if err != nil {
		logger.Exercise.Printf("%v", err)
		return Exercise.Result{Passed: false, Output: err.Error()}
	}
	expectedOutput := expectedMainOutput()
	if output != expectedOutput {
		return Exercise.AssertionError(expectedOutput, output)
	}
	return Exercise.Passed("OK")
}

func ex05() Exercise.Exercise {
	return Exercise.NewExercise("05", "ex05", []string{"src/main.rs", "Cargo.toml"}, 15, ex05Test)
}
