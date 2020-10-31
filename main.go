package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"flag"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
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

func defaultPwOptions() *pwOptions {
	return &pwOptions{
		pwLen: 16,
		numPw: 1,
	}
}

type pwOptions struct {
	pwLen int
	numPw int
}

func filterValidArgs(allOsArgs []string) []int {
	// The valid argument should be int
	isValidArgument := func(arg string) (int, bool) {
		isFlag := strings.HasPrefix("-", arg)
		if isFlag {
			return 0, false
		}

		val, err := strconv.Atoi(arg)
		return val, err == nil
	}

	validArgs := make([]int, 0)
	for _, rawArg := range allOsArgs {
		if val, ok := isValidArgument(rawArg); ok {
			validArgs = append(validArgs, val)
		}
	}

	return validArgs
}

func getOptions(pwArgs []int, pwOptions *pwOptions) *pwOptions {
	switch {
	case len(pwArgs) >= 1:
		pwOptions.pwLen = pwArgs[0]
		fallthrough
	case len(pwArgs) >= 2:
		pwOptions.numPw = pwArgs[1]
	}

	return pwOptions
}

func main() {
	var pwFlags byte

	capitalize := flag.Bool("c", false, "Include at least one capital letter in the password")
	nonCapitalize := flag.Bool("A", false, "Don't include capital letters in the password")
	number := flag.Bool("n", false, "Include at least one number in the password")
	nonNumber := flag.Bool("0", false, "Don't include numbers in the password")
	symbol := flag.Bool("y", false, "Include at least one special symbol in the password")
	help := flag.Bool("h", false, "Get help")
	// random := flag.Bool("s", false, "Generate completely random passwords")
	ambigous := flag.Bool("B", false, "Don't include ambiguous characters in the password")
	sha1 := flag.String("H", "", "Use sha1 hash of given file as a (not so) random generator")
	// column := flag.Bool("C", false, "Print the generated passwords in columns")
	// nonColumn := flag.Bool("1", false, "Don't print the generated passwords in columns")
	nonVowels := flag.Bool("v", false, "Do not use any vowels so as to avoid accidental nasty words")

	flag.Parse()

	pwOptions := getOptions(
		filterValidArgs(os.Args[1:]),
		defaultPwOptions(),
	)

	type generateFunc int
	const (
		typeRand = 1
		typeSha1 = 2
	)
	var typeFunc generateFunc = typeRand

	switch {
	case *capitalize:
		pwFlags |= PW_UPPERS
	case *nonCapitalize:
		pwFlags &^= PW_UPPERS
	case *number:
		pwFlags |= PW_DIGITS
	case *nonNumber:
		pwFlags ^= PW_DIGITS
	case *symbol:
		pwFlags |= PW_SYMBOLS
	// case *random:
	case *ambigous:
		pwFlags |= PW_AMBIGUOUS
	case *sha1 != "":
		typeFunc = typeSha1
	// case *column:
	// case *nonColumn:
	case *nonVowels:
		pwFlags |= PW_NO_VOWELS
	case *help:
		fmt.Println("Usage: pwgen-go [ OPTIONS ] [ pw_length | default: 16 ] [ num_pw | default: 1 ]")
		fmt.Println("Options supported by pwgen-go:")
		flag.PrintDefaults()
		os.Exit(0)
	}

	switch typeFunc {
	case typeRand:
		{
			// TODO: Build a CLI with options to generate password
			rand.Seed(time.Now().UnixNano())
			pwRand(nil, pwOptions.pwLen, pwOptions.numPw, pwFlags, nil)
		}
	case typeSha1:
		{
			splitted := strings.Split(*sha1, "#")
			if len(splitted) != 2 {
				println("err: Sha1 filepath and seed are invalid, should be path/sub_path/file.extension#seed")
				os.Exit(0)
			}

			filePath, seed := splitted[0], splitted[1]
			sha1File(filePath, seed)
		}
	}
}

func sha1File(filePath string, seed string) {
	f, err := os.Open(filePath)
	if err != nil {
		println("err: Couldn't open file")
		os.Exit(0)
	}
	defer f.Close()

	h := hmac.New(sha1.New, []byte(seed))
	if _, err := io.Copy(h, f); err != nil {
		println("err: Couldn't has file content")
		os.Exit(0)
	}

	var s string
	_, _ = h.Write([]byte(s))
	bs := h.Sum(nil)
	fmt.Printf("%x\n", bs)
}

func randomize(size int, chars string, t int) string {
	newPw := make([]byte, size)
	for i := range newPw {
		newPw[i] = chars[rand.Int63()%int64(len(chars))]
	}

	return string(newPw)
}

func pwRand(buf *string, size int, numPwds int, pwFlags byte, remove *string) {
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

	if (pwFlags & PW_AMBIGUOUS) != 0 {
		chars += pwAmbiguous
	}

	if (pwFlags & PW_NO_VOWELS) == 0 {
		chars += pwVowels
	}

	chars += pwLowers

	pwds := make([]string, numPwds)

	for i := range pwds {
		pwds[i] = randomize(size, chars, i)
		fmt.Println(pwds[i])
	}
}
