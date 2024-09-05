package R00

import (
	"errors"
	"fmt"
	"path/filepath"
	"rust-piscine/internal/alloweditems"
	"strings"
	"time"

	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

const YesMain = `
fn main() {
	yes();
}
`

const CollatzMain = `
fn main() {
	collatz(%s);
}
`

const PrintBytesMain = `
fn main() {
	print_bytes("%s")
}
`

func yes(filename string) Exercise.Result {
	if err := testutils.AppendStringToFile(YesMain, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	if err := testutils.CompileWithRustc(filename); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := testutils.ExecutablePath(filename, ".rs")
	output, err := testutils.RunExecutable(executablePath, testutils.WithTimeout(500*time.Millisecond))
	if err != nil && !errors.Is(err, testutils.ErrTimeout) {
		return Exercise.RuntimeError(err.Error())
	}
	lines := strings.Split(output, "\n")
	for _, line := range lines {
		if line != "y" && line != "" {
			return Exercise.AssertionError("y", line)
		}
	}
	if err := testutils.DeleteStringFromFile(YesMain, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	return Exercise.Passed("OK")
}

func collatzInfiniteLoopTest(filename string) Exercise.Result {
	main := fmt.Sprintf(CollatzMain, "0")
	if err := testutils.AppendStringToFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	if err := testutils.CompileWithRustc(filename); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := testutils.ExecutablePath(filename, ".rs")

	if _, err := testutils.RunExecutable(executablePath, testutils.WithTimeout(500*time.Millisecond)); err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	if err := testutils.DeleteStringFromFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	return Exercise.Result{Passed: true}
}

func doCollatz(n int) string {
	if n <= 0 {
		return ""
	}
	var results []string
	for n != 1 {
		results = append(results, fmt.Sprintf("%d", n))
		if n%2 == 0 {
			n /= 2
		} else {
			n = 3*n + 1
		}
	}
	results = append(results, "1")
	return strings.Join(results, "\n") + "\n"
}

func collatzAssertionTest(filename string) Exercise.Result {
	main := fmt.Sprintf(CollatzMain, "42")
	if err := testutils.AppendStringToFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	if err := testutils.CompileWithRustc(filename); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := testutils.ExecutablePath(filename, ".rs")
	output, err := testutils.RunExecutable(executablePath, testutils.WithTimeout(500*time.Millisecond))
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	expectedOutput := doCollatz(42)

	if output != expectedOutput {
		return Exercise.AssertionError(expectedOutput, output)
	}
	if err := testutils.DeleteStringFromFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	return Exercise.Passed("OK")
}

func collatz(filename string) Exercise.Result {
	if result := collatzInfiniteLoopTest(filename); !result.Passed {
		return result
	}
	if result := collatzAssertionTest(filename); !result.Passed {
		return result
	}
	return Exercise.Passed("OK")
}

func doPrintBytes(s string) string {
	var result strings.Builder
	for i := 0; i < len(s); i++ {
		result.WriteString(fmt.Sprintf("%d\n", s[i]))
	}
	return result.String()
}

func printBytesAssertionTest(filename string) Exercise.Result {
	main := fmt.Sprintf(PrintBytesMain, "Hello, World!")
	if err := testutils.AppendStringToFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	if err := testutils.CompileWithRustc(filename); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := testutils.ExecutablePath(filename, ".rs")
	output, err := testutils.RunExecutable(executablePath, testutils.WithTimeout(500*time.Millisecond))
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	expectedOutput := doPrintBytes("Hello, World!")

	if output != expectedOutput {
		return Exercise.AssertionError(expectedOutput, output)
	}
	if err := testutils.DeleteStringFromFile(main, filename); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}
	return Exercise.Passed("OK")
}

func printBytes(filename string) Exercise.Result {
	return printBytesAssertionTest(filename)
}

func clippyCheck02(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	if _, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"init", "--lib"}); err != nil {
		return Exercise.InternalError("cargo init failed")
	}
	for _, file := range exercise.TurnInFiles {
		if _, err := testutils.RunCommandLine(workingDirectory, "cp", []string{filepath.Base(file), "src/lib.rs"}); err != nil {
			return Exercise.InternalError("unable to copy file to src/ folder")
		}
		tmp := Exercise.Exercise{
			CloneDirectory:  exercise.CloneDirectory,
			TurnInDirectory: exercise.TurnInDirectory,
			TurnInFiles:     []string{filepath.Join(workingDirectory, "src/lib.rs")},
		}
		if err := alloweditems.Check(tmp, "", map[string]int{"unsafe": 0}); err != nil {
			return Exercise.CompilationError(err.Error())
		}
	}
	return Exercise.Passed("")
}

func ex02Test(exercise *Exercise.Exercise) Exercise.Result {
	if result := clippyCheck02(exercise); !result.Passed {
		return result
	}
	if result := yes(exercise.TurnInFiles[2]); !result.Passed {
		return result
	}
	if result := collatz(exercise.TurnInFiles[0]); !result.Passed {
		return result
	}
	if result := printBytes(exercise.TurnInFiles[1]); !result.Passed {
		return result
	}
	return Exercise.Passed("OK")
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "ex02", []string{"collatz.rs", "print_bytes.rs", "yes.rs"}, 20, ex02Test)
}
