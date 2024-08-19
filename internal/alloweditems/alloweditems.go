package alloweditems

import (
	"os"
	"path/filepath"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
)

func Execute(exercise Exercise.Exercise) (err error) {
	if err = os.Remove(filepath.Join("compile-environment", exercise.TurnInDirectory, "Cargo.toml")); err != nil {
		return err
	}
	logger.Info.Printf("no forbidden items/keywords found in %s", exercise.TurnInDirectory+"/"+exercise.TurnInFiles[0])
	return nil
}
