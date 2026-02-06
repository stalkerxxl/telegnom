# Contributing to Telegnom

First off, thank you for considering contributing to Telegnom! It's people like you that make Telegnom such a great tool.

## How Can I Contribute?

### Reporting Bugs
* Check the [Issues](https://github.com/stalkerxxl/telegnom/issues) to see if the bug has already been reported.
* If not, create a new issue using the **Bug Report** template.

### Suggesting Enhancements
* Open a new issue using the **Feature Request** template to discuss your idea before implementation.

### Pull Requests
1. **Fork** the repository.
2. Create a new branch for your changes: `git checkout -b feat/your-feature-name` or `git checkout -b fix/your-bug-name`.
3. Make your changes.
4. Ensure your code follows Go standards:
   * Run `go fmt ./...`
   * Run `go mod tidy`
5. **Run tests**: Make sure all tests pass by running `go test ./...`.
6. Commit your changes: `git commit -m "feat: add some amazing feature"`.
7. Push to your fork and **submit a Pull Request to the `main` branch**.

## Development Rules
* **No external dependencies**: We aim to keep the library dependent only on the Go standard library.
* **Documentation**: All public functions and types should have comments following [GoDoc](https://go.dev/doc/comment) standards.
* **Be descriptive**: Explain *why* the change is necessary in your PR description.

## License
By contributing, you agree that your contributions will be licensed under the [MIT License](LICENSE).
