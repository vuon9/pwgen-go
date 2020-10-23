package main

import (
	"flag"
	"fmt"
	"math/rand"
	"os"
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
	pwSymbols   = "!\"#$%&'()*+,-./:<=>?@[\\]^_`{|}~"
	pwAmbiguous = "B8G6I1l0OQDS5Z2"
	pwVowels    = "01aeiouyAEIOUY"
)

func main() {
	println("Usage: pwgen-go [ OPTIONS ] [ pw_length ] [ num_pw ]")
	println("Options supported by pwgen-go:")

	var pwFlags byte

	capitalize := flag.Bool("c", false, "Include at least one capital letter in the password")
	nonCapitalize := flag.Bool("A", false, "Don't include capital letters in the password")
	number := flag.Bool("n", false, "Include at least one number in the password")
	nonNumber := flag.Bool("0", false, "Don't include numbers in the password")
	symbol := flag.Bool("y", false, "Include at least one special symbol in the password")
	// random := flag.Bool("s", false, "Generate completely random passwords")
	ambigous := flag.Bool("B", false, "Don't include ambiguous characters in the password")
	help := flag.Bool("h", false, "Print a help message")
	// sha1 := flag.Bool("H", false, "Use sha1 hash of given file as a (not so) random generator")
	// column := flag.Bool("C", false, "Print the generated passwords in columns")
	// nonColumn := flag.Bool("1", false, "Don't print the generated passwords in columns")
	nonVowels := flag.Bool("v", false, "Do not use any vowels so as to avoid accidental nasty words")

	switch {
	case *capitalize:
		pwFlags |= PW_UPPERS
	case *nonCapitalize:
		pwFlags ^= PW_UPPERS
	case *number:
		pwFlags |= PW_DIGITS
	case *nonNumber:
		pwFlags ^= PW_DIGITS
	case *symbol:
		pwFlags |= PW_SYMBOLS
	// case *random:
	case *ambigous:
		pwFlags |= PW_AMBIGUOUS
	// case *sha1:
	// case *column:
	// case *nonColumn:
	case *nonVowels:
		pwFlags |= PW_NO_VOWELS
	case *help:
		flag.PrintDefaults()
		os.Exit(1)
	}

	flag.Parse()

	// TODO: Build a CLI with options to generate password
	rand.Seed(time.Now().UnixNano())
	pwRand(nil, 16, pwFlags, nil)
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

	// fmt.Println(pwFlags, pwFlags&PW_DIGITS, pwFlags&PW_UPPERS, pwFlags&PW_SYMBOLS)

	// TODO: Clean this debugging
	fmt.Printf("%v\n%v\n%v\n%v",
		pwds[0:4],
		pwds[4:8],
		pwds[8:12],
		pwds[12:16],
	)
}
