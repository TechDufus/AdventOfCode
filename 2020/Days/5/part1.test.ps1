$AllSeatIDs = @()
Get-Content "$PSScriptRoot\test_input.txt" | Foreach-Object {
    $LastRange = 0..127
    #Calculate Row
    For ($i=0;$i -lt 7; $i++) {
        Switch($_[$i]) {
            'F' {
                $LastRange = @(($LastRange[0])..($LastRange[($LastRange.Count / 2)-1]))
            }
            'B' {
                $LastRange = @(($LastRange[($LastRange.Count / 2)])..($LastRange[-1]))
            }
        }
    }
    $Row = $LastRange[0]
    $LastRange = 0..7
    For ($j=7; $j -le 10; $j++) {
        Switch($_[$j]) {
            'L' {
                $LastRange = @(($LastRange[0])..($LastRange[($LastRange.Count / 2)-1]))
            }
            'R' {
                $LastRange = @(($LastRange[($LastRange.Count / 2)])..($LastRange[-1]))
            }
        }
    }
    $Column = $LastRange[0]
    $SeatID = ($Row * 8) + $Column
    $AllSeatIDs += $SeatID
}

"Answer: $(($AllSeatIDs | Measure-Object -Maximum).Maximum)"

