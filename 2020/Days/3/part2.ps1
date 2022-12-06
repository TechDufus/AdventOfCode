$Terrain = Get-Content "$PSScriptRoot\input.txt"

$Count = @{
    'r1d1' = 0
    'r3d1' = 0
    'r5d1' = 0
    'r7d1' = 0
    'r1d2' = 0
}

$Repeat = $Terrain[0].Length
#Loops down the square 1
For ($i = 1; $i -lt $Terrain.count; $i++) {
    #Right 1, down 1
    If ($Terrain[$i][$i%$Repeat] -eq '#') {
        $Count.'r1d1'++
    }
    # Right 3, down 1
    If ($Terrain[$i][($i*3)%$Repeat] -eq '#') {
        $Count.'r3d1'++
    }
    # Right 5, down 1
    If ($Terrain[$i][($i*5)%$Repeat] -eq '#') {
        $Count.'r5d1'++
    }
    # Right 7, down 1
    If ($Terrain[$i][($i*7)%$Repeat] -eq '#') {
        $Count.'r7d1'++
    }
    #Ensure we don't go above $Terrain[$Terrain.Count]
    $ii = $i*2
    If ($ii -le $Terrain.count) {
        # Right 1, down 2
        If ($Terrain[$($ii)][$i%$Repeat] -eq '#') {
            $Count.'r1d2'++
        }
    }
}
$Answer = 1
$Count.Keys | Foreach-Object {
    $Answer = $Answer * $Count.$_
}
Write-Host "Answer: $Answer"