package main

import (
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//partOne()
	partTwo()
}

func partOne() {
	file, err := os.ReadFile("./Day 3/day3-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var matrix [][]string

	var gears [][]Coordinates

	for _, str := range lines {
		var charSlice []string
		for _, char := range str {
			charSlice = append(charSlice, string(char))
		}
		matrix = append(matrix, charSlice)
	}
	sum := 0
	start := Coordinates{-1, -1}
	end := Coordinates{-1, -1}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] >= "0" && matrix[i][j] <= "9" {
				if start.x == -1 {
					start = Coordinates{i, j}
					end = Coordinates{i, j}
				} else {
					end = Coordinates{i, j}
				}
			} else {
				if start.x != -1 && end.x != -1 {
					if isPartNumber(start, end, matrix) {
						if isPartNearGear(start, end, matrix) {

						}
						//number := coordinatesToNumber(start, end, matrix)
						sum += number
					}
					start = Coordinates{-1, -1}
					end = Coordinates{-1, -1}
				}
			}
		}
	}
	println(sum)
}

func partTwo() {
	file, err := os.ReadFile("./Day 3/day3-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")
	var matrix [][]string

	for _, str := range lines {
		var charSlice []string
		for _, char := range str {
			charSlice = append(charSlice, string(char))
		}
		matrix = append(matrix, charSlice)
	}
	gearRatioSum := 0
	start := Coordinates{-1, -1}
	end := Coordinates{-1, -1}
	for i := 0; i < len(matrix); i++ {
		for j := 0; j < len(matrix[i]); j++ {
			if matrix[i][j] >= "0" && matrix[i][j] <= "9" {
				if start.x == -1 {
					start = Coordinates{i, j}
					end = Coordinates{i, j}
				} else {
					end = Coordinates{i, j}
				}
			} else {
				if start.x != -1 && end.x != -1 {
					if isPartNumber(start, end, matrix) {
						number := coordinatesToNumber(start, end, matrix)
						gearRatioSum += number
					}
					start = Coordinates{-1, -1}
					end = Coordinates{-1, -1}
				}
			}
		}
	}
	println(gearRatioSum)
}

type Coordinates struct {
	x int
	y int
}

func (c Coordinates) String() string {
	return fmt.Sprintf("Coordinates{x: %d, y: %d}", c.x, c.y)
}

func isPartNumber(start Coordinates, end Coordinates, matrix [][]string) bool {
	adjacentSymbols := getAdjacentSymbols(Coordinates{start.x, start.y}, Coordinates{end.x, end.y}, matrix)
	return len(adjacentSymbols) > 0
}

func isOutOfBounds(coords Coordinates, matrix [][]string) bool {
	return coords.x < 0 || coords.y < 0 || coords.x >= len(matrix) || coords.y >= len(matrix[0])
}

func isSymbol(c string) bool {
	return c == "#" || c == "%" || c == "$" || c == "&" || c == "*" || c == "/" || c == "@" || c == "=" || c == "+" || c == "-"
}

func getAdjacentSymbols(start Coordinates, end Coordinates, matrix [][]string) []Coordinates {
	var adjacentSymbols []Coordinates
	for i := start.x - 1; i <= end.x+1; i++ {
		for j := start.y - 1; j <= end.y+1; j++ {
			if isOutOfBounds(Coordinates{i, j}, matrix) {
				continue
			}
			if isSymbol(matrix[i][j]) {
				adjacentSymbols = append(adjacentSymbols, Coordinates{i, j})
			}
		}
	}
	return adjacentSymbols
}

func coordinatesToNumber(start Coordinates, end Coordinates, matrix [][]string) int {
	number := ""
	for i := start.y; i <= end.y; i++ {
		number += matrix[start.x][i]
	}
	fullNumber, _ := strconv.Atoi(number)
	return fullNumber
}

func isPartNearGear(start Coordinates, end Coordinates, matrix [][]string) bool {
	adjacentSymbols := getAdjacentSymbols(Coordinates{start.x, start.y}, Coordinates{end.x, end.y}, matrix)
	for _, coords := range adjacentSymbols {
		if matrix[coords.x][coords.y] == "*" {
			return true
		}
	}
	return false
}
