package main

func Solve01v1() {
	sum := 0
	var first, last int
	var m48 int
	lines := extractInput("input.txt")
	for _, line := range lines {
		first, last = -1, -1
		for _, c := range line {
			m48 = int(c - 48)
			if m48 < 10 {
				if first == -1 {
					first = m48
				}
				last = m48
			}
		}
		sum += first*10 + last
	}
	println(sum)
}

func Solve01v2() {
	digitString := []string{"zero", "one", "two", "three", "four", "five", "six", "seven", "eight", "nine"}
	currentDigits := []int8{0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	sum := 0
	var first, last, extractedNumber int8
	for _, line := range extractInput("input.txt") {
		first, last = -1, -1
		for _, c := range line {
			extractedNumber = -1
			for i := range digitString {
				if digitString[i][currentDigits[i]] != byte(c) {
					currentDigits[i] = 0
				}

				if digitString[i][currentDigits[i]] == byte(c) {
					currentDigits[i]++
					if currentDigits[i] == int8(len(digitString[i])) {
						extractedNumber = int8(i)
						currentDigits[i] = 0
					}
				}
			}
			if extractedNumber == -1 && c-48 >= 0 && c-48 < 10 {
				if c-48 > 0 && c-48 < 10 {
					extractedNumber = int8(c - 48)
				}
			}
			// test code
			// if extractedNumber != -1 {
			// 	print(extractedNumber)
			// }
			//
			if extractedNumber != -1 {
				if first == -1 {
					first = extractedNumber
				}
				last = extractedNumber
			}
		}
		// print("\n") // test code
		sum += int(first*10 + last)
		// println(sum)
	}
	println(sum)
}
