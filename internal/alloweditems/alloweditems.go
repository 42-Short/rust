package alloweditems

import (
	"os"
	"path/filepath"
	"slices"
	"strings"

	Exercise "github.com/42-Short/shortinette/pkg/interfaces/exercise"
	"github.com/42-Short/shortinette/pkg/logger"
	"github.com/42-Short/shortinette/pkg/testutils"
)

func prependLintLevel(filePath string, lintLevelModifications []string) (err error) {
	contentAsBytes, err := os.ReadFile(filePath)
	if err != nil {
		return err
	}

	contentAsStringSlice := strings.Split(string(contentAsBytes), "\n")
	for index, line := range contentAsStringSlice {
		if !strings.HasPrefix(line, "#![") {
			for _, modification := range lintLevelModifications {
				contentAsStringSlice = slices.Insert(contentAsStringSlice, index, modification)
			}
			break
		}
	}

	err = os.WriteFile(filePath, []byte(strings.Join(contentAsStringSlice, "\n")), 0644)
	if err != nil {
		return err
	}

	logger.Info.Printf("content of %s:\n%s", filePath, strings.Join(contentAsStringSlice, "\n"))

	return nil
}

// Checks for forbidden methods/macros using `cargo clippy`.
//
// Args:
// exercise: `Exercise.Exercise` structure containing the exercise metadata
//
// clippyTomlAsString: string representation of the `.clippy.toml file which should dictate the lint rules`
//
// lintLevelModifications: slice of strings representing the lint level modifications you want to add
//
// Example Usage:
//   - I want to ban `std::ptr::read` and `std::println`
//   - I also want to remove the documentation linter, which would fail students
//     for wrong indentation in comments (https://rust-lang.github.io/rust-clippy/master/index.html#doc_lazy_continuation)
//
// To achieve this, I can call call this function like follows:
//
//	var clippyTomlAsString := `
//	disallowed-macros = ["std::println"]
//	disallowed-methods = ["std::ptr::read"]
//	`
//	lintLevelModifications := []string{"#[allow(clippy::doc_lazy_continuation)]"}
//
//	if err := allowedItems.Check(exercise, clippyTomlAsString, lintLevelModifications); err != nil {
//		// err != nil -> linting failed, meaning the submission did not pass your static analysis.
//		// err.Error() will contain all necessary information for your trace, such as which line posed an issue,
//		// which disallowed item(s) was/were found, (...), you can simply handle this as follows:
//		return Exercise.CompilationError(err.Error())
//	}
//
// See https://rust-lang.github.io/rust-clippy/master/index.html for details.
func Check(exercise Exercise.Exercise, clippyTomlAsString string, lintLevelModifications []string) (err error) {
	for _, filePath := range exercise.TurnInFiles {
		if strings.Contains(filePath, ".rs") {
			if err = prependLintLevel(filePath, lintLevelModifications); err != nil {
				return err
			}
		}
	}
	file, err := os.Create(filepath.Join("studentcode", exercise.TurnInDirectory, ".clippy.toml"))
	if err != nil {
		return err
	}
	if _, err = file.WriteString(clippyTomlAsString); err != nil {
		return err
	}
	workingDirectory := filepath.Join("studentcode", exercise.TurnInDirectory)
	if _, err := testutils.RunCommandLine(workingDirectory, "cargo", []string{"clippy", "--", "-D", "warnings"}); err != nil {
		return err
	}
	return nil
}
