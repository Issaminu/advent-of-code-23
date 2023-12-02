package main

import (
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.ReadFile("./Day 2/day2-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	gameIdSum := 0
	for _, line := range lines {
		skip := false

		pattern := regexp.MustCompile(`(\d+) red`)
		reds := pattern.FindAllStringSubmatch(line, -1)
		maxRed, _ := strconv.Atoi(reds[0][1])
		for _, red := range reds {
			redInt, _ := strconv.Atoi(red[1])
			maxRed = max(maxRed, redInt)
			if maxRed > 12 {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		pattern = regexp.MustCompile(`(\d+) green`)
		greens := pattern.FindAllStringSubmatch(line, -1)
		maxGreen, _ := strconv.Atoi(greens[0][1])
		for _, green := range greens {
			greenInt, _ := strconv.Atoi(green[1])
			maxGreen = max(maxGreen, greenInt)
			if maxGreen > 13 {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		pattern = regexp.MustCompile(`(\d+) blue`)
		blues := pattern.FindAllStringSubmatch(line, -1)
		maxBlue, _ := strconv.Atoi(blues[0][1])
		for _, blue := range blues {
			blueInt, _ := strconv.Atoi(blue[1])
			maxBlue = max(maxBlue, blueInt)
			if maxBlue > 14 {
				skip = true
				break
			}
		}
		if skip {
			continue
		}

		pattern = regexp.MustCompile(`Game (\d+)`)
		gameID := pattern.FindStringSubmatch(line)[1]
		gameIDInt, _ := strconv.Atoi(gameID)
		gameIdSum += gameIDInt
	}
	println(gameIdSum)
}

func partTwo() {
	file, err := os.ReadFile("./Day 2/day2-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	sum := 0
	for _, line := range lines {
		pattern := regexp.MustCompile(`(\d+) red`)
		reds := pattern.FindAllStringSubmatch(line, -1)
		maxRed, _ := strconv.Atoi(reds[0][1])
		for _, red := range reds {
			redInt, _ := strconv.Atoi(red[1])
			maxRed = max(maxRed, redInt)
		}

		pattern = regexp.MustCompile(`(\d+) green`)
		greens := pattern.FindAllStringSubmatch(line, -1)
		maxGreen, _ := strconv.Atoi(greens[0][1])
		for _, green := range greens {
			greenInt, _ := strconv.Atoi(green[1])
			maxGreen = max(maxGreen, greenInt)
		}

		pattern = regexp.MustCompile(`(\d+) blue`)
		blues := pattern.FindAllStringSubmatch(line, -1)
		maxBlue, _ := strconv.Atoi(blues[0][1])
		for _, blue := range blues {
			blueInt, _ := strconv.Atoi(blue[1])
			maxBlue = max(maxBlue, blueInt)
		}

		power := maxRed * maxGreen * maxBlue
		sum += power
	}
	println(sum)
}
