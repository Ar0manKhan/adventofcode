package main

import (
	"bufio"
	"io"
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

func extractInputInByteGrid() [][]rune {
	grid := [][]rune{}
	for _, row := range extractInputByLine("input.txt") {
		r := make([]rune, len(row))
		for i, col := range row {
			r[i] = col
		}
		grid = append(grid, r)
	}
	return grid
}

func getInputOs() *os.File {
	file, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}
	return file
}

func getInputBuffer() *bufio.Reader {
	file := getInputOs()
	buf := bufio.NewReader(file)
	return buf
}

func getInputByDelimiter(delim byte) []string {
	result := []string{}
	scanner := getInputBuffer()

	for {
		chunk, err := scanner.ReadBytes(delim)
		if err != nil {
			if err == io.EOF && len(chunk) > 0 {
				result = append(result, string(chunk))
				break
			}
			panic(err)
		}
		result = append(result, string(chunk))
	}
	return result
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

func splitStringToIntDelimiter(in string, delim string) []int {
	result := []int{}
	for _, v := range strings.Split(in, delim) {
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
