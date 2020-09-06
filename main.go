package main

import (
	"fmt"
	"math/rand"
	"time"
)

const (
	CONSONANT = 0x0001
	VOWEL     = 0x0002
	DIPTHONG  = 0x0004
	NOT_FIRST = 0x0008

	PW_DIGITS    = 1
	PW_UPPERS    = 2
	PW_SYMBOLS   = 4
	PW_AMBIGUOUS = 8
	PW_NO_VOWELS = 10

	pwDigits    = "0123456789"
	pwUppers    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	pwLowers    = "abcdefghijklmnopqrstuvwxyz"
	pwSymbols   = "!\"#$%&'()*+,-./:;<=>?@[\\]^_`{|}~"
	pwAmbiguous = "B8G6I1l0OQDS5Z2"
	pwVowels    = "01aeiouyAEIOUY"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	pwRand(nil, 16, PW_DIGITS|PW_UPPERS, nil)
}

func randomize(size int, chars string, t int) string {
	newPw := make([]byte, size)
	for i := range newPw {
		newPw[i] = chars[rand.Int63()%int64(len(chars))]
	}

	return string(newPw)
}

func pwRand(buf *string, size int, pwFlags byte, remove *string) {
	// var ch, chars, wChars string
	// var i, len, featureFlags int

	chars := ""
	if (pwFlags & PW_DIGITS) != 0 {
		chars += pwDigits
	}

	if (pwFlags & PW_UPPERS) != 0 {
		chars += pwUppers
	}

	if (pwFlags & PW_SYMBOLS) != 0 {
		chars += pwSymbols
	}

	chars += pwLowers

	pwds := make([]string, 16)

	for i := range pwds {
		pwds[i] = randomize(size, chars, i)
	}

	fmt.Println(pwFlags, pwFlags&PW_DIGITS, pwFlags&PW_UPPERS, pwFlags&PW_SYMBOLS)

	fmt.Printf("%v\n%v\n%v\n%v",
		pwds[0:4],
		pwds[4:8],
		pwds[8:12],
		pwds[12:16],
	)
}
