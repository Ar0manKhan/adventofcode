package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

// function to take file path input and return array of lines - By COPILOT
func extractInputByLine(filePath string) []string {
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

func splitStringToInt(in string) []int {
	result := []int{}
	for _, v := range strings.Split(in, " ") {
		if len(v) == 0 {
			continue
		}
		val, err := strconv.Atoi(v)
		if err != nil {
			panic("Something went wrong while parsing integer")
		}
		result = append(result, val)
	}
	return result
}
