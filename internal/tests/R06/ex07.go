package R06

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func checkGlobalAttributesUsage(exercise *Exercise.Exercise) (err error) {
	fileAsBytes, err := os.ReadFile(filepath.Join(exercise.RepoDirectory, exercise.TurnInDirectory, "ft_putchar.rs"))
	if err != nil {
		return err
	}
	missingAttributes := map[string]bool{
		"#![no_std]":  true,
		"#![no_main]": true,
	}
	fileAsSlice := strings.Split(string(fileAsBytes), "\n")

	for _, line := range fileAsSlice {
		delete(missingAttributes, line)
	}
	if len(missingAttributes) != 0 {
		return fmt.Errorf("missing global attributes: %v", missingAttributes)
	}
	return nil
}

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := checkGlobalAttributesUsage(exercise); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	if output, err := testutils.RunCommandLine(filepath.Dir(exercise.TurnInFiles[0]), "rustc", []string{"-C", "panic=abort", "-C", "link-args=-nostartfiles", "-o", "ft_putchar", "ft_putchar.rs"}); err != nil {
		return Exercise.CompilationError(fmt.Sprintf("%s: %s", err.Error(), output))
	}
	executablePath := strings.TrimSuffix(exercise.TurnInFiles[0], ".rs")

	output, err := testutils.RunExecutable(executablePath)
	if err != nil {
		if err.Error() != "exit status 42" {
			return Exercise.RuntimeError("invalid exit code:" + err.Error())
		}
	}

	if !strings.Contains(output, "42") {
		return Exercise.AssertionError("42", output)
	}

	return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "studentcode", "ex07", []string{"ft_putchar.rs"}, map[string]int{"unsafe": 100}, 10, ex07Test)
}
