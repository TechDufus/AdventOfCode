$ValidPassports = 0
(Get-Content "$PSScriptRoot\input.txt" -raw) -split '\n\r' | Foreach-Object {
    $Passport = (($_ -Split ' ') -split [System.Environment]::NewLine)
    $RequiredFields = ($_.Contains('cid:')) ? 8 : 7
    If ($Passport.Count -ge $RequiredFields) {
        $PassportObject = [Hashtable]::new()
        $ValidFields = 0
        foreach ($Item in $Passport) {
            $Values = $Item.trim() -split ':'
            $PassportObject[$Values[0]] = $Values[1]
            switch($Values[0]) {
                'byr' {
                    if (($PassportObject[$Values[0]] -ge 1920) -and ($PassportObject[$Values[0]] -le 2002)) {
                        $ValidFields++
                    }
                }
                'iyr' {
                    if (($PassportObject[$Values[0]] -ge 2010) -and ($PassportObject[$Values[0]] -le 2020)) {
                        $ValidFields++
                    }
                }
                'eyr' {
                    if (($PassportObject[$Values[0]] -ge 2020) -and ($PassportObject[$Values[0]] -le 2030)) {
                        $ValidFields++
                    }
                }
                'hgt' {
                    Switch($PassportObject[$Values[0]].Substring($PassportObject[$Values[0]].Length - 2)) {
                        'in' {
                            $InValue = $PassportObject[$Values[0]].TrimEnd('in')
                            if (($InValue -ge 59) -and ($InValue -le 76)) {
                                $ValidFields++
                            }
                        }
                        'cm' {
                            $CmValue = $PassportObject[$Values[0]].TrimEnd('cm')
                            if (($CmValue -ge 150) -and ($CmValue -le 193)) {
                                $ValidFields++
                            }
                        }
                        DEFAULT {}
                    }
                }
                'hcl' {
                    $HclValue = $PassportObject[$Values[0]].Substring(1)
                    If ($HclValue -match "^[a-zA-Z0-9]{6}$") {
                        $ValidFields++
                    }
                }
                'ecl' {
                    $EclValue = $PassportObject[$Values[0]]
                    Switch($EclValue) {
                        'amb' {$ValidFields++}
                        'blu' {$ValidFields++}
                        'brn' {$ValidFields++}
                        'gry' {$ValidFields++}
                        'grn' {$ValidFields++}
                        'hzl' {$ValidFields++}
                        'oth' {$ValidFields++}
                        DEFAULT {}
                    }
                }
                'pid' {
                    If ($PassportObject[$Values[0]] -match "^\d{9}$") {
                        $ValidFields++
                    }
                }
                DEFAULT {}
            }
        }
        If ($RequiredFields -eq 8) {$RequiredFields = 7}
        If ($ValidFields -ge $RequiredFields) {
            $ValidPassports++
        }
    }
}
Write-Host "Answer: $ValidPassports"