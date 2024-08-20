package alloweditems

import (
	"os"
	"path/filepath"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func Execute(exercise Exercise.Exercise, clippyTomlAsString string) (err error) {
	file, err := os.Create(filepath.Join("compile-environment", exercise.TurnInDirectory, ".clippy.toml"))
	if err != nil {
		return err
	}
	if _, err = file.WriteString(clippyTomlAsString); err != nil {
		return err
	}
	workingDirectory := filepath.Join("compile-environment", exercise.TurnInDirectory)
	if output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"clippy", "--", "-D", "warnings"}); err != nil {
		
	}
	return nil
}
