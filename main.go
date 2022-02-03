package main

import (
	"crypto/hmac"
	"crypto/sha1"
	"errors"
	"fmt"
	"io"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

const (
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
		pwLen: 20,
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
		if val, ok := isValidArgument(rawArg); ok && val > 0 {
			validArgs = append(validArgs, val)
		}
	}

	return validArgs
}

func getOptions(pwArgs []int, pwOptions *pwOptions) (*pwOptions, error) {
	if len(pwArgs) >= 1 {
		pwOptions.pwLen = pwArgs[0]
	}

	if len(pwArgs) >= 2 {
		pwOptions.numPw = pwArgs[1]
	}

	if pwOptions.numPw < 1 || pwOptions.pwLen < 5 {
		return nil, errors.New("The number of password should be >= 1 and the length of password shoul be >= 5")
	}

	return pwOptions, nil
}

var (
	cmdCapitalize   cmdName = "capitalize"
	cmdNoCapitalize cmdName = "no-capitalize"
	cmdHelp         cmdName = "help"
	cmdNumerals     cmdName = "numerals"
	cmdNoNumerals   cmdName = "no-numerals"
	cmdSymbol       cmdName = "symbol"
	cmdRemoveChars  cmdName = "remove-chars"
	cmdSha1         cmdName = "sha1"
	cmdAmbiguous    cmdName = "ambiguous"
	cmdNoVowels     cmdName = "no-vowels"
	cmdSecure       cmdName = "secure"
	cmdColumn       cmdName = "column"
	cmdNoColumn     cmdName = "no-column"
	cmdDebug        cmdName = "debug"
)

func main() {
	commands := NewCommandController(
		NewItems(
			NewBoolCommand(cmdHelp, "h", "false", "Get help"),
			NewBoolCommand(cmdCapitalize, "c", "false", "Include at least one capital letter in the password"),
			NewBoolCommand(cmdNoCapitalize, "A", "true", "Don't include capital letters in the password"),
			NewBoolCommand(cmdNumerals, "n", "false", "Include at least one number in the password"),
			NewBoolCommand(cmdNoNumerals, "0", "true", "Don't include numbers in the password"),
			NewBoolCommand(cmdSymbol, "y", "false", "Include at least one special symbol in the password"),
			NewStringCommand(cmdRemoveChars, "r", "", "Remove characters from the set of characters to generate passwords (ex: -r <chars> or --remove-chars=<chars>)"),
			NewStringCommand(cmdSha1, "H", "", "Use sha1 hash of given file as a (not so) random generator (ex: -H or -sha1=path/to/file[#seed])"),
			NewBoolCommand(cmdAmbiguous, "B", "false", "Don't include ambiguous characters in the password"),
			NewBoolCommand(cmdNoVowels, "v", "false", "Do not use any vowels so as to avoid accidental nasty words"),
			NewBoolCommand(cmdSecure, "s", "false", "Generate completely random passwords"),
			NewBoolCommand(cmdColumn, "", "false", "Print the generated passwords in columns"),
			NewBoolCommand(cmdNoColumn, "", "true", "Don't print the generated passwords in columns"),
			NewBoolCommand(cmdDebug, "vvv", "false", "Enable debug mode"),
		),
		WithUsageHeader("Usage: pwgen-go [ OPTIONS ] [pw_length] [num_pw]\nOptions supported by pwgen-go:"),
	)

	commands.Ready()

	var hasSha1 string = commands.GetString("sha1")
	if hasSha1 != "" {
		splitted := strings.Split(hasSha1, "#")
		if len(splitted) != 2 {
			println("err: Sha1 filepath and seed are invalid, should be path/sub_path/file.extension#seed")
			os.Exit(0)
		}

		filePath, seed := splitted[0], splitted[1]
		sha1File(filePath, seed)
		os.Exit(0)
	}

	pwOptions, err := getOptions(filterValidArgs(os.Args[0:]), defaultPwOptions())
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	var pwFlags byte
	var withColumn bool
	removeChars := ""
	var debug bool

	switch {
	case commands.GetBool(cmdCapitalize):
		pwFlags |= PW_UPPERS
		fallthrough
	case commands.GetBool(cmdNoCapitalize):
		pwFlags &^= PW_UPPERS
	case commands.GetBool(cmdNumerals):
		pwFlags |= PW_DIGITS
		fallthrough
	case commands.GetBool(cmdNoNumerals):
		pwFlags ^= PW_DIGITS
		fallthrough
	case commands.GetBool(cmdSecure):
		pwFlags = PW_DIGITS | PW_UPPERS
	case commands.GetBool(cmdSymbol):
		pwFlags |= PW_SYMBOLS
	case commands.GetBool(cmdAmbiguous):
		pwFlags |= PW_AMBIGUOUS
	case commands.GetBool(cmdNoVowels):
		pwFlags |= PW_NO_VOWELS | PW_DIGITS | PW_UPPERS
	case commands.GetBool(cmdNoColumn):
		withColumn = false
	case commands.GetBool(cmdColumn):
		withColumn = true
	case commands.GetString(cmdRemoveChars) != "":
		removeChars = commands.GetString(cmdRemoveChars)
	case commands.GetBool(cmdDebug):
		debug = true
	case commands.GetBool(cmdHelp):
		commands.Usage()
		os.Exit(0)
	}

	// Randomize passwords by flags & eligible chars
	var t1 time.Time
	if debug {
		t1 = time.Now()
	}

	passwords, err := pwRand(nil, pwOptions, eligibleChars(pwFlags, removeChars))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}

	// Print passwords by column or no column
	const itemsPerColumn = 4
	for i, pwd := range passwords {
		fmt.Printf("%s\t", pwd)
		if withColumn && i+1 >= itemsPerColumn && (i+1)%itemsPerColumn == 0 {
			fmt.Print("\n")
		}
	}

	if debug {
		fmt.Println("\nElapsed time: ", time.Since(t1))
	}
	os.Exit(0)
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

func eligibleChars(pwFlags byte, removeChars string) string {
	chars := pwLowers
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

	for _, rChar := range removeChars {
		chars = strings.ReplaceAll(chars, string(rChar), "")
	}

	return chars
}

func pwRand(buf *string, pwOptions *pwOptions, chars string) ([]string, error) {
	if len(chars) == 0 {
		return nil, errors.New("no available chars for generating password")
	}

	randomize := func(size int, chars string) []byte {
		newPw := make([]byte, size)
		for i := range newPw {
			newPw[i] = chars[rand.Int63()%int64(len(chars))]
		}

		return newPw
	}

	rand.Seed(time.Now().UnixNano())

	var wg sync.WaitGroup
	wg.Add(pwOptions.numPw)

	passwords := make([]string, pwOptions.numPw)
	for i := range passwords {
		go func(i int) {
			defer wg.Done()
			passwords[i] = string(randomize(pwOptions.pwLen, chars))
		}(i)
	}
	wg.Wait()

	return passwords, nil
}
