package main

import (
	"flag"
	"fmt"
	"os"
)

type userInput struct {
	output string
	target string
}

func parseCommand() (filepath string, target string, output string) {
	filepath = os.Args[1]
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&target, "t", "", "target of command")
	mySet.StringVar(&output, "o", "", "output of command")
	mySet.Parse(os.Args[2:])
	return filepath, target, output
}
func main() {
	filepath, target, output := parseCommand()
	fmt.Println(filepath, target, output)
}
