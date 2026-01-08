package commands

import (
	"fmt"

	"github.com/yourusername/gwt/internal/git"
)

// Prune prunes stale worktree references
func Prune() error {
	if err := git.PruneWorktrees(); err != nil {
		return err
	}

	fmt.Println("Pruned stale worktree references")
	return nil
}
