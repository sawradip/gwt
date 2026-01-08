package commands

import (
	"fmt"
	"path/filepath"
	"strings"

	"github.com/yourusername/gwt/internal/config"
	"github.com/yourusername/gwt/internal/git"
	"github.com/yourusername/gwt/internal/path"
)

// ConfigCmd handles config operations
func ConfigCmd(args []string) error {
	if len(args) == 0 {
		// Show current config
		return showConfig()
	}

	// Check for --path flag
	if len(args) >= 2 && args[0] == "--path" {
		pattern := args[1]
		return setConfig(pattern)
	}

	return fmt.Errorf("unknown config option: %s", args[0])
}

// showConfig displays the current path pattern
func showConfig() error {
	pattern, err := config.GetPathPattern()
	if err != nil {
		return err
	}

	repoRoot, err := git.GetRepoRoot()
	if err != nil {
		// Not in a repo, just show the pattern
		fmt.Printf("Path Pattern: %s\n", pattern)
		fmt.Println("\nPattern placeholders:")
		fmt.Println("  {repo}   - Repository name (from directory basename)")
		fmt.Println("  {branch} - Branch/worktree name")
		fmt.Println("\nExamples:")
		fmt.Println("  ../{repo}_wt/{branch}")
		fmt.Println("  ~/worktrees/{repo}/{branch}")
		fmt.Println("  /tmp/wt/{repo}/{branch}")
		return nil
	}

	repoName := filepath.Base(repoRoot)

	fmt.Printf("Path Pattern: %s\n", pattern)
	fmt.Println("\nPattern placeholders:")
	fmt.Println("  {repo}   - Repository name")
	fmt.Println("  {branch} - Branch/worktree name")

	// Show examples
	fmt.Printf("\nExamples for repository '%s':\n", repoName)
	exampleBranches := []string{"auth", "feat/new-ui", "bugfix"}

	for _, branch := range exampleBranches {
		resolved, err := path.ResolvePatternForRepo(pattern, repoRoot, branch)
		if err != nil {
			fmt.Printf("  %s: <error> %v\n", branch, err)
		} else {
			fmt.Printf("  %s -> %s\n", branch, resolved)
		}
	}

	return nil
}

// setConfig sets a new path pattern
func setConfig(pattern string) error {
	if !strings.Contains(pattern, "{branch}") {
		return fmt.Errorf("path pattern must contain {branch} placeholder")
	}

	if err := config.SetPathPattern(pattern); err != nil {
		return err
	}

	fmt.Printf("Path pattern updated: %s\n", pattern)

	// Show examples
	repoRoot, err := git.GetRepoRoot()
	if err == nil {
		repoName := filepath.Base(repoRoot)
		fmt.Printf("\nExamples for repository '%s':\n", repoName)

		exampleBranches := []string{"auth", "feat/new-ui", "bugfix"}
		for _, branch := range exampleBranches {
			resolved, err := path.ResolvePatternForRepo(pattern, repoRoot, branch)
			if err != nil {
				fmt.Printf("  %s: <error> %v\n", branch, err)
			} else {
				fmt.Printf("  %s -> %s\n", branch, resolved)
			}
		}
	}

	return nil
}
