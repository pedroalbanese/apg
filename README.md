# apg(2)
[![ISC License](http://img.shields.io/badge/license-ISC-blue.svg)](https://github.com/pedroalbanese/apg/blob/master/LICENSE.md) 
[![GoDoc](https://godoc.org/github.com/pedroalbanese/apg?status.png)](http://godoc.org/github.com/pedroalbanese/apg)
[![GitHub downloads](https://img.shields.io/github/downloads/pedroalbanese/apg/total.svg?logo=github&logoColor=white)](https://github.com/pedroalbanese/apg/releases)
[![Go Report Card](https://goreportcard.com/badge/github.com/pedroalbanese/apg)](https://goreportcard.com/report/github.com/pedroalbanese/apg)
[![GitHub go.mod Go version](https://img.shields.io/github/go-mod/go-version/pedroalbanese/apg)](https://golang.org)
[![GitHub release (latest by date)](https://img.shields.io/github/v/release/pedroalbanese/apg)](https://github.com/pedroalbanese/apg/releases)

### Automated Password Generator 
Automated Password Generator (APG) is a Linux program that helps webmasters and server administrators with creating random, secure passwords, through the SSH of server's. A wide range of Linux distros use APG in their repository. 

## Usage
```
Usage of apg:
  -H    Avoid ambiguous characters
  -L    Use lowercase characters (default true)
  -N    Use numeric characters (default true)
  -S    Use special characters
  -U    Use uppercase characters (default true)
  -l int
        Password length (default 12)
  -n int
        Number of passwords to generate (default 6)
  -seed string
        Specify a seed for random number generation
  -spell
        Spell passwords using phonetic alphabet
```

### Syntax:
```
./apg -L -N -U -spell
``` 

### Output:
```
vwItH3bGEEFZ victor-whiskey-India-tango-Hotel-THREE-bravo-Golf-Echo-Echo-Foxtrot-Zulu
xOGfGFkHL20e x-ray-Oscar-Golf-foxtrot-Golf-Foxtrot-kilo-Hotel-Lima-TWO-ZERO-echo
e5KLYOK6mA1m echo-FIVE-Kilo-Lima-Yankee-Oscar-Kilo-SIX-mike-Alpha-ONE-mike
DLBAm5HgPofk Delta-Lima-Bravo-Alpha-mike-FIVE-Hotel-golf-Papa-oscar-foxtrot-kilo
738e3h7k2ZQ3 SEVEN-THREE-EIGHT-echo-THREE-hotel-SEVEN-kilo-TWO-Zulu-Quebec-THREE
aQ2TdulqJ780 alfa-Quebec-TWO-Tango-delta-uniform-lima-quebec-Juliett-SEVEN-EIGHT-ZERO
```

This project is licensed under the ISC License.
##### Copyright (c) 2016-2023 Pedro F. Albanese - ALBANESE Research Lab.
