package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	file, err := os.Open("input.txt")
	scanner := bufio.NewScanner(file)

	if err != nil {
		fmt.Println(err)
	}

	//Part1(scanner)

	Part2(scanner)

}

func Part1(scanner *bufio.Scanner) {

	x := 0
	y := 0

	for scanner.Scan() {
		commandParts := strings.Fields(scanner.Text())
		command := commandParts[0]
		units, _ := strconv.Atoi(commandParts[1])

		//fmt.Println("Command:", command, " / Spaces:", units)

		switch command {
		case "forward":
			x = x + units
		case "down":
			y = y + units
		case "up":
			y = y - units
		}
	}

	fmt.Println("X:", x)
	fmt.Println("Y:", y)
	fmt.Println("Answer:", x*y)
}

func Part2(scanner *bufio.Scanner) {

	x := 0
	y := 0
	a := 0

	for scanner.Scan() {
		commandParts := strings.Fields(scanner.Text())
		command := commandParts[0]
		units, _ := strconv.Atoi(commandParts[1])

		//fmt.Println("Command:", command, " / Spaces:", units)

		switch command {
		case "forward":
			x = x + units
			y = y + (a * units)
		case "down":
			//y = y + units
			a = a + units
		case "up":
			//y = y - units
			a = a - units
		}
	}

	fmt.Println("X:", x)
	fmt.Println("Y:", y)
	fmt.Println("A:", a)
	fmt.Println("Answer:", x*y)
}
