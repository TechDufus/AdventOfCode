$Terrain = Get-Content "$PSScriptRoot\input.txt"

$TreeCount = 0

#Loops down the square 1
For ($i,$j = 0,0; $i -lt $Terrain.count; $i++) {
    $Repeat = $Terrain[$i].Length
    If ($Terrain[$i][$j%$Repeat] -eq '#') {
        $TreeCount++
    }
    $j = $j+3
}
Write-Host "Answer: $TreeCount"