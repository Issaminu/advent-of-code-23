package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	inputPath := "./Day 1/day1-input.txt"
	partOne(inputPath)
	partTwo(inputPath)
}

func partOne(path string) int {
	file, err := os.ReadFile(path)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	result := 0
	for _, line := range lines {
		first := ""
		second := ""
		for _, char := range line {
			_, err := strconv.Atoi(string(char))
			if err != nil {
				continue
			}
			if first == "" {
				first = string(char)
				continue
			}
			second = string(char)
		}
		if second == "" {
			second = first
		}
		concat := first + second
		num, _ := strconv.Atoi(concat)
		result += num
	}
	println(result)
	return result
}

type Pair struct {
	First  string
	Second string
}

func partTwo(inputPath string) {
	readFile, readError := os.ReadFile(inputPath)
	if readError != nil {
		panic(readError)
	}
	lines := strings.Split(string(readFile), "\n")
	parsedLines := make([]string, 0)
	for _, line := range lines {
		parsedLines = append(parsedLines, strToStrParsableByPartOne(line))
	}

	newPath := "./Day 1/day1-input-parsed.txt"
	writeFile, err := os.Create(newPath)
	if err != nil {
		fmt.Println("Error creating writeFile:", err)
		return
	}
	defer writeFile.Close()

	for _, parsedLine := range parsedLines {
		_, err := fmt.Fprintln(writeFile, parsedLine)
		if err != nil {
			fmt.Println("Error writing to file:", err)
			return
		}
	}
	partOne(newPath) // <---- Magic happens here
}

func strToStrParsableByPartOne(str string) string {
	pairs := []Pair{
		{"one", "on1e"},
		{"two", "tw2o"},
		{"three", "th3ree"},
		{"four", "fo4ur"},
		{"five", "fi5ve"},
		{"six", "s6ix"},
		{"seven", "se7ven"},
		{"eight", "e8ight"},
		{"nine", "n9ine"},
	}
	for _, pair := range pairs {
		str = strings.Replace(str, pair.First, pair.Second, -1)
	}
	return str
}
