package R05

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex00TestMod = `
#[cfg(test)]
mod shortinette_rust_test_module05_ex00_0001 {
    use super::*;

    #[test]
    fn u32() {
        let a = Cell::new(1);
        let b = Cell::new(3);

        swap_u32(&a, &b);

        assert_eq!(a.get(), 3);
        assert_eq!(b.get(), 1);
    }

    #[test]
    fn string() {
        let a = Cell::new("ABC".into());
        let b = Cell::new("DEF".into());

        swap_string(&a, &b);

        assert_eq!(a.into_inner(), "DEF");
        assert_eq!(b.into_inner(), "ABC");
    }
}
`

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex00TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test", "--release", "shortinette_rust_test_module05_ex00_0001"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "ex00", []string{"src/main.rs", "Cargo.toml"}, 25, ex00Test)
}
