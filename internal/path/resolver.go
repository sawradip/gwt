package path

import (
	"fmt"
	"os"
	"os/user"
	"path/filepath"
	"strings"
)

const defaultPattern = "../{repo}_wt/{branch}"

// ResolvePath resolves a path pattern with variable substitution
func ResolvePath(pattern string, repo string, branch string) (string, error) {
	if !strings.Contains(pattern, "{branch}") {
		return "", fmt.Errorf("path pattern must contain {branch} placeholder")
	}

	// Replace variables
	path := strings.ReplaceAll(pattern, "{repo}", repo)
	path = strings.ReplaceAll(path, "{branch}", branch)

	// Expand home directory
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		path = filepath.Join(homeDir, path[2:])
	}

	// Handle relative paths
	if !filepath.IsAbs(path) {
		// If path starts with ../, it's relative to the repo root's parent
		if strings.HasPrefix(path, ".."+string(filepath.Separator)) || strings.HasPrefix(path, "../") {
			// We need the repo root to resolve relative paths
			cwd, err := os.Getwd()
			if err != nil {
				return "", err
			}
			// Get the parent of the repo root
			repoRoot := filepath.Dir(cwd)
			path = filepath.Join(repoRoot, path)
		} else {
			// Relative to current directory
			cwd, err := os.Getwd()
			if err != nil {
				return "", err
			}
			path = filepath.Join(cwd, path)
		}
	}

	// Normalize the path
	path = filepath.Clean(path)

	return path, nil
}

// ResolvePatternForRepo resolves a path for a given repo and branch
// This function handles the case where we need to know the actual repo root
func ResolvePatternForRepo(pattern string, repoRoot string, branch string) (string, error) {
	if !strings.Contains(pattern, "{branch}") {
		return "", fmt.Errorf("path pattern must contain {branch} placeholder")
	}

	// Get repo name from repo root path
	repoName := filepath.Base(repoRoot)

	// Replace variables
	path := strings.ReplaceAll(pattern, "{repo}", repoName)
	path = strings.ReplaceAll(path, "{branch}", branch)

	// Expand home directory
	if strings.HasPrefix(path, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", fmt.Errorf("failed to get home directory: %w", err)
		}
		path = filepath.Join(homeDir, path[2:])
	}

	// Handle relative paths
	if !filepath.IsAbs(path) {
		if strings.HasPrefix(path, ".."+string(filepath.Separator)) || strings.HasPrefix(path, "../") {
			// Relative to the parent of repo root
			// First normalize repoRoot to get its true parent
			repoParent := filepath.Dir(filepath.Clean(repoRoot))
			// Remove the leading ../ from the pattern before joining
			relativePath := strings.TrimPrefix(path, "../")
			relativePath = strings.TrimPrefix(relativePath, ".."+string(filepath.Separator))
			// Join the relative pattern from the repo's parent
			path = filepath.Join(repoParent, relativePath)
		} else {
			// Relative to repo root
			path = filepath.Join(repoRoot, path)
		}
	}

	// Normalize the path
	path = filepath.Clean(path)

	return path, nil
}

// GetDefaultPattern returns the default path pattern
func GetDefaultPattern() string {
	return defaultPattern
}

// ExtractWorktreeNameFromPath extracts the branch/worktree name from a worktree path
// given the pattern and repo root
func ExtractWorktreeNameFromPath(path string, pattern string, repoRoot string) (string, error) {
	repoName := filepath.Base(repoRoot)

	// Build expected path structure by replacing {branch} with a placeholder
	placeholder := "__BRANCH_PLACEHOLDER__"
	expectedPattern := strings.ReplaceAll(pattern, "{repo}", repoName)
	expectedPattern = strings.ReplaceAll(expectedPattern, "{branch}", placeholder)

	// Expand home directory in pattern
	if strings.HasPrefix(expectedPattern, "~/") {
		homeDir, err := os.UserHomeDir()
		if err != nil {
			return "", err
		}
		expectedPattern = filepath.Join(homeDir, expectedPattern[2:])
	}

	// Normalize paths for comparison
	path = filepath.Clean(path)
	expectedPattern = filepath.Clean(expectedPattern)

	// Find the placeholder in the normalized pattern
	if !strings.Contains(expectedPattern, placeholder) {
		return "", fmt.Errorf("could not extract worktree name from path")
	}

	// Simple extraction: split by placeholder and extract the middle part
	parts := strings.Split(expectedPattern, placeholder)
	if len(parts) != 2 {
		return "", fmt.Errorf("invalid pattern structure")
	}

	prefix := parts[0]
	suffix := parts[1]

	if !strings.HasPrefix(path, prefix) {
		return "", fmt.Errorf("path does not match pattern")
	}

	withoutPrefix := strings.TrimPrefix(path, prefix)

	if suffix != "" && strings.HasSuffix(withoutPrefix, suffix) {
		withoutPrefix = strings.TrimSuffix(withoutPrefix, suffix)
	}

	// Clean up path separators
	branchName := strings.Trim(withoutPrefix, string(filepath.Separator))

	return branchName, nil
}

// NormalizePath normalizes a path for comparison
func NormalizePath(p string) string {
	// Expand home directory
	if strings.HasPrefix(p, "~/") {
		homeDir, _ := os.UserHomeDir()
		p = filepath.Join(homeDir, p[2:])
	}

	// Resolve to absolute path if needed
	if !filepath.IsAbs(p) {
		if cwd, err := os.Getwd(); err == nil {
			p = filepath.Join(cwd, p)
		}
	}

	return filepath.Clean(p)
}

// ExpandHome expands the home directory in a path
func ExpandHome(p string) (string, error) {
	if strings.HasPrefix(p, "~") {
		usr, err := user.Current()
		if err != nil {
			return "", err
		}
		return strings.Replace(p, "~", usr.HomeDir, 1), nil
	}
	return p, nil
}
