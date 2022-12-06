Get-Content "$PSScriptRoot\input.txt" | Foreach-Object -Begin {
    $CompliantPasswords = 0
} -Process {
    $Object = $_ -Split ' '
    $Policy = [PSCustomObject] @{
        Range   = $Object[0]
        Character = $Object[1][0]
        Password = $Object[2]
    }
    $CharInstance = 0
    For ($i = 0; $i -lt $Policy.Password.Length; $i++) {
        If ($Policy.Password[$i] -eq $Policy.Character) {
            $CharInstance++
        }
    }
    If (($CharInstance -ge ($Policy.Range -Split '-')[0]) -and ($CharInstance -le ($Policy.Range -Split '-')[1])) {
        $CompliantPasswords++
    }
}
Write-Host "Answer: $CompliantPasswords"