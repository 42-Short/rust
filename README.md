# rust
Rust Piscine built using [Nils Mathieu's subjects](https://github.com/nils-mathieu/piscine-rust) and the [shortinette framework](https://pkg.go.dev/github.com/42-Short/shortinette).

## Usage
Everything you need to know for environment and test configuration is documented [here](https://pkg.go.dev/github.com/42-Short/shortinette).

Since shortinette is supposed to be a language-agnostic framework for bootcamp/piscine organisation, enforcing of allowed items (functions, macros) could not be abstracted, since it is highly language-specific.
In the case of Rust, we happen to have a pre-made solution: [**clippy**](https://rust-lang.github.io/rust-clippy/master/index.html).

The wrapper for this can be found [here](~/internal/alloweditems/alloweditems.go). We did not find any robust & simple way to exclusively allow the items from each exercise, which is why we went with Clippy, which bans specific items. This way, we can also point students in a direction without the need to ban EVERYTHING which is not specifically allowed - enabling them to use their creativity and solve problems in their unique way.
The idea is to only ban the items which would do the heavy lifting for them in order to enforce the use of lower level methods.

### Example Usage of alloweditems.Check()
  - I want to ban `std::ptr::read` and `std::println`
  - I also want to remove the documentation linter, which would fail students
    for wrong indentation in comments (https://rust-lang.github.io/rust-clippy/master/index.html#doc_lazy_continuation)

To achieve this, I can call call alloweditems.Check() like follows:

	var clippyTomlAsString := `
	disallowed-macros = ["std::println"]
	disallowed-methods = ["std::ptr::read"]
	`
	lintLevelModifications := []string{"#[allow(clippy::doc_lazy_continuation)]"}

	if err := allowedItems.Check(exercise, clippyTomlAsString, lintLevelModifications); err != nil {
		// err != nil -> linting failed, meaning the submission did not pass your static analysis.
		// err.Error() will contain all necessary information for your trace, such as which line posed an issue,
		// which disallowed item(s) was/were found, (...), you can simply handle this as follows:
		return Exercise.CompilationError(err.Error())
	}

See https://rust-lang.github.io/rust-clippy/master/index.html for details on .clippy.toml configuration.

