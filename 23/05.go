package main

import (
	"fmt"
	"strconv"
	"strings"
)

func Solve05v1() {
	input := extractInput("input.txt")
	seeds := getSeeds(input[0])
	mappings := getOtherMappings(input[2:])
	current_min_seed := 99999999999999
	for _, seed := range seeds {
		current_min_seed = min(current_min_seed, getSeedToLocationMapping(seed, mappings))
	}
	fmt.Println(current_min_seed)
}

func getSeeds(input string) []int {
	seeds := []int{}
	for _, seed := range strings.Split(input, " ")[1:] {
		temp, _ := strconv.Atoi(seed)
		seeds = append(seeds, temp)
	}
	return seeds
}

type SourceDestRangeMapping struct {
	source, destination, increment_range int
}

func getOtherMappings(input []string) [][]SourceDestRangeMapping {
	var result [][]SourceDestRangeMapping
	current_input := []SourceDestRangeMapping{} // store mapping for one conversion
	for _, line := range input {
		// case 1: current line is empty - insert current inputs to result
		if len(line) == 0 {
			result = append(result, current_input)
			current_input = []SourceDestRangeMapping{}
			continue
		}
		// case 2: last line was empty (according to flag) - skip it
		if 'a' <= line[0] && line[0] <= 'z' {
			continue
		}
		// case 3: current line has number - input number
		data := strings.SplitN(line, " ", 3)
		var temp [3]int = [3]int{}
		for i := range data {
			x, _ := strconv.Atoi(data[i])
			temp[i] = x
		}
		current_input = append(current_input, SourceDestRangeMapping{temp[1], temp[0], temp[2]})
	}
	result = append(result, current_input)
	return result
}

func getSeedToLocationMapping(seed int, mappings [][]SourceDestRangeMapping) int {
	current_source := seed
	for _, mapping := range mappings {
		for _, each_map := range mapping {
			if current_source-each_map.source > -1 && current_source-each_map.source < each_map.increment_range {
				current_source = each_map.destination + (current_source - each_map.source)
				break
			}
		}
	}
	return current_source
}

func Solve05v2() {
	input := extractInput("input.txt")
	seedsMapping := getSeedsMapping(input[0])
	mappings := getOtherMappings(input[2:])
	current_min_seed := 99999999999999
	// fmt.Println(seedsMapping)
	for _, seed := range seedsMapping {
		a := seed.source
		r := seed.increment_range
		for {
			seed_map_range := getSeedToLocationRangeMapping(a, mappings)
			current_min_seed = min(seed_map_range.source, current_min_seed)
			if r <= seed_map_range.increment_range {
				break
			} else {
				a += seed_map_range.increment_range
				r -= seed_map_range.increment_range
			}
		}
		// current_min_seed = min(current_min_seed, getSeedToLocationMapping(seed, mappings))
	}
	fmt.Println(current_min_seed)
}

type SourceRangeMapping struct {
	source, increment_range int
}

func getSeedsMapping(input string) []SourceRangeMapping {
	result := []SourceRangeMapping{}
	seeds := getSeeds(input)
	for i := 0; i < len(seeds); i += 2 {
		// skipped index out of range error, in my case, it won't give error
		result = append(result, SourceRangeMapping{seeds[i], seeds[i+1]})
	}
	return result
}

// idea behind this function is that this function will be get called
// with seed and mapping and it will return the location for that seed but
// also the range upto which seed and location can be mapped
func getSeedToLocationRangeMapping(seed int, mappings [][]SourceDestRangeMapping) SourceRangeMapping {
	current_source := seed
	increment_range := 99999999999999
	for _, mapping := range mappings {
		free_range := 99999999999999
		found_mapping := false
		for _, each_map := range mapping {
			if current_source-each_map.source > -1 && current_source-each_map.source < each_map.increment_range {
				found_mapping = true
				increment_range = min(increment_range, each_map.source+each_map.increment_range-current_source)
				current_source = each_map.destination + (current_source - each_map.source)
				break
			} else if current_source < each_map.source {
				free_range = min(free_range, each_map.source-current_source)
			}
		}
		if !found_mapping {
			increment_range = min(increment_range, free_range)
		}
	}
	return SourceRangeMapping{current_source, increment_range}
}
