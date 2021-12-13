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
	sums := []int{}

	i := 0
	measurementIncrease := 0

	for scanner.Scan() {
		i++

		// Cast string to int
		depth, _ := strconv.Atoi(scanner.Text())
		depths[i] = depth

		if i < 3 {
			continue
		}

		sum := depths[i-2] + depths[i-1] + depths[i]
		sums = append(sums, sum)
		//s := fmt.Sprintf("%d + %d + %d", depths[i-2], depths[i-1], depths[i])
		//fmt.Println(sum, " :: ", s)

	}

	previousSum := 0
	for _, sum := range sums {

		if previousSum != 0 && sum > previousSum {
			measurementIncrease++
		}

		previousSum = sum
	}

	fmt.Println("Measurement increase: ", measurementIncrease)

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}

}
