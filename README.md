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
    Generate random passwords with digits, symbols, ambiguous, uppers
 -sha1, -H (default: <empty>)
    Use sha1 hash of given file as a (not so) random generator (ex: -H or -sha1=path/to/file[#seed])
 -symbol, -y (default: false)
    Include at least one special symbol in the password
```

## Examples

```
> go run . -secure -column 20 10
LBCI_H0#M2q0H'2LYfnl    56M?O1ZY\G5h(YX6_^0J    R4O)1S:9[8vge"YGDZqa    SyJEEO5G03hI_Q!..n)$
*)G<~ZQG^Uy`q'F|]CKU    m`(2onEi&_S1el8F\G1/    f!1ou|e0SO#/S^a\BrB`    _R8O|K},KBy6gOI29UQr
IA!n[D5$mI'x~~X{DsyI    /ltSpA8\vT90Cxau!4qb
```