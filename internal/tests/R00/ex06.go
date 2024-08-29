package R00

import (
	"path/filepath"
	"strings"
	"time"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func ex06Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run"}, testutils.WithTimeout(5*time.Second))
	if err != nil && !strings.Contains(err.Error(), "EOF reached") {
		logger.Exercise.Printf("%v", err)
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "ex06", []string{"src/main.rs", "Cargo.toml"}, 25, ex06Test)
}
