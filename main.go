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

func main() {
	filePath := os.Args[1]
	var target string
	var output string
	mySet := flag.NewFlagSet("", flag.ExitOnError)
  mySet.StringVar(&target, "t", "", "target of command")
  mySet.StringVar(&output, "o", "", "output of command")
  mySet.Parse(os.Args[2:])
	fmt.Println(filePath)
	fmt.Println(target)
	fmt.Println(output)
}
