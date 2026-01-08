# gwt - Git Worktree Manager

A CLI tool for managing git worktrees with a filesystem-like interface.

## Quick Install

### macOS

```bash
# Apple Silicon (M1/M2/M3)
curl -L https://github.com/yourusername/gwt/releases/latest/download/gwt-darwin-arm64 -o gwt
chmod +x gwt && sudo mv gwt /usr/local/bin/

# Intel Mac
curl -L https://github.com/yourusername/gwt/releases/latest/download/gwt-darwin-amd64 -o gwt
chmod +x gwt && sudo mv gwt /usr/local/bin/
```

### Linux

```bash
# x64
curl -L https://github.com/yourusername/gwt/releases/latest/download/gwt-linux-amd64 -o gwt
chmod +x gwt && sudo mv gwt /usr/local/bin/

# ARM64
curl -L https://github.com/yourusername/gwt/releases/latest/download/gwt-linux-arm64 -o gwt
chmod +x gwt && sudo mv gwt /usr/local/bin/
```

### Windows (PowerShell as Admin)

```powershell
# Download to a folder in your PATH
Invoke-WebRequest -Uri "https://github.com/yourusername/gwt/releases/latest/download/gwt-windows-amd64.exe" -OutFile "$env:LOCALAPPDATA\Microsoft\WindowsApps\gwt.exe"
```

### Build from Source

```bash
git clone https://github.com/yourusername/gwt.git
cd gwt
go build -o gwt . && sudo mv gwt /usr/local/bin/
```

## Shell Integration (Required for `gwt cd`)

Add this to your shell config, then restart your terminal:

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
cd ~/projects/myapp

gwt create feat/auth     # Create a worktree
gwt ls                   # List all worktrees
gwt cd feat/auth         # Switch to worktree
gwt pwd                  # Show current worktree
gwt cd main              # Go back to main repo
gwt rm feat/auth         # Remove worktree
```

## Commands

| Command | Description |
|---------|-------------|
| `gwt create <branch>` | Create worktree for branch |
| `gwt ls` | List all worktrees |
| `gwt cd <name>` | Switch to worktree (or `main`) |
| `gwt pwd` | Show current worktree name |
| `gwt rm <name>` | Remove a worktree |
| `gwt prune` | Clean stale references |
| `gwt config` | View path pattern |
| `gwt config --path <pattern>` | Set path pattern |

All commands support `--help` for detailed usage.

## Path Patterns

Control where worktrees are created. Default: `../{repo}_wt/{branch}`

```bash
# View current pattern
gwt config

# Set custom pattern
gwt config --path ~/worktrees/{repo}/{branch}
```

**Placeholders:**
- `{repo}` - Repository directory name
- `{branch}` - Branch name

**Examples:**

| Pattern | Result for `feat/auth` in `myapp` repo |
|---------|----------------------------------------|
| `../{repo}_wt/{branch}` | `../myapp_wt/feat/auth` |
| `~/worktrees/{repo}/{branch}` | `~/worktrees/myapp/feat/auth` |

## Troubleshooting

**`gwt cd` doesn't change directory** - Add the shell function above and restart your terminal.

**"Not in a git repository"** - Run gwt from inside a git repo.

**Reset config** - `git config --global --unset gwt.pathpattern`

## License

MIT
