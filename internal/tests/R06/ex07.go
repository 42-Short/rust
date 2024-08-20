package R06

import (
	"fmt"
	"path/filepath"
	"strings"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

// TODO:
// Add tests for the #[no_std] & #[no_main] requirements
func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
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
