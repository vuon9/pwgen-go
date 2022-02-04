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
    Include ambiguous characters in the password
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
> pwgen-go -secure -column 20 10
LBCI_H0#M2q0H'2LYfnl    56M?O1ZY\G5h(YX6_^0J    R4O)1S:9[8vge"YGDZqa    SyJEEO5G03hI_Q!..n)$
*)G<~ZQG^Uy`q'F|]CKU    m`(2onEi&_S1el8F\G1/    f!1ou|e0SO#/S^a\BrB`    _R8O|K},KBy6gOI29UQr
IA!n[D5$mI'x~~X{DsyI    /ltSpA8\vT90Cxau!4qb
```

```
> pwgen-go -column -r="bzaX" 20 22
vS2E2GkEwW1MV2eOAnUm    yUoI7FJ7ygnR1BBloRfU    hTNS8jrkFOhfIN0UZTd0    J4yLWgFTuSwY0ROImvJI
UAu3U7i2FYOy7HiDkep8    cTRALWAYgP1tOmPIVyF5    OlmrcfkDoI1tstCIlDIF    BAKud0swVeE9O9DW4yqA
yjPOGMEL1e2wVLdR0JIg    f5ui30osIeV0dt7qGoYh    0fGeIYyTvIx0w6Feufrc    1qPxVWrKWFRcD1uAHA7Y
IQe1ycYCe0tR5KU3ii8A    gmf6iAiRUupUyusN1y4E    AryAOBoJJlxqUi3dutwB    Onmj0oeS4giDAyU3O1BI
rDN8My30HgDY0yHwsomO    EPT9yfNidJuwUHdRU1Us    T0wTG8Em7Kj8VVw1FU3Z    C0wuwfjeHT6BcklB5VKG
qxIAhFTNoPuecpGeD2Um    EOf0DdpIEtCxpYm1omWk
```

```
> pwgen-go -no-capitalize -r="bzaX" 5 10
04dr8   9q82t   9v1q4   k29nx   644ik   i320j   637o0   6vce0   3q31k   qq4n6
```