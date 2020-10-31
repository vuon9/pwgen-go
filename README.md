# pwgen-go
Password generator practice in Go

Download
```bash
go get -u github.com/vuongggggg/pwgen-go
```

## Usages

```bash

# Generate password with length = 32, total = 5
pwgen-go 32 5
# EEud1EmkguaOe0IeIlOzdinIlrzte0Av
# flovabponeugc1lnpgOEzbOpkzzuneqe
# oEqeyirdsrviruatslUmwglwfsbEEmyY
# suUhsvypUwieslYu1parljIYuuidghys
# pjE1kanujxulqmEjeojyeUkUonxq0iUw

# generate SHA1 with filepath & seed
pwgen-go -H="D:\go\src\github.com\vuongggggg\pwgen-go\main.go#abcww"
# output: a4392a17f6c6c6938dadb74a1ce531065880d597
```

Inspired by https://github.com/jbernard/pwgen
