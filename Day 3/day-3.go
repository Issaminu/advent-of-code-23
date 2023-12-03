package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
	"strings"
)

var gearAdjacentParts map[string][]int // Key: hash of gear coordinates, Value: adjacent part numbers

func main() {
	partOne()
	partTwo()
}

func partOne() {
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

func partTwo() {
	file, err := os.ReadFile("./Day 3/day3-input.txt")
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(file), "\n")

	var matrix [][]string
	gearAdjacentParts = make(map[string][]int)
	sum := 0

	for _, str := range lines {
		var charSlice []string
		for _, char := range str {
			charSlice = append(charSlice, string(char))
		}
		charSlice = append(charSlice, ".")
		matrix = append(matrix, charSlice)
	}

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
					if isAdjacentToGear(start, end, matrix) {
						checkAdjacentsForGear(start, end, matrix)
					}
				}
				start = Coordinates{-1, -1}
				end = Coordinates{-1, -1}
			}
		}
	}

	for _, valueArray := range gearAdjacentParts {
		if len(valueArray) == 2 {
			multiplication := 1
			for _, value := range valueArray {
				multiplication *= value
			}
			sum += multiplication
		}
	}

	println(sum)
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

func isAdjacentToGear(start Coordinates, end Coordinates, matrix [][]string) bool {
	adjacentSymbols := getAdjacentSymbols(Coordinates{start.x, start.y}, Coordinates{end.x, end.y}, matrix)
	for _, coords := range adjacentSymbols {
		if matrix[coords.x][coords.y] == "*" {
			return true
		}
	}
	return false
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

func checkAdjacentsForGear(start Coordinates, end Coordinates, matrix [][]string) {
	adjacentSymbols := getAdjacentSymbols(Coordinates{start.x, start.y}, Coordinates{end.x, end.y}, matrix)
	for _, coords := range adjacentSymbols {
		if matrix[coords.x][coords.y] == "*" {
			addCoordinate(coords, coordinatesToNumber(start, end, matrix))
		}
	}
}

func hashCoordinates(coord Coordinates) string {
	coordString := fmt.Sprintf("%d-%d", coord.x, coord.y)
	hasher := md5.New()
	hasher.Write([]byte(coordString))
	hash := hex.EncodeToString(hasher.Sum(nil))
	return hash
}

// Append `newPart` as new item in the Value array of the Key `gear` in the `gearAdjacentParts` map
func addCoordinate(gear Coordinates, newPart int) {
	hash := hashCoordinates(gear)
	gearAdjacentParts[hash] = append(gearAdjacentParts[hash], newPart)
}
