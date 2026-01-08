package commands

import (
	"fmt"
	"os"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
)

// Pwd prints the current worktree name
func Pwd() error {
	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		return err
	}

	// Check if we're in the main repo
	if isInMainRepo(cwd, repoRoot) {
		fmt.Println("main")
		return nil
	}

	// Get path pattern
	pattern, err := config.GetPathPattern()
	if err != nil {
		return err
	}

	// Try to extract branch name from current path
	branchName, err := path.ExtractWorktreeNameFromPath(cwd, pattern, repoRoot)
	if err != nil {
		// If we can't extract it, we're not in a known worktree
		fmt.Println("main")
		return nil
	}

	fmt.Println(branchName)
	return nil
}
