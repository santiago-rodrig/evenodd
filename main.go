package main

import (
	"errors"
	"fmt"
	"flag"
	"os"
	"strconv"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		printError(notEnoughArgsMsg)
		printUsage()
		os.Exit(1)
	}

	if argsAreInvalid() {
		printError(invalidArgsMsg)
		printUsage()
		os.Exit(1)
	}

	start, end, _ := parseArgs()
	evens, odds := 0, 0

	for i := start; i <= end; i++ {
		if i % 2 == 0 {
			evens++
		} else {
			odds++
		}
	}

	fmt.Printf(responseTemplate, start, end, start, end, evens, odds, evens+odds)
}

const responseTemplate = `for all the integers between %d and %d ([%d, %d])...
%d are even
%d are odd
%d is the number of integers
`

const (
	notEnoughArgsMsg = "please provide two numbers"
	invalidArgsMsg = "the given arguments are invalid"
)

var errUnparseableArgs = errors.New("some arguments are not integers")

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", msg)
}

func printUsage() {
	fmt.Println("usage: evenodd startN endN")
}

func parseArgs() (start int, end int, err error) {
	start, err = strconv.Atoi(flag.Arg(0))

	if err != nil {
		return 0, 0, errUnparseableArgs
	}

	end, err = strconv.Atoi(flag.Arg(1))

	if err != nil {
		return 0, 0, errUnparseableArgs
	}

	return
}

func argsAreInvalid() bool {
	start, end, err := parseArgs()

	if err != nil {
		return true
	}

	if start < 0 {
		return true
	}

	if end < start {
		return true
	}

	return false
}
