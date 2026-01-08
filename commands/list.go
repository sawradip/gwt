package commands

import (
	"fmt"
	"path/filepath"

	"github.com/fatih/color"
	"github.com/yourusername/gwt/internal/git"
)

// List lists all worktrees
func List() error {
	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	repoName := filepath.Base(repoRoot)

	// Get current worktree path (where we actually are)
	currentWtPath, err := git.GetCurrentWorktreePath()
	if err != nil {
		return err
	}

	fmt.Printf("Worktrees for %s:\n\n", repoName)

	// Check if we're in the main repo
	isInMain := filepath.Clean(currentWtPath) == filepath.Clean(repoRoot)

	// Print main repo
	if isInMain {
		color.Green("  * main (main repository)")
	} else {
		fmt.Println("    main (main repository)")
	}

	// Get all worktrees
	worktrees, err := git.ListWorktrees()
	if err != nil {
		return err
	}

	// Filter and display worktrees (excluding main)
	for wtPath, branch := range worktrees {
		// Skip the main repository entry from git worktree list
		if filepath.Clean(wtPath) == filepath.Clean(repoRoot) {
			continue
		}

		// Check if this is the current worktree
		isCurrentWT := filepath.Clean(wtPath) == filepath.Clean(currentWtPath)

		if isCurrentWT {
			color.Green("  * %s (%s)", branch, branch)
		} else {
			fmt.Printf("    %s (%s)\n", branch, branch)
		}
	}

	return nil
}
