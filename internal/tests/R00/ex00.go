package R00

import (
	"bytes"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"

	"github.com/42-Short/shortinette/pkg/logger"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

func ex00Compile(exercise *Exercise.Exercise) error {
	cmd := exec.Command("rustc", filepath.Base(exercise.TurnInFiles[0]))
	dirPath := filepath.Dir(exercise.TurnInFiles[0])
	if _, err := os.Stat(dirPath); os.IsNotExist(err) {
		return err
	}
	cmd.Dir = filepath.Dir(exercise.TurnInFiles[0])

	output, err := cmd.CombinedOutput()
	if err != nil {
		logger.Exercise.Println(err)
		return fmt.Errorf("could not compile: %s", output)
	}
	logger.Exercise.Printf("%s compiled with rustc\n", exercise.TurnInFiles[0])
	return nil
}

func runExecutable(executablePath string) (string, error) {
	cmd := exec.Command("./" + filepath.Base(executablePath))
	cmd.Dir = filepath.Dir(executablePath)
	var stdout bytes.Buffer
	var stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	if err := cmd.Run(); err != nil {
		return stderr.String(), err
	}
	return stdout.String(), nil
}

func ex00Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := ex00Compile(exercise); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	executablePath := strings.TrimSuffix(exercise.TurnInFiles[0], filepath.Ext(exercise.TurnInFiles[0]))
	output, err := runExecutable(executablePath)
	if err != nil {
		return Exercise.RuntimeError(err.Error())
	}
	if output != "Hello, World!\n" {
		return Exercise.AssertionError("Hello, World!\n", output)
	}
	return Exercise.Passed("OK")
}

func ex00() Exercise.Exercise {
	return Exercise.NewExercise("00", "studentcode", "ex00", []string{"hello.rs"}, map[string]int{"unsafe": 0}, 10, ex00Test)
}
