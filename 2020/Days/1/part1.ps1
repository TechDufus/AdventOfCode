[int[]] $Numbers = Get-Content "$PSScriptRoot\numbers.txt"

for ($i = 0; $i -lt $Numbers.count; $i++) {
    if ($Numbers[$i] -gt 2020) {
        Continue
    }
    for ($j = $i + 1; $j -lt $Numbers.Count; $j++) {
        $SumIJ = $Numbers[$i] + $Numbers[$j]
        if ($SumIJ -eq 2020) {
            Write-Host "Answer: $($Numbers[$i] * $Numbers[$j])"
        }
    }
}