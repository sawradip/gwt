package commands

import (
	"fmt"
)

// Help displays help information
func Help() {
	fmt.Println(`gwt - Git Worktree Manager

A simple CLI tool for managing git worktrees with a filesystem-like interface.

USAGE:
    gwt <command> [options]

COMMANDS:
    <branch-name>    Create a new worktree for the given branch
                     Example: gwt feat/auth

    ls, list         List all worktrees
                     Example: gwt ls

    cd <name>        Print path to a worktree (use with shell function)
                     Example: gwt cd feat/auth

    rm, remove <name> Remove a worktree
                     Example: gwt rm feat/auth

    pwd              Show current worktree name
                     Example: gwt pwd

    prune            Clean up stale worktree references
                     Example: gwt prune

    config           View or set path pattern configuration
                     Examples:
                       gwt config
                       gwt config --path ~/worktrees/{repo}/{branch}

    help, -h, --help Show this help message

SHELL INTEGRATION:

Add this function to your shell configuration (.bashrc, .zshrc, etc.):

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

For Fish shell (~/.config/fish/config.fish):

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

PATH PATTERNS:

Configure where worktrees are created. Patterns support:
  {repo}   - Repository name (directory basename)
  {branch} - Branch/worktree name

Examples:
  ../repo_wt/{branch}         - Relative to repo parent
  ~/worktrees/{repo}/{branch} - Under home directory
  /tmp/wt/{repo}/{branch}     - Absolute path

EXAMPLES:

    # Configure custom path pattern
    gwt config --path ~/worktrees/{repo}/{branch}

    # Create a new worktree
    gwt feat/auth

    # List all worktrees
    gwt ls

    # Switch to a worktree (requires shell function)
    gwt cd feat/auth

    # See where you are
    gwt pwd

    # Remove a worktree
    gwt rm feat/auth

    # Cleanup stale references
    gwt prune

For more information, visit: https://github.com/yourusername/gwt
`)
}
