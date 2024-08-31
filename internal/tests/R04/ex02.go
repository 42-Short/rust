package R04

import (
	"io/fs"
	"os"
	"os/exec"
	"path/filepath"
	"regexp"
	"rust-piscine/internal/alloweditems"
	"strings"
	"sync"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
)

var clippyTomlAsString02 = ``

type output struct {
	out []byte
	err error
}

func removeANSICodes(input string) string {
	ansiPattern := `\x1b\[[0-9;]*[a-zA-Z]`

	re := regexp.MustCompile(ansiPattern)

	cleanedString := re.ReplaceAllString(input, "")
	return cleanedString
}

func testMb(workingDirectory string, fooPath string) Exercise.Result {
	kbString := strings.Repeat("69", 500000)
	if err := os.WriteFile(filepath.Join(fooPath, "69.txt"), []byte(kbString), fs.FileMode(os.O_WRONLY)); err != nil {
		return Exercise.InternalError(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan output)
	go func() {
		defer wg.Done()
		cmd := exec.Command("cargo", "run", "./foo")
		cmd.Dir = workingDirectory
		out, err := cmd.CombinedOutput()
		ch <- output{out, err}
	}()
	out := <-ch
	wg.Wait()
	if out.err != nil {
		return Exercise.RuntimeError(out.err.Error())
	}
	codeOutput := strings.Split(string(out.out), "\n")
	if len(codeOutput) < 4 {
		return Exercise.AssertionError("1.0 megabytes", strings.Join(codeOutput, "\n"), "ls -lh ./foo\n-rw-r--r-- 1 root root 1000000 Aug 31 17:15 69.txt", "cargo run ./foo")
	}
	if !strings.Contains(codeOutput[3], "1.0 megabytes") {
		return Exercise.AssertionError("1.0 megabytes", removeANSICodes(codeOutput[3]), "ls -lh ./foo\n-rw-r--r-- 1 root root 1000000 Aug 31 17:15 69.txt", "cargo run ./foo")
	}
	return Exercise.Passed("OK")
}

func testRecursive(workingDirectory string, fooPath string) Exercise.Result {
	if err := os.Mkdir(filepath.Join(fooPath, "bar"), 0755); err != nil {
		return Exercise.InternalError(err.Error())
	}
	mbString := strings.Repeat("69", 500000)
	if err := os.WriteFile(filepath.Join(fooPath, "69.txt"), []byte(mbString), fs.FileMode(os.O_WRONLY)); err != nil {
		return Exercise.InternalError(err.Error())
	}
	if err := os.WriteFile(filepath.Join(fooPath, "bar", "69.txt"), []byte(mbString), fs.FileMode(os.O_WRONLY)); err != nil {
		return Exercise.InternalError(err.Error())
	}
	var wg sync.WaitGroup
	wg.Add(1)

	ch := make(chan output)
	go func() {
		defer wg.Done()
		cmd := exec.Command("cargo", "run", "./bar")
		cmd.Dir = workingDirectory
		out, err := cmd.CombinedOutput()
		ch <- output{out, err}
	}()
	out := <-ch
	wg.Wait()
	if out.err != nil {
		return Exercise.RuntimeError(out.err.Error())
	}
	codeOutput := strings.Split(string(out.out), "\n")
	if len(codeOutput) < 5 {
		return Exercise.AssertionError("2.0 megabytes", strings.Join(codeOutput, "\n"), "ls -lhR ./bar\nbar:\n-rw-r--r-- 1 root root 977K Aug 31 17:15 69.txt\ndrwxr-xr-x 2 root root 4.0K Aug 31 18:12 bar\n\nbar/bar:\n-rw-r--r-- 1 root root 977K Aug 31 18:12 69.txt\n", "cargo run ./bar")
	}
	if !strings.Contains(codeOutput[4], "2.0 megabytes") {
		return Exercise.AssertionError("2.0 megabytes", removeANSICodes(codeOutput[4]), "ls -lhR ./bar\nbar:\n-rw-r--r-- 1 root root 977K Aug 31 17:15 69.txt\ndrwxr-xr-x 2 root root 4.0K Aug 31 18:12 bar\n\nbar/bar:\n-rw-r--r-- 1 root root 977K Aug 31 18:12 69.txt\n", "cargo run ./bar")
	}
	return Exercise.Passed("OK")
}

func ex02Test(exercise *Exercise.Exercise) (result Exercise.Result) {
	if err := alloweditems.Check(*exercise, clippyTomlAsString02, nil); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	fooPath := filepath.Join(workingDirectory, "foo")
	if err := os.Mkdir(fooPath, 0755); err != nil {
		return Exercise.InternalError(err.Error())
	}
	defer os.RemoveAll(fooPath)
	if result = testMb(workingDirectory, fooPath); !result.Passed {
		return result
	}
	fooPath = filepath.Join(workingDirectory, "bar")
	if err := os.Mkdir(fooPath, 0755); err != nil {
		return Exercise.InternalError(err.Error())
	}
	defer os.RemoveAll(fooPath)
	if result = testRecursive(workingDirectory, fooPath); !result.Passed {
		return result
	}
	return Exercise.Passed("OK")
}

func ex02() Exercise.Exercise {
	return Exercise.NewExercise("02", "ex02", []string{"Cargo.toml", "src/main.rs"}, 10, ex02Test)
}
