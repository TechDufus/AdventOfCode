[int[]] $Numbers = Get-Content "$PSScriptRoot\numbers.txt"

for ($i = 0; $i -lt $Numbers.count; $i++) {
    if ($Numbers[$i] -gt 2020) {
        Continue
    }
    for ($j = $i + 1; $j -lt $Numbers.Count; $j++) {
        $SumIJ = $Numbers[$i] + $Numbers[$j]
        if ($SumIJ -gt 2020) {
            Continue
        }
        for ($k = $j + 1; $k -lt $Numbers.Count; $k++) {
            $SumIJK = $Numbers[$i] + $Numbers[$j] + $Numbers[$k]
            if ($SumIJK -eq 2020) {
                Write-Host "Anser: $($Numbers[$i] * $Numbers[$j] * $Numbers[$k])"
            }
        }
    }
}