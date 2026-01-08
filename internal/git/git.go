package git

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

// GetRepoRoot returns the absolute path to the git repository root
func GetRepoRoot() (string, error) {
	cmd := exec.Command("git", "rev-parse", "--show-toplevel")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("not in a git repository")
	}
	return strings.TrimSpace(string(output)), nil
}

// GetRepoName returns the basename of the repository directory
func GetRepoName() (string, error) {
	root, err := GetRepoRoot()
	if err != nil {
		return "", err
	}
	return filepath.Base(root), nil
}

// GetCurrentBranch returns the currently checked out branch name
func GetCurrentBranch() (string, error) {
	cmd := exec.Command("git", "branch", "--show-current")
	output, err := cmd.Output()
	if err != nil {
		return "", fmt.Errorf("failed to get current branch: %w", err)
	}
	return strings.TrimSpace(string(output)), nil
}

// GetCurrentWorktree returns the path of the current worktree
func GetCurrentWorktree() (string, error) {
	cwd, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return cwd, nil
}

// BranchExists checks if a branch exists locally
func BranchExists(branch string) (bool, error) {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/heads/"+branch)
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// RemoteBranchExists checks if a branch exists on origin remote
func RemoteBranchExists(branch string) (bool, error) {
	cmd := exec.Command("git", "show-ref", "--verify", "--quiet", "refs/remotes/origin/"+branch)
	err := cmd.Run()
	if err != nil {
		if _, ok := err.(*exec.ExitError); ok {
			return false, nil
		}
		return false, err
	}
	return true, nil
}

// CreateWorktree creates a new worktree at the given path
func CreateWorktree(path string, branch string) error {
	// Check if we need to create the branch
	exists, err := BranchExists(branch)
	if err != nil {
		return err
	}

	remoteExists, err := RemoteBranchExists(branch)
	if err != nil {
		return err
	}

	var cmd *exec.Cmd
	if exists {
		// Branch exists locally, just add the worktree
		cmd = exec.Command("git", "worktree", "add", path, branch)
	} else if remoteExists {
		// Branch exists on remote, track it
		cmd = exec.Command("git", "worktree", "add", "--track", "-b", branch, path, "origin/"+branch)
	} else {
		// Create new branch
		cmd = exec.Command("git", "worktree", "add", "-b", branch, path)
	}

	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to create worktree: %s\n%w", string(output), err)
	}
	return nil
}

// RemoveWorktree removes a worktree at the given path
func RemoveWorktree(path string) error {
	cmd := exec.Command("git", "worktree", "remove", path)
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to remove worktree: %s\n%w", string(output), err)
	}
	return nil
}

// ListWorktrees returns a list of all worktrees (paths and branches)
func ListWorktrees() (map[string]string, error) {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return nil, fmt.Errorf("failed to list worktrees: %w", err)
	}

	worktrees := make(map[string]string)
	lines := strings.Split(string(output), "\n")

	var currentPath string
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			continue
		}

		// Parse "worktree <path>" line
		if strings.HasPrefix(line, "worktree ") {
			currentPath = strings.TrimPrefix(line, "worktree ")
			continue
		}

		// Parse "branch refs/heads/<branch>" line
		if strings.HasPrefix(line, "branch ") && currentPath != "" {
			branchRef := strings.TrimPrefix(line, "branch ")
			branchName := strings.TrimPrefix(branchRef, "refs/heads/")
			worktrees[currentPath] = branchName
			continue
		}

		// Handle detached HEAD
		if line == "detached" && currentPath != "" {
			worktrees[currentPath] = "(detached)"
			continue
		}
	}

	return worktrees, nil
}

// PruneWorktrees prunes stale worktree references
func PruneWorktrees() error {
	cmd := exec.Command("git", "worktree", "prune")
	output, err := cmd.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to prune worktrees: %s\n%w", string(output), err)
	}
	return nil
}

// WorktreeExists checks if a worktree exists at the given path
func WorktreeExists(path string) bool {
	cmd := exec.Command("git", "worktree", "list", "--porcelain")
	output, err := cmd.Output()
	if err != nil {
		return false
	}

	// Normalize the path we're looking for
	normalizedPath := filepath.Clean(path)

	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if strings.HasPrefix(line, "worktree ") {
			wtPath := filepath.Clean(strings.TrimPrefix(line, "worktree "))
			if wtPath == normalizedPath {
				return true
			}
		}
	}
	return false
}
