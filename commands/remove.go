package commands

import (
	"fmt"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
)

// Remove removes a worktree
func Remove(branch string) error {
	if branch == "" {
		return fmt.Errorf("branch name required")
	}

	if branch == "main" {
		return fmt.Errorf("cannot remove the main repository")
	}

	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// Get path pattern
	pattern, err := config.GetPathPattern()
	if err != nil {
		return err
	}

	// Resolve the worktree path
	worktreePath, err := path.ResolvePatternForRepo(pattern, repoRoot, branch)
	if err != nil {
		return err
	}

	// Check if worktree exists
	if !git.WorktreeExists(worktreePath) {
		return fmt.Errorf("worktree '%s' not found at %s", branch, worktreePath)
	}

	// Remove the worktree
	if err := git.RemoveWorktree(worktreePath); err != nil {
		return err
	}

	fmt.Printf("Removed worktree '%s' at %s\n", branch, worktreePath)
	return nil
}
