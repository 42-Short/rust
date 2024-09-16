package R03

import (
	"path/filepath"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var Ex07TestMod = `
[cfg(test)]
mod shortinette_rust_test_module03_ex07_0001 {
    use super::*;

    #[derive(Debug, PartialEq)]
    struct User {
        name: String,
        age: u32,
    }

    impl Record for User {
        fn encode(&self, target: &mut String) -> Result<(), EncodingError> {
            self.name.encode(target)?;
            target.push(',');
            self.age.encode(target)?;

            Ok(())
        }

        fn decode(line: &str) -> Result<Self, DecodingError> {
            let mut split = line.split(',');

            let name = match split.next() {
                Some(value) => Field::decode(value),
                None => Err(DecodingError),
            }?;

            let age = match split.next() {
                Some(value) => Field::decode(value),
                None => Err(DecodingError),
            }?;

            if split.next().is_some() {
                return Err(DecodingError);
            }

            Ok(Self { name, age })
        }
    }

    #[test]
    fn test_encode() {
        let database = [
            User {
                name: "aaa".into(),
                age: 23,
            },
            User {
                name: "bb".into(),
                age: 2,
            },
        ];

        let csv = encode_csv(&database).unwrap();

        assert_eq!(
            csv,
            "\
        aaa,23\n\
        bb,2\n\
        "
        );
    }

    #[test]
    fn test_decode() {
        let csv = "\
        hello,2\n\
        yes,5\n\
        no,100\n\
    ";

        let database: Vec<User> = decode_csv(csv).unwrap();

        assert_eq!(
            database,
            [
                User {
                    name: "hello".into(),
                    age: 2
                },
                User {
                    name: "yes".into(),
                    age: 5
                },
                User {
                    name: "no".into(),
                    age: 100
                },
            ]
        );
    }

    #[test]
    fn decoding_error() {
        let csv = "\
        hello,2\n\
        yes,6\n\
        no,23,hello\n\
    ";

        decode_csv::<User>(csv).unwrap_err();

        let csv = "\
        hello,2\n\
        yes,6\n\
        no\n\
    ";

        decode_csv::<User>(csv).unwrap_err();
    }

    #[test]
    fn empty_line() {
        let csv = "\
        hello,2\n\
        \n\
        bye,3\n\
        ";

        decode_csv::<User>(csv).unwrap_err();
    }
}

`

// TODO: since this is a bonus this should only be included if the macro actually exists
var Ex07TestModMacro = `
#[cfg(test)]
mod shortinette_rust_test_module03_ex07_0002 {
    use super::*;

    struct MyType {
        id: u32,
        name: String,
    }

    impl_record!(MyType(id, name));

    #[test]
    fn test_impl_record() {
        let records = [
            MyType {
                id: 10,
                name: "Marvin".into(),
            },
            MyType {
                id: 11,
                name: "Marvin".into(),
            },
            MyType {
                id: 12,
                name: "Marvin".into(),
            },
        ];

        let csv = encode_csv(&records).unwrap();
        assert_eq!(
            csv,
            "\
        10,Marvin\n\
        11,Marvin\n\
        12,Marvin\n\
        "
        );
    }
}
`

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(Ex07TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test", "--release", "shortinette_rust_test_module03_ex07_0001"})
	if err != nil {
		return Exercise.AssertionError("", output)
	}
	return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "ex07", []string{"src/lib.rs", "Cargo.toml"}, 25, ex07Test)
}
