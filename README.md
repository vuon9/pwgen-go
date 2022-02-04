# pwgen-go
![Go](https://github.com/vuon9/pwgen-go/workflows/Go/badge.svg)

## Download

```bash
go install github.com/vuon9/pwgen-go
```

## Debug

```PWGEN_GO_DEBUG=true```

## Usages

```md
> pwgen-go -help
Usage: pwgen-go [ OPTIONS ] [pw_length] [num_pw]
Options supported by pwgen-go:
 -ambiguous, -B (default: false)
    Don't include ambiguous characters in the password
 -column,  (default: false)
    Print the generated passwords in columns
 -debug, -vvv (default: false)
    Enable debug mode
 -help, -h (default: false)
    Get help
 -no-capitalize, -A (default: false)
    Don't include capital letters in the password
 -no-numerals, -0 (default: false)
    Don't include numbers in the password
 -no-vowels, -v (default: false)
    Don't include any vowels so as to avoid accidental nasty words
 -remove-chars, -r (default: <empty>)
    Remove characters from the set of characters to generate passwords (ex: -r <chars> or --remove-chars=<chars>)
 -secure, -s (default: false)
    Generate completely random passwords
 -sha1, -H (default: <empty>)
    Use sha1 hash of given file as a (not so) random generator (ex: -H or -sha1=path/to/file[#seed])
 -symbol, -y (default: false)
    Include at least one special symbol in the password
```

## Examples

```
> go run . -column -remove-chars="bzaa" 20 22
KxTC7fuirIBOOwUiuou0    uHStEiGFKhecGE1In8uj    IEe1LkwwiWCfdUmiQYSE    yc7MtORu3e5URPtOZSJ5
LFVoYurmYnEFnsOXvRBu    9CCthoPI4cxfyy1Hglv0    D1U3wE5GiRFU8TwCHG0m    wisY928pxY9Zv94Gew31
eWe2e7SGYeE95tk1Ye6A    SRODOEUUSrsLY6o31Npo    EY9SvRsySx0LxOvtpI03    iAvuEknMcHoRTKHmsGjp
sCqjUwmQRC5cTO4ZHlym    thL1Oo97KeAx3MX9sFhq    cPUBtOsgXZovQLSukjoN    p6KYXiVYcpnSmA0JHdPj
qN2HRWBPedHrkIButjUr    i1igUwjGBUR3Auu4UMfV    DO4WEdIon1Bwv6Ip4I5R    iOK4VMneREZDATUE2LLV
4uTAZyTiWIHj181sHTcU    ueAMTFximso9oc3iHwnj
```