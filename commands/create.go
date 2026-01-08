package commands

import (
	"fmt"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
)

// Reserved names that cannot be used as worktree names
var reservedNames = map[string]bool{
	"main":   true,
	"master": true,
}

// Create creates a new worktree for the given branch
func Create(branch string) error {
	if branch == "" {
		return fmt.Errorf("branch name required")
	}

	if reservedNames[branch] {
		return fmt.Errorf("'%s' is a reserved name and cannot be used as a worktree name", branch)
	}

	// Get repo root
	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// Check if worktree with this branch name already exists
	worktrees, err := git.ListWorktrees()
	if err != nil {
		return err
	}

	for path, b := range worktrees {
		if b == branch && path != repoRoot {
			return fmt.Errorf("worktree '%s' already exists at %s\nUse 'gwt cd %s' to switch to it", branch, path, branch)
		}
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
