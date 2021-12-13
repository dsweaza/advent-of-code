package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	file, err := os.Open("../input.txt")

	if err != nil {
		fmt.Println(err)
	}

	scanner := bufio.NewScanner(file)
	depths := make(map[int]int)
	i := 0
	measurementIncrease := 0

	for scanner.Scan() {

		// Cast string to int
		depth, _ := strconv.Atoi(scanner.Text())
		depths[i] = depth

		if i != 0 && depth > depths[i-1] {
			measurementIncrease++
		}

		i++
	}

	println("# Increased: ", measurementIncrease)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}
