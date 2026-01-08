package commands

import (
	"fmt"

	"github.com/yourusername/gwt/internal/git"
)

// Remove removes a worktree
func Remove(branch string) error {
	if branch == "" {
		return fmt.Errorf("branch name required")
	}

	if branch == "main" {
		return fmt.Errorf("cannot remove the main repository")
	}

	// Get all worktrees from git
	worktrees, err := git.ListWorktrees()
	if err != nil {
		return err
	}

	// Find the worktree path for this branch
	var worktreePath string
	var mainRepoPath string
	repoRoot, _ := git.GetRepoRoot()

	for path, b := range worktrees {
		// Track main repo path
		if path == repoRoot {
			mainRepoPath = path
			continue
		}

		// Found the branch we want to remove
		if b == branch {
			worktreePath = path
			break
		}
	}

	// Also check main repo path in case GetRepoRoot returned something different
	if mainRepoPath != "" && worktreePath == "" {
		// Branch not found in worktrees
		return fmt.Errorf("worktree '%s' not found", branch)
	}

	if worktreePath == "" {
		return fmt.Errorf("worktree '%s' not found\nAvailable worktrees:\n%s", branch, formatWorktreeList(worktrees, repoRoot))
	}

	// Remove the worktree using the actual path from git
	if err := git.RemoveWorktree(worktreePath); err != nil {
		return err
	}

	fmt.Printf("Removed worktree '%s' at %s\n", branch, worktreePath)
	return nil
}

func formatWorktreeList(worktrees map[string]string, repoRoot string) string {
	result := ""
	for path, branch := range worktrees {
		if path != repoRoot {
			result += fmt.Sprintf("  - %s (%s)\n", branch, path)
		}
	}
	return result
}
