package commands

import (
	"fmt"
)

// Help displays general help information
func Help() {
	fmt.Print(`gwt - Git Worktree Manager

USAGE
    gwt <command> [options]

COMMANDS
    create <branch>   Create a new worktree
    ls, list          List all worktrees
    cd <name>         Switch to a worktree (use with shell function)
    rm, remove <name> Remove a worktree
    pwd               Show current worktree name
    prune             Clean up stale worktree references
    config            View or set path pattern
    version           Show version information
    help              Show this help message

OPTIONS
    -h, --help        Show help (works with any command)
    -v, --version     Show version information

EXAMPLES
    gwt create feat/auth      Create worktree for feat/auth branch
    gwt ls                    List all worktrees
    gwt cd feat/auth          Switch to worktree
    gwt rm feat/auth          Remove worktree
    gwt config --path ~/wt/{repo}/{branch}
    gwt --version             Show version

SHELL INTEGRATION
    Add to ~/.bashrc or ~/.zshrc for 'gwt cd' to work:

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

Run 'gwt <command> --help' for detailed help on a specific command.
`)
}

// HelpCreate displays help for the create command
func HelpCreate() {
	fmt.Print(`gwt create - Create a new worktree

USAGE
    gwt create <branch>

DESCRIPTION
    Creates a new git worktree for the specified branch at a path
    determined by the configured path pattern.

    Branch behavior:
    - If branch exists locally: checks it out in the new worktree
    - If branch exists on origin: creates local branch tracking origin
    - If branch doesn't exist: creates a new branch

ARGUMENTS
    <branch>    Branch name (e.g., feat/auth, bugfix/issue-42)

RESERVED NAMES
    The following names cannot be used as branch names:
    main, master

EXAMPLES
    gwt create feat/auth
    gwt create bugfix/issue-42

SEE ALSO
    gwt config --help         Configure worktree path pattern
    gwt ls --help             List existing worktrees
`)
}

// HelpLs displays help for the ls command
func HelpLs() {
	fmt.Print(`gwt ls - List all worktrees

USAGE
    gwt ls
    gwt list

DESCRIPTION
    Lists all git worktrees for the current repository.
    The current worktree is marked with '*' and shown in green.

OUTPUT FORMAT
    Worktrees for <repo-name>:

      * main (main repository)      <- current worktree
        feat/auth (feat/auth)
        bugfix (bugfix)

EXAMPLES
    gwt ls
    gwt list
`)
}

// HelpCd displays help for the cd command
func HelpCd() {
	fmt.Print(`gwt cd - Switch to a worktree

USAGE
    gwt cd <name>
    gwt cd main

DESCRIPTION
    Prints the path to the specified worktree. When used with the
    shell function wrapper, this changes the current directory.

    NOTE: A subprocess cannot change the parent shell's directory.
    You must add the shell function to your shell config for this
    command to actually change directories.

ARGUMENTS
    <name>      Worktree name (branch name used when creating)
    main        Special name for the main repository

SHELL INTEGRATION
    Add to ~/.bashrc or ~/.zshrc:

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

EXAMPLES
    gwt cd feat/auth          Switch to feat/auth worktree
    gwt cd main               Switch to main repository

SEE ALSO
    gwt ls --help             List available worktrees
`)
}

// HelpRm displays help for the rm command
func HelpRm() {
	fmt.Print(`gwt rm - Remove a worktree

USAGE
    gwt rm <name>
    gwt remove <name>

DESCRIPTION
    Removes a git worktree and cleans up its git references.
    The worktree directory will be deleted.

    WARNING: Make sure you have committed or stashed any changes
    in the worktree before removing it.

ARGUMENTS
    <name>      Worktree name to remove

RESTRICTIONS
    - Cannot remove 'main' (the main repository)
    - Worktree must exist

EXAMPLES
    gwt rm feat/auth
    gwt remove bugfix/issue-42

SEE ALSO
    gwt ls --help             List existing worktrees
    gwt prune --help          Clean up stale references
`)
}

// HelpPwd displays help for the pwd command
func HelpPwd() {
	fmt.Print(`gwt pwd - Show current worktree name

USAGE
    gwt pwd

DESCRIPTION
    Prints the name of the current worktree based on your
    current working directory.

    Returns "main" if you're in the main repository or if
    the current directory doesn't match any known worktree.

OUTPUT
    Prints the worktree name to stdout.

EXAMPLES
    $ cd ~/projects/myapp
    $ gwt pwd
    main

    $ gwt cd feat/auth
    $ gwt pwd
    feat/auth
`)
}

// HelpPrune displays help for the prune command
func HelpPrune() {
	fmt.Print(`gwt prune - Clean up stale worktree references

USAGE
    gwt prune

DESCRIPTION
    Removes stale worktree references from git. This is useful when
    worktree directories have been manually deleted without using
    'gwt rm'.

    Equivalent to running 'git worktree prune'.

WHEN TO USE
    - After manually deleting worktree directories
    - When 'gwt ls' shows worktrees that no longer exist
    - As periodic maintenance

EXAMPLES
    gwt prune
`)
}

// HelpConfig displays help for the config command
func HelpConfig() {
	fmt.Print(`gwt config - View or set path pattern

USAGE
    gwt config
    gwt config --path <pattern>

DESCRIPTION
    Manages the path pattern that determines where worktrees are created.
    Configuration is stored in git global config (gwt.pathpattern).

OPTIONS
    --path <pattern>    Set a new path pattern

PATTERN PLACEHOLDERS
    {repo}      Repository name (directory basename)
    {branch}    Branch name provided when creating worktree

PATTERN TYPES
    Relative to repo parent:    ../{repo}_wt/{branch}
    Home directory:             ~/worktrees/{repo}/{branch}
    Absolute path:              /tmp/wt/{repo}/{branch}

DEFAULT PATTERN
    ../{repo}_wt/{branch}

EXAMPLES
    gwt config
        Show current pattern with example paths

    gwt config --path ~/worktrees/{repo}/{branch}
        Set pattern to use home directory

    gwt config --path /tmp/wt/{repo}/{branch}
        Set pattern to use temp directory

RESET TO DEFAULT
    git config --global --unset gwt.pathpattern

SEE ALSO
    gwt create --help         Create worktrees using the pattern
`)
}
