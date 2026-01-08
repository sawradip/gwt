package commands

import (
	"fmt"
	"path/filepath"

	"github.com/yourusername/gwt/internal/git"
)

// Pwd prints the current worktree name
func Pwd() error {
	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// Get current worktree path
	currentWtPath, err := git.GetCurrentWorktreePath()
	if err != nil {
		return err
	}

	// Check if we're in the main repo
	if filepath.Clean(currentWtPath) == filepath.Clean(repoRoot) {
		fmt.Println("main")
		return nil
	}

	// Get worktree list to find our branch name
	worktrees, err := git.ListWorktrees()
	if err != nil {
		return err
	}

	// Find the branch name for our current worktree
	for wtPath, branch := range worktrees {
		if filepath.Clean(wtPath) == filepath.Clean(currentWtPath) {
			fmt.Println(branch)
			return nil
		}
	}

	// Fallback to main if we can't find it
	fmt.Println("main")
	return nil
}
