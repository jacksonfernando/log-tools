package main

import (
	"bufio"
	"flag"
	"fmt"
	"os"
)

func readAndWriteFile(filePath string, target string) error{
  if(target == "" || target == "text"){
  }
  file, err := os.Open(filePath)
  if err != nil{
    return err;
  }
  defer file.Close()
  scanner := bufio.NewScanner(file)
  for scanner.Scan() {
    fmt.Println(scanner.Text())
  }
  if err := scanner.Err(); err != nil{
    return err;
  }
  return nil
}

func parseCliCommand() (filePath string, target string, output string) {
	filePath = os.Args[1]
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&target, "t", "", "Target of command")
	mySet.StringVar(&output, "o", "", "Output of command")
	mySet.Parse(os.Args[2:])
	return filePath, target, output
}
func main() {
	filePath, target, _ := parseCliCommand()
  readAndWriteFile(filePath, target)
}
