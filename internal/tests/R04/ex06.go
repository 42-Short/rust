package R04

import (
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"rust-piscine/internal/alloweditems"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
)

var clippyTomlAsString06 = ``

func testSimpleCBinary(workingDirectory string) Exercise.Result {
	if err := os.WriteFile(filepath.Join(workingDirectory, "test.c"), []byte("int main(){return 0;}"), fs.FileMode(os.O_WRONLY)); err != nil {
		return Exercise.InternalError(err.Error())
	}
	if _, err := testutils.RunCommandLine(workingDirectory, "cc", []string{filepath.Join(workingDirectory, "test.c")}); err != nil {
		return Exercise.InternalError(err.Error())
	}
	commandLine := []string{"cargo", "run", filepath.Join(workingDirectory, "a.out")}
	output, err := testutils.RunCommandLine(workingDirectory, commandLine[0], commandLine[:1])
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	fmt.Println(output)
	return Exercise.Passed("OK")
}

func ex06Test(exercise *Exercise.Exercise) (result Exercise.Result) {
	if err := alloweditems.Check(*exercise, clippyTomlAsString06, nil); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)

	if result = testNoInput(workingDirectory); !result.Passed {
		return result
	}
	if result = testSimpleCBinary(workingDirectory); !result.Passed {
		return result
	}

	return Exercise.Passed("OK")
}

func ex06() Exercise.Exercise {
	return Exercise.NewExercise("06", "ex06", []string{"Cargo.toml", "src/main.rs"}, 15, ex06Test)
}
