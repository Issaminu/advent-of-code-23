package main

import (
	"math"
	"os"
	"strings"
)

func main() {
	partOne()
	partTwo()
}

func partOne() {
	file, err := os.ReadFile("./Day 4/day4-input.txt")
	if err != nil {
		panic(err)
	}
	cards := strings.SplitAfter(string(file), "\n")
	lastIndex := len(cards) - 1
	if cards[lastIndex][len(cards[lastIndex])-1] != '\n' {
		cards[lastIndex] += "\n"
	}

	sum := 0
	for _, card := range cards {
		wins := getWinnings(card)
		if wins > 0 {
			sum += int(math.Pow(2, float64(wins-1)))
		}
	}
	println(sum)
}

func partTwo() {
	file, err := os.ReadFile("./Day 4/day4-input.txt")
	if err != nil {
		panic(err)
	}
	cards := strings.SplitAfter(string(file), "\n")
	lastIndex := len(cards) - 1
	if cards[lastIndex][len(cards[lastIndex])-1] != '\n' {
		cards[lastIndex] += "\n"
	}
	originalCards := cards
	for i := 0; i < len(cards)-1; i++ {
		wins := getWinnings(cards[i])
		originalIndex := getOriginalIndex(originalCards, cards[i])
		if originalIndex == -1 {
			continue
		}
		newCards := createNewCards(originalCards, originalIndex, wins)
		cards = append(cards, newCards...)
	}
	println(len(cards))
}

func createNewCards(cards []string, index int, wins int) []string {
	var newCards []string
	for i := 1; i <= wins; i++ {
		newCards = append(newCards, cards[index+i])
	}
	return newCards
}

func getWinnings(card string) int {
	startOfWinnings := strings.Index(card, ":")
	endOfWinnings := strings.Index(card, "|")
	winningNumbers := strings.Fields(card[startOfWinnings+2 : endOfWinnings])

	startOfNumbers := strings.Index(card, "|")
	endOfNumbers := strings.Index(card, "\n")
	numbers := strings.Fields(card[startOfNumbers+2 : endOfNumbers])

	wins := 0
	for _, winningNumber := range winningNumbers {
		for i := 0; i < len(numbers); i++ {
			if winningNumber == numbers[i] {
				wins++
				numbers = append(numbers[:i], numbers[i+1:]...)
				break
			}
		}
	}
	return wins
}

func getOriginalIndex(originalCards []string, card string) int {
	colonIndex := strings.Index(card, ":")
	if colonIndex != -1 {
		cardNumberStr := strings.TrimSpace(card[:colonIndex])
		for i := 0; i < len(originalCards); i++ {
			if strings.Contains(originalCards[i], cardNumberStr+":") {
				return i
			}
		}
	}
	return -1
}
