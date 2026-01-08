package commands

import (
	"fmt"
	"os"
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

	// Get current working directory
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	fmt.Printf("Worktrees for %s:\n\n", repoName)

	// Check if we're in the main repo
	isInMain := isInMainRepo(cwd, repoRoot)

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
		if wtPath == repoRoot {
			continue
		}

		// Check if this is the current worktree
		currentPath := filepath.Clean(cwd)
		isCurrentWT := filepath.Clean(wtPath) == currentPath || isPathInWorktree(currentPath, wtPath)

		if isCurrentWT {
			color.Green("  * %s (%s)", branch, branch)
		} else {
			fmt.Printf("    %s (%s)\n", branch, branch)
		}
	}

	return nil
}

// isInMainRepo checks if the current directory is in the main repository
func isInMainRepo(cwd string, repoRoot string) bool {
	cleanCwd := filepath.Clean(cwd)
	cleanRepoRoot := filepath.Clean(repoRoot)

	// Check if cwd is within repoRoot
	return cleanCwd == cleanRepoRoot || (len(cleanCwd) > len(cleanRepoRoot) &&
		cleanCwd[:len(cleanRepoRoot)] == cleanRepoRoot &&
		(len(cleanCwd) == len(cleanRepoRoot) || cleanCwd[len(cleanRepoRoot)] == filepath.Separator))
}

// isPathInWorktree checks if a path is within a worktree directory
func isPathInWorktree(cwd string, wtPath string) bool {
	cleanCwd := filepath.Clean(cwd)
	cleanWtPath := filepath.Clean(wtPath)

	return cleanCwd == cleanWtPath || (len(cleanCwd) > len(cleanWtPath) &&
		cleanCwd[:len(cleanWtPath)] == cleanWtPath &&
		(len(cleanCwd) == len(cleanWtPath) || cleanCwd[len(cleanWtPath)] == filepath.Separator))
}
