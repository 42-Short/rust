package R00

import (
	"fmt"
	"path/filepath"
	"rust-piscine/internal/alloweditems"
	"strings"
	"time"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func doFizzBuzz() string {
	var result strings.Builder
	for number := 1; number <= 100; number++ {
		switch {
		case number%3 == 0 && number%5 == 0:
			result.WriteString("fizzbuzz\n")
		case number%3 == 0:
			result.WriteString("fizz\n")
		case number%5 == 0:
			result.WriteString("buzz\n")
		case number%11 == 3:
			result.WriteString("FIZZ\n")
		case number%11 == 5:
			result.WriteString("BUZZ\n")
		default:
			result.WriteString(fmt.Sprintf("%d\n", number))
		}
	}
	return result.String()
}

func fizzBuzzOutputTest(exercise *Exercise.Exercise) Exercise.Result {
	if err := CompileWithRustc(exercise.TurnInFiles[0]); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := testutils.ExecutablePath(exercise.TurnInFiles[0], filepath.Ext(exercise.TurnInFiles[0]))
	output, err := testutils.RunExecutable(executablePath, testutils.WithTimeout(500*time.Millisecond))
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	expectedOutput := doFizzBuzz()

	if output != expectedOutput {
		return Exercise.AssertionError(expectedOutput, output)
	}
	return Exercise.Passed("OK")
}

func clippyCheck03(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	if _, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"init"}); err != nil {
		return Exercise.InternalError("cargo init failed")
	}
	if _, err := testutils.RunCommandLine(workingDirectory, "cp", []string{"fizzbuzz.rs", "src/main.rs"}); err != nil {
		return Exercise.InternalError("unable to copy file to src/ folder")
	}
	tmp := Exercise.Exercise{
		CloneDirectory:  exercise.CloneDirectory,
		TurnInDirectory: exercise.TurnInDirectory,
		TurnInFiles:     []string{filepath.Join(workingDirectory, "src/main.rs")},
	}
	if err := alloweditems.Check(tmp, "", map[string]int{"unsafe": 0}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	return Exercise.Passed("OK")
}

func ex03Test(exercise *Exercise.Exercise) Exercise.Result {
	if result := clippyCheck03(exercise); !result.Passed {
		return result
	}
	return fizzBuzzOutputTest(exercise)
}

func ex03() Exercise.Exercise {
	return Exercise.NewExercise("03", "ex03", []string{"fizzbuzz.rs"}, 10, ex03Test)
}
