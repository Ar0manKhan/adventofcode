package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve02v1() {
	answer := 0
	const redLimit, greenLimit, blueLimit = 12, 13, 14
	for i, line := range extractInput("input.txt") {
		fetchIsPossible := true
		colorsData := strings.Split(line, ":")[1]
		for _, grab := range strings.Split(colorsData, ";") {
			redCount, greenCount, blueCount := 0, 0, 0
			grab = strings.Trim(grab, " ")
			// INFO: Can be improved a lot
			for _, colorInGrab := range strings.Split(grab, ", ") {
				spaceIdx := strings.Index(colorInGrab, " ")
				count, err := strconv.Atoi(colorInGrab[:spaceIdx])
				if err != nil {
					panic(err)
				}
				color := colorInGrab[spaceIdx+1:]
				switch color {
				case "red":
					redCount += count
				case "green":
					greenCount += count
				case "blue":
					blueCount += count
				default:
					panic("unknown color found" + color)
				}
			}
			if redCount > redLimit || greenCount > greenLimit || blueCount > blueLimit {
				fetchIsPossible = false
				break
			}
		}
		if fetchIsPossible {
			answer += i + 1
		}
	}
	fmt.Println(answer)
}

func Solve02v2() {
	answer := 0
	for _, line := range extractInput("input.txt") {
		colorsData := strings.Split(line, ":")[1]
		red, green, blue := 0, 0, 0
		for _, grab := range strings.Split(colorsData, ";") {
			grab = strings.Trim(grab, " ")
			for _, colorInGrab := range strings.Split(grab, ", ") {
				spaceIdx := strings.Index(colorInGrab, " ")
				count, err := strconv.Atoi(colorInGrab[:spaceIdx])
				if err != nil {
					panic(err)
				}
				color := colorInGrab[spaceIdx+1:]
				switch color {
				case "red":
					red = max(red, count)
				case "green":
					green = max(green, count)
				case "blue":
					blue = max(blue, count)
				default:
					panic("unknown color found" + color)
				}
			}
		}
		answer += (red * green * blue)
	}
	fmt.Println(answer)

}
