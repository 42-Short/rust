# rust
Rust Piscine built using [Nils Mathieu's subjects](https://github.com/nils-mathieu/piscine-rust) and the [shortinette framework](https://pkg.go.dev/github.com/42-Short/shortinette).

## Usage
Everything you need to know for environment and test configuration is documented [here](https://pkg.go.dev/github.com/42-Short/shortinette).

Since shortinette is supposed to be a language-agnostic framework for bootcamp/piscine organisation, enforcing of allowed items (functions, macros) could not be abstracted, since it is highly language-specific.
In the case of Rust, we happen to have a pre-made solution: [**clippy**](https://rust-lang.github.io/rust-clippy/master/index.html).

The wrapper for this can be found [here](internal/alloweditems/alloweditems.go).

