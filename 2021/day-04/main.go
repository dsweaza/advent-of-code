package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	//Part1()
	Part2()
}

func Part1() {
	balls, cards := setVars()

	winnerBool, winningCardint, lastChosenBall, chosenBalls, winningCard := PlayGame(true, balls, cards)

	if winnerBool {
		fmt.Println("Winning Card", winningCardint)
		fmt.Println("Last Chosen Ball", lastChosenBall)
		fmt.Println("Total Chosen Balls", len(chosenBalls))
		sum := SumOfUnmarkedBalls(winningCard, chosenBalls)
		fmt.Println("Sum", sum)

		answer := sum * lastChosenBall
		fmt.Println("Part 1 Answer: ", answer)
	}

}

func Part2() {
	balls, cards := setVars()

	winnerBool, _, lastChosenBall, chosenBalls, winningCard := PlayGame(false, balls, cards)

	if winnerBool {
		fmt.Println("Last Winning Card", winningCard)

		fmt.Println("Last Chosen Ball", lastChosenBall)
		fmt.Println("Total Chosen Balls", len(chosenBalls))
		fmt.Println("Cards Left", len(cards))
		sum := SumOfUnmarkedBalls(winningCard, chosenBalls)
		fmt.Println("Sum", sum)

		answer := sum * lastChosenBall
		fmt.Println("Part 2 Answer: ", answer)
	}

}

/*
Return : 	bool Winner
			Int Winning Card
			Int Most Recent Number
			Chosen Balls
*/
func PlayGame(normalMode bool, balls []int, cards map[int][][]int) (bool, int, int, []int, [][]int) {
	var chosenBalls []int
	lastBall := 0
	lastBalls := []int{}
	winningCards := []int{}

	// Last Winners
	lastWinningCardInt := 0
	lastWinningCard := [][]int{}
	lastWinningBall := 0
	lastWinningBalls := []int{}

	for i := 0; i < len(balls); i++ {
		chosenBalls = append(chosenBalls, balls[i])
		currentBall := balls[i]

		winnerBool, winningCardInt, winningCard := CheckWinnersExist(cards, chosenBalls)
		if winnerBool == true {
			winningCards = append(winningCards, winningCardInt)
			lastBalls = chosenBalls
			lastBall = currentBall
		}
		if winnerBool && normalMode == true {
			return winnerBool, winningCardInt, lastBall, lastBalls, winningCard
		} else if winnerBool && normalMode == false {
			lastWinningCardInt = winningCardInt
			lastWinningCard = winningCard
			lastWinningBalls = chosenBalls
			lastWinningBall = currentBall
			fmt.Println("Last removed *******", winningCardInt)
			delete(cards, winningCardInt)
		} else {
			//return true, lastCardInt, lastBall, lastBalls, lastCard
		}

	}

	//fmt.Println(lastBalls)

	if normalMode == false {
		return true, lastWinningCardInt, lastWinningBall, lastWinningBalls, lastWinningCard
	}

	return false, 0, 0, []int{}, [][]int{}

}

/*
Return : 	bool Winner
			Int Winning Card
			Card
*/
func CheckWinnersExist(cards map[int][][]int, chosenBalls []int) (bool, int, [][]int) {

	for i := 0; i < 100; i++ {
		card, ok := cards[i]
		if ok {
			for _, line := range cards[i] {
				lineMatches := 0
				for _, number := range line {
					if IsNumberChosen(chosenBalls, number) == true {
						lineMatches++
					}
				}
				if lineMatches == 5 {
					return true, i, card
				}
			}
		}
	}
	//for cardNum, card := range cards {

	//}

	return false, 0, [][]int{}
}

func SumOfUnmarkedBalls(winningCard [][]int, chosenBalls []int) int {
	sumUnmarkedBalls := 0

	for i := 0; i < 5; i++ {
		line := winningCard[i]

		for _, number := range line {
			if IsNumberChosen(chosenBalls, number) == false {
				sumUnmarkedBalls += number
			}
		}
	}

	return sumUnmarkedBalls
}

func setVars() ([]int, map[int][][]int) {
	var cards = make(map[int][][]int)
	var balls []int
	fileLines, _ := readLines("input.txt")
	cardnum := 0

	// Line 1 is Balls, comma separated
	for _, ball := range strings.Split(fileLines[0], ",") {
		ballint, _ := strconv.Atoi(ball)
		balls = append(balls, ballint)
	}

	// Set Bingo Lines
	// For each line...
	//var currentCard [][]int
	cardLineNum := 1
	var cardColumns = make([][]int, 5)
	for l := 1; l < len(fileLines); l++ {

		currentLine := fileLines[l]
		var cardLine []int

		if currentLine == "" {
			cardnum++
			cardLineNum = 1
			continue
		}

		// Go through each number in the line of the card
		for lni, lineNumber := range strings.Fields(currentLine) {
			lineNumberint, _ := strconv.Atoi(lineNumber)
			cardLine = append(cardLine, lineNumberint)
			cardColumns[lni] = append(cardColumns[lni], lineNumberint)
			//fmt.Println(cardColumns[lni])
		}

		cards[cardnum] = append(cards[cardnum], cardLine)

		// On the 5th, add cardColumns
		if cardLineNum == 5 {
			cards[cardnum] = append(cards[cardnum], cardColumns...)
			cardColumns = nil
			cardColumns = make([][]int, 5)
		}

		cardLineNum++
	}

	return balls, cards

}

func readLines(path string) ([]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines []string
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
		i++
	}

	return lines, scanner.Err()
}

func IsNumberChosen(s []int, str int) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}
	return false
}
