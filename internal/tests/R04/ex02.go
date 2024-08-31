package R04

import (
	"os/exec"
	"path/filepath"
	"rust-piscine/internal/alloweditems"
	"sync"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var clippyTomlAsString02 = ``

func ex02Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, clippyTomlAsString02, nil); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	var wg sync.WaitGroup
	wg.Add(1)

	type output struct {
		out []byte
		err error
	}
	ch := make(chan output)
	go func() {
		defer wg.Done()
		cmd := exec.Command("cargo", "run", "..")
		cmd.Dir = workingDirectory
		out, err := cmd.CombinedOutput()
		ch <- output{out, err}
	}()
	x := <-ch
	wg.Wait()
	if x.err != nil {
		return Exercise.RuntimeError(x.err.Error())
	}
	return Exercise.Passed("OK")
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "ex02", []string{"Cargo.toml", "src/main.rs"}, 10, ex02Test)
}
