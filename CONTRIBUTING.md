# Contributing to gwt

Thank you for your interest in contributing to gwt! This document provides guidelines and instructions for contributing.

## Code of Conduct

Please be respectful and constructive in all interactions with other contributors.

## Getting Started

1. Fork the repository on GitHub
2. Clone your fork locally
3. Create a new branch for your feature or bug fix
4. Set up development environment:
   ```bash
   make deps
   ```

## Development Setup

```bash
# Build the tool
make build

# Run it
./dist/gwt help

# Install locally for testing
make install

# Run tests
make test

# Format code
make fmt

# Check for issues
make lint
```

## Making Changes

1. Create a branch with a descriptive name:
   ```bash
   git checkout -b feature/add-some-feature
   git checkout -b fix/resolve-issue-123
   ```

2. Make your changes following these guidelines:
   - Keep changes focused and atomic
   - Write clear commit messages
   - Add/update tests as needed
   - Update documentation if behavior changes

3. Test your changes:
   ```bash
   make build
   make test
   ```

4. Format and lint:
   ```bash
   make fmt
   make lint
   ```

## Project Structure

```
gwt/
├── main.go                 # Entry point
├── commands/               # Command implementations
├── internal/
│   ├── git/               # Git command wrappers
│   ├── config/            # Configuration management
│   └── path/              # Path pattern resolution
├── go.mod
├── README.md
├── Makefile
└── .github/workflows/
```

## Code Style

- Follow standard Go conventions (gofmt, go vet)
- Use clear, descriptive variable names
- Keep functions focused and small
- Add comments for non-obvious logic
- Use error wrapping with `fmt.Errorf` or `%w`

## Testing

While comprehensive testing isn't implemented yet, when adding features:

1. Test manually with different scenarios
2. Test on multiple platforms if possible
3. Test edge cases (spaces in names, special characters, etc.)

## Submitting Changes

1. Push to your fork:
   ```bash
   git push origin feature/your-feature
   ```

2. Create a Pull Request on GitHub with:
   - Clear title describing the change
   - Description of what was changed and why
   - Reference any related issues (e.g., "Fixes #123")

3. Respond to any feedback or review comments

## Reporting Bugs

Use GitHub Issues to report bugs. Include:
- How to reproduce the issue
- What you expected to happen
- What actually happened
- Your environment (OS, git version, etc.)

## Suggesting Features

Use GitHub Issues to suggest features. Include:
- Description of the feature
- Why it would be useful
- Example usage if applicable

## Release Process

Maintainers handle releases. Releases follow semantic versioning:
- Create a git tag: `git tag v1.2.3`
- Push the tag: `git push origin v1.2.3`
- GitHub Actions automatically builds and releases binaries

## Questions?

Feel free to open an issue or discussion if you have questions about contributing.

Thank you for helping make gwt better!
