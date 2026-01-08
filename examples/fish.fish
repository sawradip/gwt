# gwt shell integration for Fish shell
#
# Add this to your ~/.config/fish/config.fish file to enable full gwt functionality
#
# After adding, reload your shell:
#   source ~/.config/fish/config.fish
#

# gwt wrapper function
function gwt
    if test "$argv[1]" = "cd"
        # For 'cd' command, actually change directory
        set target_path (command gwt cd $argv[2])
        if test $status -eq 0 -a -n "$target_path"
            cd "$target_path"
        end
    else
        # For all other commands, just run gwt normally
        command gwt $argv
    end
end

# Export the function so it's available
export -f gwt

# Fish completion for gwt
complete -c gwt -f
complete -c gwt -n "__fish_use_subcommand_from_list" -a "ls list cd rm remove pwd prune config help" -d "gwt command"
complete -c gwt -n "__fish_seen_subcommand_from cd; or __fish_seen_subcommand_from rm; or __fish_seen_subcommand_from remove" -a "(gwt ls 2>/dev/null | awk 'NR>2 {print \$2}' | tr -d '()')" -d "Worktree"
