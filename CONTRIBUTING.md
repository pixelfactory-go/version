# Contributing to version

Thank you for your interest in contributing to the version package! We welcome contributions from the community.

## Code of Conduct

By participating in this project, you agree to maintain a respectful and inclusive environment for all contributors.

## How to Contribute

### Reporting Issues

- Check if the issue has already been reported
- Use the GitHub issue tracker to report bugs
- Provide clear reproduction steps
- Include relevant version information

### Submitting Pull Requests

1. Fork the repository
2. Create a new branch for your feature or bugfix (`git checkout -b feature/my-new-feature`)
3. Make your changes
4. Write or update tests as needed
5. Ensure all tests pass (`make test`)
6. Run linters (`make lint`)
7. Format your code (`make fmt`)
8. Commit your changes with clear commit messages
9. Push to your fork
10. Submit a pull request

## Development Setup

### Prerequisites

- Go 1.21 or higher
- golangci-lint (for linting)

### Getting Started

```bash
# Clone the repository
git clone https://github.com/pixelfactory-go/version.git
cd version

# Install dependencies
go mod download

# Run tests
make test

# Run linters
make lint

# Check formatting
make fmt
```

## Coding Standards

### Go Style Guide

- Follow the [Effective Go](https://golang.org/doc/effective_go) guidelines
- Use `gofmt` for code formatting
- Write clear, idiomatic Go code
- Add comments for exported functions and types

### Testing

- Write unit tests for new functionality
- Maintain or improve test coverage
- Use table-driven tests where appropriate
- Include both positive and negative test cases

### Commit Messages

Follow the Conventional Commits specification:

- `feat:` - New features
- `fix:` - Bug fixes
- `docs:` - Documentation changes
- `refactor:` - Code refactoring
- `test:` - Test additions or modifications
- `chore:` - Maintenance tasks

Examples:
```
feat: add new GetVersion function
fix: correct BUILDDATE format handling
docs: update README installation section
test: add edge case tests for REVISION
```

## Pull Request Process

1. Update the README.md with details of changes if applicable
2. Update tests to reflect your changes
3. Ensure CI checks pass
4. Request review from maintainers
5. Address review feedback
6. Once approved, maintainers will merge your PR

## Testing Guidelines

### Running Tests

```bash
# Run all tests
make test

# Run tests with coverage
go test -v -race -coverprofile coverage.txt -covermode atomic ./...

# View coverage report
go tool cover -html=coverage.txt
```

### Writing Tests

- Place tests in `*_test.go` files
- Use the `testing` package
- Write clear test names that describe what is being tested
- Use subtests for related test cases

Example:
```go
func TestFeature(t *testing.T) {
    t.Run("positive case", func(t *testing.T) {
        // Test code
    })

    t.Run("negative case", func(t *testing.T) {
        // Test code
    })
}
```

## Documentation

- Keep README.md up to date
- Add godoc comments for exported functions and types
- Include usage examples
- Update CHANGELOG.md for significant changes

## Questions?

If you have questions about contributing, feel free to:
- Open an issue for discussion
- Reach out to the maintainers

## License

By contributing, you agree that your contributions will be licensed under the MIT License.
