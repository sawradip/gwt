# gwt Usage Workflows

This document shows real-world workflows and usage patterns for gwt.

## Table of Contents

1. [Basic Setup](#basic-setup)
2. [Feature Branch Development](#feature-branch-development)
3. [Multiple Repositories](#multiple-repositories)
4. [Team Workflows](#team-workflows)
5. [Advanced Patterns](#advanced-patterns)

## Basic Setup

### Initial Configuration

```bash
# View current configuration (uses default if not set)
gwt config

# Set a custom path pattern
gwt config --path ~/dev/wt/{repo}/{branch}

# Or use system temp directory
gwt config --path /tmp/wt/{repo}/{branch}
```

## Feature Branch Development

### Single Feature

```bash
# Go to your repository
cd ~/projects/my-app

# Create a new feature worktree
gwt feat/user-auth

# This creates the worktree and prints the path
# Output: Created worktree for 'feat/user-auth' at:
#         /Users/you/dev/wt/my-app/feat/user-auth

# Switch to the worktree
gwt cd feat/user-auth

# Do your development work...
# Make commits, test, etc.

# Check where you are
gwt pwd
# Output: feat/user-auth

# List all worktrees to see what you have
gwt ls

# When done, switch back
gwt cd main

# Remove the worktree
gwt rm feat/user-auth
```

### Parallel Feature Development

```bash
# Create multiple feature branches to work on in parallel
cd ~/projects/my-app

gwt feat/authentication
gwt feat/payment-integration
gwt feat/dashboard-redesign

# List all your work
gwt ls
# Output:
# Worktrees for my-app:
#
#   * main (main repository)
#     feat/authentication (feat/authentication)
#     feat/payment-integration (feat/payment-integration)
#     feat/dashboard-redesign (feat/dashboard-redesign)

# Quickly switch between them
gwt cd feat/authentication    # Work on auth
gwt cd feat/payment-integration  # Switch to payments
gwt cd feat/dashboard-redesign   # Work on UI

# Back to main when needed
gwt cd main
```

### Bug Fixes and Hotfixes

```bash
# Create worktrees for different bug fixes
gwt fix/critical-security-issue
gwt fix/customer-reported-bug
gwt fix/edge-case-handling

# Work on each in isolation
gwt cd fix/critical-security-issue

# Complete one, remove it
gwt rm fix/critical-security-issue

# Move to the next
gwt cd fix/customer-reported-bug
```

## Multiple Repositories

### Organizing Multiple Projects

```bash
# Configure a global pattern for all your projects
gwt config --path ~/dev/worktrees/{repo}/{branch}

# Now you can have organized worktrees across all repos
cd ~/projects/backend
gwt api/v2-endpoints
gwt db/migration-auth

cd ~/projects/frontend
gwt ui/new-design
gwt perf/optimization

cd ~/projects/mobile
gwt ios/navigation
gwt android/onboarding

# All organized under ~/dev/worktrees/
```

### Repository-Specific Configuration

```bash
# Check current config
gwt config

# Each repo can have its own pattern by setting it in that repo's git config
cd ~/projects/special-project

# This sets it in the global config
gwt config --path ./local-worktrees/{branch}

# Other repos will use their own patterns
```

## Team Workflows

### Code Review Workflow

```bash
# Pull latest and see what branches need review
git fetch
git branch -r

# Create worktrees for each PR you're reviewing
gwt review/feature-123
gwt review/bugfix-456

# Review each one
gwt cd review/feature-123
# ... review code, run tests ...

# Check the next one
gwt cd review/bugfix-456
# ... review code, run tests ...

# Clean up when done
gwt rm review/feature-123
gwt rm review/bugfix-456
```

### Release Branch Management

```bash
# Create release worktrees
gwt release/v1.2.0
gwt release/v1.3.0-rc1

# Each can be tested independently
gwt cd release/v1.2.0
# ... run release tests ...

# List all active releases
gwt ls
```

## Advanced Patterns

### Ephemeral Testing Branches

```bash
# Quickly create test branches
gwt test/integration-tests
gwt test/performance-benchmarks
gwt test/database-migration

# Test something quickly without affecting main branch
gwt cd test/integration-tests
# ... run tests ...
# ... if successful, merge to main ...

# If tests fail, just remove and try again
gwt rm test/integration-tests
gwt test/integration-tests
```

### Cross-Repository Development

```bash
# When working on related changes across multiple repos
gwt config --path ~/dev/feature-xyz/{repo}

cd ~/projects/api
gwt xyz-implementation

cd ~/projects/web
gwt xyz-frontend

cd ~/projects/mobile
gwt xyz-mobile

# All feature-related worktrees are grouped together
ls ~/dev/feature-xyz/
# Output:
# api/
# web/
# mobile/
```

### Maintaining Multiple Versions

```bash
# Keep worktrees for different versions
gwt v1.0/maintenance
gwt v1.1/maintenance
gwt v2.0/dev

# List all version-specific work
gwt ls

# Switch between versions
gwt cd v1.0/maintenance    # Backport a fix
gwt cd v1.1/maintenance    # Backport a fix
gwt cd v2.0/dev            # Continue development
```

### Clean Up and Maintenance

```bash
# List all worktrees
gwt ls

# Remove old/completed worktrees
gwt rm feature-old
gwt rm hotfix/resolved

# Clean up stale references (e.g., after manual deletions)
gwt prune

# View what's left
gwt ls
```

## Performance Tips

### Use Path Patterns Wisely

```bash
# Good: Fast to navigate, organized
gwt config --path ~/dev/wt/{repo}/{branch}

# Also good: Using relative paths (relative to repo parent)
gwt config --path ../wt/{branch}

# Slower: Absolute paths on network drives
gwt config --path /mnt/network/wt/{repo}/{branch}
```

### Batch Operations

```bash
# Clean up multiple worktrees at once
gwt rm fix/1
gwt rm fix/2
gwt rm fix/3
gwt rm fix/4

# Or use a script
for branch in fix/1 fix/2 fix/3 fix/4; do
    gwt rm "$branch"
done
```

## Troubleshooting

### Worktree Creation Failed

```bash
# Check if worktree already exists
gwt ls

# If it shows in the list but you can't create it again
gwt rm old-branch
gwt old-branch
```

### Can't Switch Directories

```bash
# Make sure you have the shell function installed
# For bash/zsh, add to ~/.bashrc or ~/.zshrc:
# gwt() {
#     if [ "$1" = "cd" ]; then
#         local target_path=$(command gwt cd "$2")
#         if [ $? -eq 0 ] && [ -n "$target_path" ]; then
#             cd "$target_path"
#         fi
#     else
#         command gwt "$@"
#     fi
# }

# Then reload your shell
source ~/.bashrc  # or source ~/.zshrc
```

### Stale Worktree References

```bash
# If you manually deleted worktree directories, clean up references
gwt prune

# Verify they're gone
gwt ls
```

## Scripts and Aliases

### Useful Aliases

```bash
# Add to your ~/.bashrc or ~/.zshrc
alias gw='gwt'
alias gwl='gwt ls'
alias gwp='gwt pwd'
```

### Cleanup All Worktrees Script

```bash
#!/bin/bash
# Clean up all worktrees except main

gwt ls | tail -n +3 | awk '{print $2}' | tr -d '()' | while read branch; do
    echo "Removing $branch..."
    gwt rm "$branch"
done

gwt prune
gwt ls
```

Save as `cleanup-worktrees.sh` and run with `bash cleanup-worktrees.sh`

### Quick Context Switch Script

```bash
#!/bin/bash
# Quick worktree switcher using fzf

selected=$(gwt ls | tail -n +3 | fzf | awk '{print $2}' | tr -d '()')
if [ -n "$selected" ]; then
    gwt cd "$selected"
fi
```

Save as `gwt-switch.sh` and run with `bash gwt-switch.sh`
