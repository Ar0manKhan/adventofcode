package main

import (
	"bufio"
	"log"
	"os"
)

// function to take file path input and return array of lines - By COPILOT
func extractInput(filePath string) []string {
	// read file
	file, err := os.Open(filePath)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	// extract lines
	scanner := bufio.NewScanner(file)
	var lines []string
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	// return lines
	return lines
}
