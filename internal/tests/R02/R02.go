package R02

import (
	Module "github.com/42-Short/shortinette/pkg/interfaces/module"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"

	"path/filepath"
	"github.com/42-Short/shortinette/pkg/testutils"
	"github.com/42-Short/shortinette/pkg/logger"
	"rust-piscine/internal/alloweditems"
	"time"
)

func runDefaultTest(exercise *Exercise.Exercise, cargoTestModAsString string, clippyTomlAsString string) Exercise.Result {
	if err := alloweditems.Check(*exercise, clippyTomlAsString, nil); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	if err := testutils.AppendStringToFile(cargoTestModAsString, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"valgrind", "test"}, testutils.WithTimeout(60 * time.Second)) //TODO: maybe adjust the time
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func R02() *Module.Module {
	exercises := map[string]Exercise.Exercise{
		"00": ex00(),
		"01": ex01(),
		"02": ex02(),
		"03": ex03(),
		"04": ex04(),
	}
	r02 := Module.NewModule("02", 70, exercises, "subject.md") // TODO: add actual subject, add actual minimum grade
	return &r02
}
