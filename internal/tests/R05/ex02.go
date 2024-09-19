package R05

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex02TestMod = `
[cfg(test)]
mod shortinette_rust_test_module05_ex02_0001 {
    use std::thread;

    use super::*;

    #[test]
    fn last() {
        match Error::last() {
            Error::Success => (),
            _ => panic!("If Error::make_last() was never called, Error::last() call should return Error::Success")
        }
    }

    #[test]
    fn make_last() {
        {
            let err = Error::FileNotFound;
            err.make_last();
            match Error::last() {
                Error::FileNotFound => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        }

        {
            let err = Error::IsDirectory;
            err.make_last();
            match Error::last() {
                Error::IsDirectory => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        }

        {
            let err = Error::WriteError;
            err.make_last();
            match Error::last() {
                Error::WriteError => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        }

        {
            let err = Error::ReadError;
            err.make_last();
            match Error::last() {
                Error::ReadError => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        }
    }

    #[test]
    fn multiple_threads() {
        let err = Error::FileNotFound;
        err.make_last();

        let handle = thread::spawn(|| {
            match Error::last() {
                Error::Success => (),
                _ => panic!("Error::make_last() should not affect values across threads."),
            }

            let err = Error::ReadError;
            err.make_last();
            match Error::last() {
                Error::ReadError => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        });

        let handle2 = thread::spawn(|| {
            match Error::last() {
                Error::Success => (),
                _ => panic!("Error::make_last() should not affect values across threads."),
            }

            let err = Error::IsDirectory;
            err.make_last();
            match Error::last() {
                Error::IsDirectory => (),
                _ => panic!("Error::make_last() did not set the value correctly"),
            }
        });

        handle.join().unwrap();
        handle2.join().unwrap();

        match Error::last() {
            Error::FileNotFound => (),
            _ => panic!("Error::make_last() did not set the value correctly"),
        }
    }
}
`

func ex02Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex02TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test", "--release", "shortinette_rust_test_module05_ex02_0001"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "ex02", []string{"src/lib.rs", "Cargo.toml"}, 25, ex02Test)
}
