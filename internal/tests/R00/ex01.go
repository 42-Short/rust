package R00

import (
	"fmt"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

const CargoTest = `
#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_0() {
        assert_eq!(min(1, 2), 1);
    }

    #[test]
    fn test_1() {
        assert_eq!(min(2, 1), 1);
    }

    #[test]
    fn test_2() {
        assert_eq!(min(1, 1), 1);
    }

    #[test]
    fn test_3() {
        assert_eq!(min(-1, 0), -1);
    }
}
`

func compileWithRustcTestOption(turnInFile string) error {
	cmd := exec.Command("rustc", "--test", filepath.Base(turnInFile))
	cmd.Dir = filepath.Dir(turnInFile)

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Exercise.Println(err)
		return fmt.Errorf("could not compile: %s", output)
	}
	logger.Exercise.Printf("%s/%s compiled with rustc --test\n", cmd.Dir, turnInFile)
	return nil
}

func ex01Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := testutils.AppendStringToFile(CargoTest, exercise.TurnInFiles[0]); err != nil {
		logger.Exercise.Printf("could not write to %s: %v", exercise.TurnInFiles[0], err)
		return Exercise.InternalError(err.Error())
	}
	if err := compileWithRustcTestOption(exercise.TurnInFiles[0]); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	if output, err := testutils.RunExecutable(strings.TrimSuffix(exercise.TurnInFiles[0], ".rs")); err != nil {
		return Exercise.AssertionError("", err.Error()+output)
	}
	return Exercise.Passed("OK")
}

func ex01() Exercise.Exercise {
	return Exercise.NewExercise("01", "studentcode", "ex01", []string{"min.rs"}, []string{"println"}, map[string]int{"unsafe": 0}, 10, ex01Test)
}
