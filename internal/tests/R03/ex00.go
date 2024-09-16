package R03

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex00TestMod = `
#[cfg(test)]
mod shortinette_rust_test_module03_ex00_0001 {
    use super::*;

    #[test]
    #[should_panic]
    fn empty() {
        let empty: Vec<i32> = Vec::new();

        let value = choose(&empty);
        println!(
            "What in the world? This should not be printed! Value: {:?}",
            value
        );
    }

    #[test]
    fn single() {
        let alone = &[0];

        let value = choose(alone);
        let value2 = choose(alone);

        assert_eq!(value, value2, "How is choose(&[0]) not returning 0?");
    }

    #[test]
    fn generic() {
        let numbers = [1, 2, 3];
        choose(&numbers);

        let slices = ["a", "b", "c"];
        choose(&slices);

        let bools = [true, false];
        choose(&bools);

        struct Foo;
        let foos = [Foo, Foo];
        choose(&foos);
    }

    #[test]
    fn randomness() {
        let huge: Vec<_> = (0..100_000).collect();

        // Well this is one of the cases where a second grademe could actually
        // pass the exercise.
        // Although it is very unlikely that this will ever happen.
        let value = choose(&huge);
        let value2 = choose(&huge);

        assert_ne!(
            value, value2,
            "choose(&huge) returned the same value twice. Do you really return a random element?"
        );
    }
}
`

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex00TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "ex00", []string{"src/main.rs", "Cargo.toml"}, 25, ex00Test)
}
