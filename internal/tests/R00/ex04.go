package R00

import (
	"fmt"
	"os"
	"path/filepath"
	"rust-piscine/internal/alloweditems"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/testutils"
	"github.com/pelletier/go-toml/v2"
)

type Config struct {
	Package PackageConfig
}

type PackageConfig struct {
	Name        string
	Edition     string
	Description string
	Authors     []string
	Publish     *bool
}

type ReleaseMode int

const (
	Debug   ReleaseMode = 0
	Release ReleaseMode = 1
)

// var expectedTomlContent = map[string]string{
// 	"package.name":        "module00-ex04",
// 	"package.edition":     "2021",
// 	"package.description": "my answer to the fifth exercise of the first module of 42's Rust Piscine",
// }

func testNmReleaseMode(exercise *Exercise.Exercise, binary string, releaseMode ReleaseMode) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	output, err := testutils.RunCommandLine(workingDirectory, "nm", []string{binary})
	if err != nil {
		return Exercise.Result{Passed: false, Output: fmt.Sprintf("runtime error: nm did not execute as expected: %v", err)}
	}
	if releaseMode == Release && output != "" {
		return Exercise.AssertionError("", output)
	} else if releaseMode == Debug && output == "" {
		return Exercise.AssertionError("<non_empty>", output)
	}
	return Exercise.Passed("")
}

func testCargoRunBinOtherReleaseMode(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--release", "--bin", "other"})
	if err != nil {
		return Exercise.Result{Passed: false, Output: fmt.Sprintf("runtime error: cargo run: %s", err)}
	}
	if output != "Hey! I'm the other bin target!\nI'm in release mode!\n" {
		return Exercise.AssertionError("Hey! I'm the other bin target!\nI'm in release mode!\n", output)
	}
	return testNmReleaseMode(exercise, "target/release/other", Release)
}

func testCargoRunBinOther(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--bin", "other"})
	if err != nil {
		return Exercise.Result{Passed: false, Output: fmt.Sprintf("runtime error: %v", err)}
	}
	if output != "Hey! I'm the other bin target!\n" {
		return Exercise.AssertionError("Hey! I'm the other bin target!\n", output)
	}
	return testNmReleaseMode(exercise, "target/debug/other", Debug)
}

func testCargoRun(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run"})
	if err != nil {
		return Exercise.Result{Passed: false, Output: fmt.Sprintf("runtime error: %v", err)}
	}
	if output != "Hello, Cargo!\n" {
		return Exercise.AssertionError("Hello, Cargo!\n", output)
	}
	return testNmReleaseMode(exercise, "target/debug/module00-ex04", Debug)
}

func testOverflow(exercise *Exercise.Exercise) Exercise.Result {
	workingDirectory := filepath.Join(exercise.CloneDirectory, exercise.TurnInDirectory)
	_, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--bin", "test-overflows"})
	if err == nil {
		return Exercise.Result{Passed: false, Output: "test-overflows should panic in debug mode"}
	}
	output, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"run", "--profile", "no-overflows", "--bin", "test-overflows"})
	if err != nil {
		return Exercise.Result{Passed: false, Output: "execution error with 'no-overflows' profile"}
	}
	if output != "255u8 + 1u8 == 0\n" {
		return Exercise.AssertionError("255u8 + 1u8 == 0\n", output)
	}
	return testNmReleaseMode(exercise, "target/no-overflows/test-overflows", Debug)
}

func cargoTomlTest(filename string) Exercise.Result {
	var cargoToml Config

	content, err := os.ReadFile(filename)
	if err != nil {
		return Exercise.InternalError(fmt.Sprintf("error reading cargo.toml: %v", err))
	}
	err = toml.Unmarshal(content, &cargoToml)
	if err != nil {
		return Exercise.InternalError(fmt.Sprintf("error parsing cargo.toml: %v", err))
	}
	if cargoToml.Package.Name != "module00-ex04" {
		return Exercise.AssertionError("module00-ex04", cargoToml.Package.Name)
	}
	if cargoToml.Package.Edition != "2021" {
		return Exercise.AssertionError("2021", cargoToml.Package.Edition)
	}
	if cargoToml.Package.Description != "my answer to the fifth exercise of the first module of 42's Rust Piscine" {
		return Exercise.AssertionError("my answer to the fifth exercise of the first module of 42's Rust Piscine", cargoToml.Package.Description)
	}
	if len(cargoToml.Package.Authors) != 1 || cargoToml.Package.Authors[0] == "" {
		return Exercise.Result{Passed: false, Output: "author check failed in cargo.toml"}
	}
	if cargoToml.Package.Publish == nil || *cargoToml.Package.Publish {
		return Exercise.AssertionError("false", fmt.Sprintf("%v", cargoToml.Package.Publish))
	}
	return Exercise.Passed("")
}

func ex04Test(exercise *Exercise.Exercise) Exercise.Result {
	if err := alloweditems.Check(*exercise, "", map[string]int{"unsafe": 0}); err != nil {
		return Exercise.CompilationError(err.Error())
	}
	if result := cargoTomlTest(exercise.TurnInFiles[3]); !result.Passed {
		return result
	}
	if result := testCargoRun(exercise); !result.Passed {
		return result
	}
	if result := testCargoRunBinOther(exercise); !result.Passed {
		return result
	}
	if result := testCargoRunBinOtherReleaseMode(exercise); !result.Passed {
		return result
	}
	if result := testOverflow(exercise); !result.Passed {
		return result
	}
	return Exercise.Passed("OK")
}

func ex04() Exercise.Exercise {
	return Exercise.NewExercise("04", "ex04", []string{"src/main.rs", "src/overflow.rs", "src/other.rs", "Cargo.toml"}, 10, ex04Test)
}
