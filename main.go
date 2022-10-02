package main

import (
	"bufio"
	"flag"
	"fmt"
	"io"
	"os"
)

func writeToFile(
	filePath os.File, target string, scanner bufio.Scanner, fileDestination string) error {
	if fileDestination == "" {
		fileInfo, _ := filePath.Stat()
		fileDestination = "./" + fileInfo.Name()
	}
	fileDst, err := os.Create(fileDestination)
	if err != nil {
		return err
	}
	if target == "text" || target == "" {
		io.Copy(fileDst, &filePath)
		return nil
	}
	defer fileDst.Close()
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	return nil
}

func readAndWriteToFile(filePath string, target string, fileDestination string) error {
	file, err := os.Open(filePath)
	if err != nil {
		return err
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	writeToFile(*file, target, *scanner, fileDestination)
	if err := scanner.Err(); err != nil {
		return err
	}
	return nil
}

func parseCliCommand() (filePath string, target string, fileDestination string) {
	filePath = os.Args[1]
	mySet := flag.NewFlagSet("", flag.ExitOnError)
	mySet.StringVar(&target, "t", "", "File type of log file")
	mySet.StringVar(&fileDestination, "o", "", "File destination of log file")
	mySet.Parse(os.Args[2:])
	return filePath, target, fileDestination
}

func main() {
	filePath, target, fileDestination := parseCliCommand()
	readAndWriteToFile(filePath, target, fileDestination)
}
