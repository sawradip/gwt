# gwt shell integration for Bash and Zsh
#
# Add this to your ~/.bashrc or ~/.zshrc file to enable full gwt functionality
#
# After adding, reload your shell:
#   source ~/.bashrc     # for bash
#   source ~/.zshrc      # for zsh
#

# gwt wrapper function
gwt() {
    if [ "$1" = "cd" ]; then
        # For 'cd' command, actually change directory
        local target_path=$(command gwt cd "$2")
        if [ $? -eq 0 ] && [ -n "$target_path" ]; then
            cd "$target_path"
        fi
    else
        # For all other commands, just run gwt normally
        command gwt "$@"
    fi
}

# Optional: Add bash completion if you have bash-completion installed
# This provides tab completion for gwt commands
if command -v _get_comp_words_by_ref &> /dev/null; then
    _gwt_completion() {
        local cur prev opts
        COMPREPLY=()
        cur="${COMP_WORDS[COMP_CWORD]}"
        prev="${COMP_WORDS[COMP_CWORD-1]}"

        if [[ ${COMP_CWORD} -eq 1 ]]; then
            # First argument - show available commands
            opts="ls list cd rm remove pwd prune config help"
            COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        elif [[ "${prev}" == "cd" || "${prev}" == "rm" || "${prev}" == "remove" ]]; then
            # For cd/rm/remove, we could list available worktrees
            # This is a simple implementation - you can enhance it
            opts=$(command gwt ls 2>/dev/null | awk 'NR>2 {print $2}' | tr -d '()')
            COMPREPLY=( $(compgen -W "${opts}" -- ${cur}) )
        fi
    }

    complete -F _gwt_completion gwt
fi
