package R06

import (
	"fmt"
	"path/filepath"
	"strings"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	fullFilePath := filepath.Join(exercise.RepoDirectory, exercise.TurnInDirectory, exercise.TurnInFiles[0])
	if output, err := testutils.RunCommandLine(filepath.Dir(fullFilePath), "rustc", []string{"-C panic=abort", "-C link-args=-nostartfiles", "-o ft_putchar", "ft_putchar.rs"}); err != nil {
		return Exercise.CompilationError(fmt.Sprintf("%s: %s", err.Error(), output))
	}
	executablePath := strings.TrimSuffix(exercise.TurnInFiles[0], ".rs")

	output, err := testutils.RunExecutable(executablePath)
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	
	if output != "42" {
		return Exercise.AssertionError("42", output)
	}

	if output, _ := testutils.RunCommandLine(".", "echo", []string{"$?"}); output != "42" {
		return Exercise.RuntimeError(fmt.Sprintf("expected exit code '42', got '%s'", output))
	}
 
	return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "studentcode", "ex07", []string{"ft_putchar.rs"}, []string{"std::mem::MaybeUninit", "std::ffi::CStr", "std::ffi::{c_int, c_char}"}, map[string]int{"unsafe": 100}, 10, ex07Test)
}
