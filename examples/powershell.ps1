# gwt shell integration for PowerShell
#
# Add this to your PowerShell profile to enable full gwt functionality
#
# To find your profile location, run:
#   $PROFILE
#
# To edit your profile:
#   notepad $PROFILE
#
# After adding, reload your PowerShell profile:
#   . $PROFILE
#

# gwt wrapper function for PowerShell
function gwt {
    param(
        [Parameter(Position = 0)]
        [string]$Command,

        [Parameter(Position = 1, ValueFromRemainingArguments = $true)]
        [string[]]$Arguments
    )

    if ($Command -eq "cd") {
        # For 'cd' command, actually change directory
        $targetPath = & gwt.exe cd $Arguments[0] 2>$null
        if ($LASTEXITCODE -eq 0 -and $targetPath) {
            Set-Location $targetPath
        }
    }
    else {
        # For all other commands, just run gwt
        & gwt.exe $Command $Arguments
    }
}

# Make gwt a cmdlet-like function
Export-ModuleMember -Function gwt

# Optional: Register argument completer for tab completion
$scriptBlock = {
    param($wordToComplete, $commandAst, $cursorPosition)

    # Get available worktrees from gwt ls
    $worktrees = @()
    try {
        $output = & gwt.exe ls 2>$null
        $worktrees = $output | Select-Object -Skip 2 | ForEach-Object {
            $_ -replace '^\s*\*?\s*', '' | -replace '\s*\(.*\)$', ''
        } | Where-Object { $_ -and $_ -notmatch '^main' }
    }
    catch {
        # If gwt command fails, just return empty
    }

    # Return matching worktrees
    $worktrees | Where-Object { $_ -like "$wordToComplete*" } | ForEach-Object {
        [System.Management.Automation.CompletionResult]::new($_, $_, 'ParameterValue', $_)
    }
}

# Register the argument completer
Register-ArgumentCompleter -CommandName gwt -ScriptBlock $scriptBlock
