$ValidCount = 0
(Get-Content "$PSScriptRoot\input.txt" -raw) -split '\n\r' | Foreach-Object {
    $Passport = (($_ -Split ' ') -split [System.Environment]::NewLine)
    $RequiredFields = ($_.Contains('cid:')) ? 8 : 7
    If ($Passport.Count -ge $RequiredFields) {
        $ValidCount++
    }
}
Write-Host "Answer: $ValidCount"