package main

import (
	"bufio"
	"encoding/json"
	"flag"
	"io"
	"io/ioutil"
	"os"
	"strings"
)

type LogJson struct {
  Logs []string
}

func (l *LogJson) appendToLogs (value string){
  l.Logs = append(l.Logs, value);
}

func defineFileExtension(filePath os.File, target string, fileDestination string) string{
	if fileDestination == "" {
		fileInfo, _ := filePath.Stat()
    fileExtension := ".text";
    if(target == "json"){
      fileExtension = ".json"
    }
		fileDestination =  strings.Split(fileInfo.Name(), ".")[0] + fileExtension
	}
  return fileDestination
}

func writeToFile(
	filePath os.File, target string, scanner bufio.Scanner, fileDestination string) error {
  fileDestination = defineFileExtension(filePath, target, fileDestination)
	fileDst, err := os.Create(fileDestination)
	if err != nil {
		return err
	}
	if target == "text" || target == "" {
		io.Copy(fileDst, &filePath)
		return nil
	}
	defer fileDst.Close()
  logJson := LogJson{}
	for scanner.Scan() {
    logJson.appendToLogs(scanner.Text())
	}
  content , _ := json.Marshal(logJson);
  _ = ioutil.WriteFile(fileDestination, content, 0644)
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
