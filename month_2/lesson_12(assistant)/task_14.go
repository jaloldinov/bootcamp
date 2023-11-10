package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

// Exercise 14: Create a program that uses goroutines and channels to perform concurrent file processing. Each goroutine should process a separate file and return the count of lines in that file.

func main() {
	filePaths := []string{"task_01.go", "task_02.go", "task_03.go"}
	ch := make(chan int)

	for _, filePath := range filePaths {
		go countLinesInFile(filePath, ch)
	}

	for range filePaths {
		lineCount := <-ch
		fmt.Printf("Total number of lines: %d\n", lineCount)

	}
}

func countLinesInFile(filePath string, ch chan int) {
	file, err := os.Open(filePath)
	if err != nil {
		fmt.Printf("Error opening file: %s\n", err)
		ch <- 0
		return
	}

	lineCount := 0
	reader := bufio.NewReader(file)

	for {
		_, err := reader.ReadString('\n')
		if err == io.EOF {
			break
		}
		lineCount++
	}

	ch <- lineCount
}
