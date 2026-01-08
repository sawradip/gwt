package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/yourusername/gwt/commands"
)

func main() {
	args := os.Args[1:]

	if len(args) == 0 {
		commands.Help()
		return
	}

	cmd := args[0]
	cmdArgs := args[1:]

	var err error

	switch cmd {
	case "help", "-h", "--help":
		commands.Help()
		return

	case "ls", "list":
		err = commands.List()

	case "cd":
		if len(cmdArgs) == 0 {
			fmt.Fprintf(os.Stderr, "Error: worktree name required\n")
			os.Exit(1)
		}
		err = commands.Cd(cmdArgs[0])

	case "rm", "remove":
		if len(cmdArgs) == 0 {
			fmt.Fprintf(os.Stderr, "Error: branch name required\n")
			os.Exit(1)
		}
		err = commands.Remove(cmdArgs[0])

	case "pwd":
		err = commands.Pwd()

	case "prune":
		err = commands.Prune()

	case "config":
		err = commands.ConfigCmd(cmdArgs)

	default:
		// Assume it's a branch name for creation
		if !strings.HasPrefix(cmd, "-") {
			err = commands.Create(cmd)
		} else {
			fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n", cmd)
			fmt.Fprintf(os.Stderr, "Use 'gwt help' for usage information\n")
			os.Exit(1)
		}
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}
