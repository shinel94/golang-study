package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"sync"
)

var Args []string
var wg sync.WaitGroup

func checkFile(taskCh chan string, targetWord string, ch chan string) {
	for fileName := range taskCh {
		var result = fileName + "\n"
		result += "-----------------------------\n"
		filePointer, err := os.Open(fileName)

		if err != nil {
			result += "File Can't Open"
		} else {
			reader := bufio.NewScanner(filePointer)
			lineNo := 1
			for reader.Scan() {
				line := strings.TrimSpace(reader.Text())
				if strings.Contains(line, targetWord) {
					result += fmt.Sprintf("\t\t %05d \t %s\n", lineNo, line)
				}
				lineNo++
			}
		}
		filePointer.Close()

		result += "-----------------------------\n"
		ch <- result
	}
	ch <- "DONE"
	wg.Done()

}

func resultPrinter(ch chan string) {
	done := 0
	resultPointer, err := os.Create("result.txt")
	if err != nil {
		fmt.Println("Can't Open File")
		wg.Done()
		return
	}
	for result := range ch {
		if result == "DONE" {
			done++
		} else {
			fmt.Fprintln(resultPointer, result)
		}
		if done == 3 {
			break
		}
	}
	wg.Done()
}

func main() {
	Args = os.Args
	if len(Args) < 3 {
		fmt.Println("Must Provide two Argument {target} {filename}")
		return
	}
	var targetWord = Args[1]
	var fileNameList = Args[2:]
	var totalFileList []string
	if len(fileNameList) < 1 {
		fmt.Println("Must Provide Target File Name or WildCard")
		return
	}
	for _, fileName := range fileNameList {
		fileList, err := filepath.Glob(fileName)
		if err != nil {
			fmt.Println("File Find Error - ", fileName)
		}
		totalFileList = append(totalFileList, fileList...)
	}
	wg.Add(4)
	var taskCh = make(chan string)
	var resultCh = make(chan string, 3)
	go checkFile(taskCh, targetWord, resultCh)
	go checkFile(taskCh, targetWord, resultCh)
	go checkFile(taskCh, targetWord, resultCh)
	go resultPrinter(resultCh)

	for _, fileName := range totalFileList {
		taskCh <- fileName
	}
	close(taskCh)
	wg.Wait()
	fmt.Println("Please Check file : ./result.txt")
}
