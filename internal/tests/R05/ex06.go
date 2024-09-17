package R05

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex06TestMod = `
[cfg(test)]
mod shortinette_rust_test_module05_ex06_0001 {}
`

// TODO: this needs to test with different mains
func ex06Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex06TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test", "--release", "shortinette_rust_test_module05_ex06_0001"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "ex06", []string{"src/main.rs", "Cargo.toml"}, 25, ex06Test)
}
