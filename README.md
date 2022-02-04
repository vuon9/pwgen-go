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
 -no-capitalize, -A (default: true)
    Don't include capital letters in the password
 -no-numerals, -0 (default: true)
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
> pwgen-go -column -remove-chars="bzaa" 20 22
ldtovroffwuvklekjnvm    nenkcpheqdoipdthvnjm    fucgnkuytggileohwsve    xgwdoqvwouqmkslsiqvh
olirhswrjkquohpxsfit    ofokhpnksdyedktxvxtn    prlkkmsecvvdkwwfgtlk    rcnfokrewyerjfcjndrs
qniywvgejecychuetrqh    cxsquoueyvxjnlownnwx    fsyopntknulmniqeinwx    leemsrwtfmrmmmvxtpjg
mispnpnoemlhitgrtjhh    odkunevrtcvrsrysipfp    umixnivwqhilqwuskxhw    dcohkgoukhrysdyfwmho
fyenfrysthmfqirfvonx    mctxoocsgskxevgmhiev    uokntlwwrmthjlwrghqk    sgfhfxfkskdljoxtmwfv
hxigrlmjrqikflupkvcx    ntfjmdqcilchrthwwmyp
```