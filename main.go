package main

import (
	"fmt"
	"flag"
	"os"
)

func main() {
	flag.Parse()

	if flag.NArg() < 2 {
		printError(notEnoughArgsMsg)
		printUsage()
		os.Exit(1)
	}
}

const notEnoughArgsMsg = "please provide two numbers"

func printError(msg string) {
	fmt.Fprintf(os.Stderr, "ERROR: %s\n", msg)
}

func printUsage() {
	fmt.Println("usage: evenodd startN endN")
}
