package commands

import (
	"fmt"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
)

// Create creates a new worktree for the given branch
func Create(branch string) error {
	if branch == "" {
		return fmt.Errorf("branch name required")
	}

	if branch == "main" {
		return fmt.Errorf("cannot create worktree named 'main' (reserved for main repository)")
	}

	// Get repo root
	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// Get path pattern from config
	pattern, err := config.GetPathPattern()
	if err != nil {
		return err
	}

	// Resolve the worktree path
	worktreePath, err := path.ResolvePatternForRepo(pattern, repoRoot, branch)
	if err != nil {
		return err
	}

	// Check if worktree already exists
	if git.WorktreeExists(worktreePath) {
		return fmt.Errorf("worktree already exists at %s\nUse 'gwt cd %s' to switch to it", worktreePath, branch)
	}

	// Create the worktree
	if err := git.CreateWorktree(worktreePath, branch); err != nil {
		return err
	}

	// Print success message with the path
	fmt.Printf("Created worktree for '%s' at:\n%s\n", branch, worktreePath)
	fmt.Println("\nTo navigate to this worktree, run:")
	fmt.Printf("  gwt cd %s\n", branch)

	return nil
}
