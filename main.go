package main

import (
	"fmt"
	"os"

	"github.com/yourusername/gwt/commands"
)

// Version is set via ldflags during build
var Version = "dev"

func main() {
	args := os.Args[1:]

	// Check for --version or -v first
	if len(args) > 0 && (args[0] == "--version" || args[0] == "-v") {
		fmt.Printf("gwt version %s\n", Version)
		return
	}

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

	case "version":
		fmt.Printf("gwt version %s\n", Version)
		return

	case "ls", "list":
		if hasHelpFlag(cmdArgs) {
			commands.HelpLs()
			return
		}
		err = commands.List()

	case "cd":
		if hasHelpFlag(cmdArgs) {
			commands.HelpCd()
			return
		}
		if len(cmdArgs) == 0 {
			fmt.Fprintf(os.Stderr, "Error: worktree name required\n")
			fmt.Fprintf(os.Stderr, "Usage: gwt cd <name>\n")
			fmt.Fprintf(os.Stderr, "Run 'gwt cd --help' for more information\n")
			os.Exit(1)
		}
		err = commands.Cd(cmdArgs[0])

	case "rm", "remove":
		if hasHelpFlag(cmdArgs) {
			commands.HelpRm()
			return
		}
		if len(cmdArgs) == 0 {
			fmt.Fprintf(os.Stderr, "Error: branch name required\n")
			fmt.Fprintf(os.Stderr, "Usage: gwt rm <name>\n")
			fmt.Fprintf(os.Stderr, "Run 'gwt rm --help' for more information\n")
			os.Exit(1)
		}
		err = commands.Remove(cmdArgs[0])

	case "pwd":
		if hasHelpFlag(cmdArgs) {
			commands.HelpPwd()
			return
		}
		err = commands.Pwd()

	case "prune":
		if hasHelpFlag(cmdArgs) {
			commands.HelpPrune()
			return
		}
		err = commands.Prune()

	case "config":
		if hasHelpFlag(cmdArgs) {
			commands.HelpConfig()
			return
		}
		err = commands.ConfigCmd(cmdArgs)

	case "create":
		if hasHelpFlag(cmdArgs) {
			commands.HelpCreate()
			return
		}
		if len(cmdArgs) == 0 {
			fmt.Fprintf(os.Stderr, "Error: branch name required\n")
			fmt.Fprintf(os.Stderr, "Usage: gwt create <branch>\n")
			fmt.Fprintf(os.Stderr, "Run 'gwt create --help' for more information\n")
			os.Exit(1)
		}
		err = commands.Create(cmdArgs[0])

	default:
		fmt.Fprintf(os.Stderr, "Error: unknown command '%s'\n", cmd)
		fmt.Fprintf(os.Stderr, "Run 'gwt --help' for usage information\n")
		os.Exit(1)
	}

	if err != nil {
		fmt.Fprintf(os.Stderr, "Error: %v\n", err)
		os.Exit(1)
	}
}

// hasHelpFlag checks if args contain -h or --help
func hasHelpFlag(args []string) bool {
	for _, arg := range args {
		if arg == "-h" || arg == "--help" {
			return true
		}
	}
	return false
}
