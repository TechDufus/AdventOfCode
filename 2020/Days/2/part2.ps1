Get-Content "$PSScriptRoot\input.txt" | Foreach-Object -Begin {
    $CompliantPasswords = 0
} -Process {
    $Object = $_ -Split ' '
    $Policy = [PSCustomObject] @{
        Range   = $Object[0]
        Character = $Object[1][0]
        Password = $Object[2]
    }
    If (($Policy.Character -eq $Policy.Password[(($Policy.Range -Split '-')[0])-1]) -xor ($Policy.Character -eq $Policy.Password[(($Policy.Range -Split '-')[1])-1])) {
        $CompliantPasswords++
    }
}
Write-Host "Answer: $CompliantPasswords"