# Development Guide

Quick reference for building, testing, and releasing gwt.

## Building

```bash
# Build for current platform
go build -o gwt .

# Build and run
go build -o gwt . && ./gwt --help

# Build with version info
go build -ldflags "-X main.Version=v1.0.0" -o gwt .
```

## Testing Locally

```bash
# Test help
./gwt --help
./gwt create --help
./gwt ls --help

# Test in a git repo
cd /path/to/any/git/repo
./path/to/gwt ls
./path/to/gwt create test-branch
./path/to/gwt cd test-branch
./path/to/gwt rm test-branch
```

## Cross-Platform Builds

```bash
# Linux (x64)
GOOS=linux GOARCH=amd64 go build -o gwt-linux-amd64 .

# Linux (ARM)
GOOS=linux GOARCH=arm64 go build -o gwt-linux-arm64 .

# macOS (Intel)
GOOS=darwin GOARCH=amd64 go build -o gwt-darwin-amd64 .

# macOS (Apple Silicon)
GOOS=darwin GOARCH=arm64 go build -o gwt-darwin-arm64 .

# Windows
GOOS=windows GOARCH=amd64 go build -o gwt-windows-amd64.exe .
```

## Using Makefile

```bash
make build       # Build for current platform (output: ./dist/gwt)
make build-all   # Build for all platforms
make install     # Install to $GOPATH/bin
make clean       # Remove build artifacts
make test        # Run tests
make fmt         # Format code
make lint        # Run linter
```

## Git Tags & Releases

### Creating a New Release

```bash
# 1. Make sure all changes are committed
git status
git add .
git commit -m "Your commit message"

# 2. Create a version tag
git tag v1.0.0                    # Simple tag
git tag -a v1.0.0 -m "Release v1.0.0"  # Annotated tag (recommended)

# 3. Push commits and tags
git push origin main              # Push commits
git push origin v1.0.0            # Push specific tag
# OR
git push origin --tags            # Push all tags

# GitHub Actions will automatically build and create a release!
```

### Tag Naming Convention

Use semantic versioning: `vMAJOR.MINOR.PATCH`

- `v1.0.0` - First stable release
- `v1.1.0` - New features, backward compatible
- `v1.1.1` - Bug fixes
- `v2.0.0` - Breaking changes

### Managing Tags

```bash
# List all tags
git tag

# List tags with messages
git tag -n

# Delete a local tag
git tag -d v1.0.0

# Delete a remote tag
git push origin --delete v1.0.0

# Checkout a specific tag
git checkout v1.0.0
```

### Pre-release Tags

```bash
git tag v1.0.0-beta.1
git tag v1.0.0-rc.1
```

## Dependencies

```bash
# Download dependencies
go mod download

# Tidy up go.mod and go.sum
go mod tidy

# Update dependencies
go get -u ./...
```

## Common Issues

### "go: command not found"
Install Go from https://go.dev/dl/

### Build fails with import errors
```bash
go mod tidy
```

### Permission denied when running binary
```bash
chmod +x ./gwt
```

## Project Structure

```
gwt/
├── main.go              # Entry point, command routing
├── commands/            # Command implementations
│   ├── create.go
│   ├── list.go
│   ├── cd.go
│   ├── remove.go
│   ├── pwd.go
│   ├── config_cmd.go
│   ├── prune.go
│   └── help.go
├── internal/
│   ├── git/git.go       # Git command wrappers
│   ├── config/config.go # Config management
│   └── path/resolver.go # Path pattern resolution
├── go.mod
├── go.sum
├── Makefile
├── README.md
├── dev.md               # This file
└── .github/workflows/
    └── release.yml      # Auto-release on tags
```
