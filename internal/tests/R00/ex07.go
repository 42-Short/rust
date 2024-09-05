package R00

import (
	"path/filepath"
	"rust-piscine/internal/alloweditems"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var TestMod = `
#[cfg(test)]
mod test {
	use crate::lib::strpcmp;

	#[test]
	fn test1() {
		assert!(strpcmp(b"abc", b"abc"));
	}

	#[test]
	fn test2() {
		assert!(strpcmp(b"abcd", b"ab*"));
	}
	
	#[test]
	fn test3() {
		assert!(!strpcmp(b"cab", b"ab*"));
	}
	
	#[test]
	fn test4() {
		assert!(strpcmp(b"dcab", b"*ab"));
	}
	
	#[test]
	fn test5() {
		assert!(!strpcmp(b"abc", b"*ab"));
	}

	#[test]
	fn test6() {
		assert!(strpcmp(b"ab000cd", b"ab*cd"));
	}
	
	#[test]
	fn test7() {
		assert!(strpcmp(b"abcd", b"ab*cd"));
	}

	#[test]
	fn test8() {
		assert!(strpcmp(b"", b"****"));
	}

	#[test]
	fn test9() {
		assert!(strpcmp(b"abc*def", b"abc*"));
	}

	#[test]
	fn test10() {
		assert!(strpcmp(b"abc**", b"abc*"));
	}

	#[test]
	fn test11() {
		assert_eq!(strpcmp(b"abc*", b"*abc"), false);
	}

	#[test]
	fn test12() {
		assert_eq!(strpcmp(b"ab*cd", b"abcd"), false);
	}
}
`

func ex07Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, "", map[string]int{"unsafe": 0}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if err := testutils.AppendStringToFile(TestMod, exercise.TurnInFiles[1]); err != nil {
		logger.Exercise.Printf("internal error: %v", err)
		return Exercise.InternalError(err.Error())
	}

	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"test"})
	if err != nil {
		return Exercise.AssertionError("", err.Error())
	}
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--", "abcde", "ab*"})
	if err != nil {
		return Exercise.AssertionError("", err.Error())
	}
	if output != "yes\n" {
		return Exercise.AssertionError("yes\n", output)
	}
	output, err = testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--", "abcde", "ab*ef"})
	if err != nil {
		return Exercise.AssertionError("", err.Error())
	}
	if output != "no\n" {
		return Exercise.AssertionError("no\n", output)
	}
	return Exercise.Passed("OK")
}

func ex07() Exercise.Exercise {
	return Exercise.NewExercise("07", "ex07", []string{"src/main.rs", "src/lib.rs", "Cargo.toml"}, 25, ex07Test)
}
