package commands

import (
	"fmt"
	"os"

	"github.com/yourusername/gwt/internal/git"
)

// Cd prints the path to the specified worktree
// The user needs to use a shell function to actually cd into it
func Cd(name string) error {
	if name == "" {
		return fmt.Errorf("worktree name required")
	}

	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// If name is "main", return the repo root
	if name == "main" {
		fmt.Println(repoRoot)
		return nil
	}

	// Get all worktrees from git
	worktrees, err := git.ListWorktrees()
	if err != nil {
		return err
	}

	// Find the worktree path for this branch
	var worktreePath string
	for path, branch := range worktrees {
		// Skip main repo
		if path == repoRoot {
			continue
		}
		// Found the branch
		if branch == name {
			worktreePath = path
			break
		}
	}

	if worktreePath == "" {
		fmt.Fprintf(os.Stderr, "Worktree '%s' not found\n\n", name)
		fmt.Fprintf(os.Stderr, "Available worktrees:\n")
		if err := List(); err != nil {
			return err
		}
		return fmt.Errorf("worktree not found")
	}

	// Print the path to stdout
	fmt.Println(worktreePath)
	return nil
}
