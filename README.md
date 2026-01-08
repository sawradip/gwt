# gwt - Git Worktree Manager

A CLI tool for managing git worktrees with a filesystem-like interface.

## Installation

### Download Binary

Download from [releases](https://github.com/yourusername/gwt/releases):

| Platform | Binary |
|----------|--------|
| Linux (x64) | `gwt-linux-amd64` |
| Linux (ARM) | `gwt-linux-arm64` |
| macOS (Intel) | `gwt-darwin-amd64` |
| macOS (Apple Silicon) | `gwt-darwin-arm64` |
| Windows | `gwt-windows-amd64.exe` |

```bash
# Linux/macOS
chmod +x gwt
sudo mv gwt /usr/local/bin/
```

### Build from Source

```bash
git clone https://github.com/yourusername/gwt.git
cd gwt
make build && make install
```

### Shell Integration (Required for `gwt cd`)

Since a subprocess cannot change the parent shell's directory, add this wrapper function:

**Bash/Zsh** (`~/.bashrc` or `~/.zshrc`):
```bash
gwt() {
    if [ "$1" = "cd" ]; then
        local target_path=$(command gwt cd "$2")
        if [ $? -eq 0 ] && [ -n "$target_path" ]; then
            cd "$target_path"
        fi
    else
        command gwt "$@"
    fi
}
```

**Fish** (`~/.config/fish/config.fish`):
```fish
function gwt
    if test "$argv[1]" = "cd"
        set target_path (command gwt cd $argv[2])
        if test $status -eq 0 -a -n "$target_path"
            cd "$target_path"
        end
    else
        command gwt $argv
    end
end
```

**PowerShell** (`$PROFILE`):
```powershell
function gwt {
    if ($args[0] -eq "cd") {
        $targetPath = & gwt.exe cd $args[1] 2>$null
        if ($LASTEXITCODE -eq 0 -and $targetPath) {
            Set-Location $targetPath
        }
    } else {
        & gwt.exe @args
    }
}
```

## Quick Start

```bash
# Configure where worktrees are created (optional)
gwt config --path ~/worktrees/{repo}/{branch}

# Create a worktree
cd ~/projects/myapp
gwt feat/auth

# List all worktrees
gwt ls

# Switch to a worktree
gwt cd feat/auth

# Check which worktree you're in
gwt pwd

# Go back to main repository
gwt cd main

# Remove a worktree
gwt rm feat/auth
```

## Commands

| Command | Description |
|---------|-------------|
| `gwt <branch>` | Create worktree for branch |
| `gwt ls` | List all worktrees |
| `gwt cd <name>` | Switch to worktree (or `main`) |
| `gwt pwd` | Show current worktree name |
| `gwt rm <name>` | Remove a worktree |
| `gwt prune` | Clean stale references |
| `gwt config` | View path pattern |
| `gwt config --path <pattern>` | Set path pattern |

### `gwt <branch>`

Creates a worktree. If the branch exists (locally or on origin), it checks it out. Otherwise, creates a new branch.

```bash
gwt feat/auth        # Creates worktree and checks out/creates branch
gwt bugfix/issue-42  # Branch names with / are supported
```

Blocked: `gwt main` (reserved for main repository)

### `gwt ls`

Lists worktrees with `*` marking the current one:

```
Worktrees for myapp:

  * main (main repository)
    feat/auth (feat/auth)
    bugfix (bugfix)
```

### `gwt cd <name>`

Switches to a worktree. Use `main` to return to the main repository.

```bash
gwt cd feat/auth
gwt cd main
```

### `gwt rm <name>`

Removes a worktree and cleans up git references.

```bash
gwt rm feat/auth
```

## Path Patterns

Patterns control where worktrees are created. Default: `../{repo}_wt/{branch}`

**Placeholders:**
- `{repo}` - Repository directory name
- `{branch}` - Branch name you provide

**Examples:**

| Pattern | Repo | Branch | Result |
|---------|------|--------|--------|
| `../{repo}_wt/{branch}` | `~/projects/myapp` | `feat/auth` | `~/projects/myapp_wt/feat/auth` |
| `~/worktrees/{repo}/{branch}` | `~/projects/myapp` | `feat/auth` | `~/worktrees/myapp/feat/auth` |
| `/tmp/wt/{repo}/{branch}` | `~/projects/myapp` | `bugfix` | `/tmp/wt/myapp/bugfix` |

```bash
# View current pattern with examples
gwt config

# Set a new pattern
gwt config --path ~/worktrees/{repo}/{branch}
```

Configuration is stored in git global config (`gwt.pathpattern`).

## Troubleshooting

**"Not in a git repository"**
Run gwt from inside a git repository.

**"Worktree already exists"**
Use `gwt cd <name>` to switch to it, or `gwt rm <name>` to remove it first.

**`gwt cd` doesn't change directory**
Add the shell function from [Shell Integration](#shell-integration-required-for-gwt-cd) and reload your shell config.

**Reset path pattern to default**
```bash
git config --global --unset gwt.pathpattern
```

## Development

```bash
make build      # Build for current platform
make build-all  # Build for all platforms
make install    # Install to $GOPATH/bin
make test       # Run tests
make clean      # Remove build artifacts
```

## License

MIT
