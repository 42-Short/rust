package R03

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex03TestMod = `
[cfg(test)]
mod shortinette_rust_test_module03_ex03_0001 {
	
}
`

// TODO: this needs to test with different mains
func ex03Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex03TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test", "--release", "shortinette_rust_test_module03_ex03_0001"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex03() Exercise.Exercise {
	return Exercise.NewExercise("03", "ex03", []string{"src/main.rs", "Cargo.toml"}, 25, ex03Test)
}
