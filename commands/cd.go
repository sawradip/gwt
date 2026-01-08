package commands

import (
	"fmt"
	"os"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
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

	// Get path pattern
	pattern, err := config.GetPathPattern()
	if err != nil {
		return err
	}

	// Resolve the path
	worktreePath, err := path.ResolvePatternForRepo(pattern, repoRoot, name)
	if err != nil {
		return err
	}

	// Check if worktree exists
	if !git.WorktreeExists(worktreePath) && !pathExists(worktreePath) {
		fmt.Fprintf(os.Stderr, "Worktree '%s' not found at %s\n\n", name, worktreePath)
		fmt.Fprintf(os.Stderr, "Available worktrees:\n")

		// List available worktrees
		if err := List(); err != nil {
			return err
		}

		return fmt.Errorf("worktree not found")
	}

	// Print the path to stdout
	fmt.Println(worktreePath)
	return nil
}

// pathExists checks if a path exists
func pathExists(p string) bool {
	_, err := os.Stat(p)
	return err == nil
}
