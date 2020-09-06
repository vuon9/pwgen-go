# pwgen-go
Password generator practice in Go. It is inspired by https://github.com/jbernard/pwgen

## The manual

This is following pwgen's manual: https://linux.die.net/man/1/pwgen
or can follow the usages by `pwgen-go -help` or `pwgen-go -h`.

## Download

```bash
go get -u github.com/vuon9/pwgen-go
```

## Usages

```md
Usage: pwgen-go [ OPTIONS ] [pw_length] [num_pw]
Options supported by pwgen-go:
 -h or -help
    Get help
 -c or -capitalize
    Include at least one capital letter in the password
 -A or -no-capitalize
    Don't include capital letters in the password
 -n or -numerals
    Include at least one number in the password
 -0 or -no-numerals
    Don't include numbers in the password
 -y or -symbol
    Include at least one special symbol in the password
 -r <chars> or --remove-chars=<chars>
     Remove characters from the set of characters to generate passwords
 -H or -sha1=path/to/file[#seed]
     Use sha1 hash of given file as a (not so) random generator
 -B or -ambigous
    Don't include ambiguous characters in the password
 -v or -no-vowels
    Do not use any vowels so as to avoid accidental nasty words
 -s or -secure
    Generate completely random passwords
 -column
     Print the generated passwords in columns
 -no-column
     Don't print the generated passwords in columns
 -vvv or -debug
    Enable debug mode
```
