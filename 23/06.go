package main

import (
	"math"
	"regexp"
	"strconv"
	"strings"
)

type timeDist struct {
	time, dist int
}

func Solve06v1() {
	input := extractInput("input.txt")
	questions := getTimeDistData(input)
	answer := 1
	for _, q := range questions {
		answer *= (getMaxTimeForDist(q) - getMinTimeForDist(q) + 1)
	}
	println(answer)
}

func getMinTimeForDist(data timeDist) int {
	i, j := 1, data.time-1
	for i < j {
		mid := (j + i) / 2
		if mid*(data.time-mid) <= data.dist {
			i = mid + 1
		} else {
			j = mid
		}
	}
	return j
}

// it uses binary splitting method to find lowest time which can beat the distance
func getMaxTimeForDist(data timeDist) int {
	i, j := 1, data.time-1
	for i < j {
		mid := int(math.Ceil((float64(j) + float64(i)) / 2.0))
		if mid*(data.time-mid) <= data.dist {
			j = mid - 1
		} else {
			i = mid
		}
	}
	return i
}

// it uses binary splitting method to find highest time which can beat the distance
func getTimeDistData(data []string) []timeDist {
	spaceRegexp := regexp.MustCompile(`\s+`)
	timeData := spaceRegexp.Split(data[0], -1)[1:]
	distanceData := spaceRegexp.Split(data[1], -1)[1:]
	result := make([]timeDist, len(timeData))
	for i := range timeData {
		time, _ := strconv.Atoi(timeData[i])
		dist, _ := strconv.Atoi(distanceData[i])
		result[i] = timeDist{time, dist}
	}
	return result
}

// this is the math way of solving quadratic equation
// our equations boils down to
// x = (t +- sqrt(t^2 - 4d)) / 2   --- found this equation on copy using math
// where t is time and d is distance
// we get ceil value of lowest of them and floor value of highest of them
// we assume that t^2 - 4d will never be zero
func getPossiblityCountByMath(q timeDist) int {
	rootD := math.Sqrt(float64(q.time*q.time - 4*q.dist))
	lowerVal := int(math.Ceil((float64(q.time) - rootD) / 2))
	higherVal := int(math.Floor((float64(q.time) + rootD) / 2))
	return higherVal - lowerVal + 1
}

func Solve06v2() {
	input := extractInput("input.txt")
	timeStr := strings.Replace(input[0], " ", "", -1)
	distStr := strings.Replace(input[1], " ", "", -1)
	timeData, _ := strconv.Atoi(timeStr[strings.Index(timeStr, ":")+1:])
	distData, _ := strconv.Atoi(distStr[strings.Index(distStr, ":")+1:])
	question := timeDist{timeData, distData}
	println(getMaxTimeForDist(question) - getMinTimeForDist(question) + 1)
	println(getPossiblityCountByMath(question))
}
