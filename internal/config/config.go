package config

import (
	"fmt"
	"os/exec"
	"strings"
)

const (
	configKey = "gwt.pathpattern"
	defaultPattern = "../{repo}_wt/{branch}"
)

// GetPathPattern retrieves the path pattern from git global config
// Returns default pattern if not configured
func GetPathPattern() (string, error) {
	cmd := exec.Command("git", "config", "--global", "--get", configKey)
	output, err := cmd.Output()

	if err != nil {
		// If not found, return default
		if _, ok := err.(*exec.ExitError); ok {
			return defaultPattern, nil
		}
		return "", fmt.Errorf("failed to read config: %w", err)
	}

	pattern := strings.TrimSpace(string(output))
	if pattern == "" {
		return defaultPattern, nil
	}

	return pattern, nil
}

// SetPathPattern sets the path pattern in git global config
func SetPathPattern(pattern string) error {
	if !strings.Contains(pattern, "{branch}") {
		return fmt.Errorf("path pattern must contain {branch} placeholder")
	}

	cmd := exec.Command("git", "config", "--global", configKey, pattern)
	output, err := cmd.CombinedOutput()

	if err != nil {
		return fmt.Errorf("failed to set config: %s\n%w", string(output), err)
	}

	return nil
}

// GetDefaultPattern returns the default path pattern
func GetDefaultPattern() string {
	return defaultPattern
}
