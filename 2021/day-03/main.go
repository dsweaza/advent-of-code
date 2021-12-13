package main

import (
	"bufio"
	"fmt"
	"strconv"

	"os"
)

func main() {

	Part2()

}

func Part1() {

	gamma := ""
	epsilon := ""
	bitsCounter := [][]int{
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
		{0, 0},
	}

	lines, _ := readLines("input.txt")
	for _, line := range lines {

		for i := 0; i < len(line); i++ {
			if line[i] == 48 {
				//Bit = 0
				bitsCounter[i][0] = bitsCounter[i][0] + 1
			} else {
				// Bit = 1
				bitsCounter[i][1] = bitsCounter[i][1] + 1
			}

		}

	}

	for _, v := range bitsCounter {
		if v[0] > v[1] {
			// 0 Significant
			gamma = gamma + "0"
			epsilon = epsilon + "1"
		} else {
			// 1 Significant
			gamma = gamma + "1"
			epsilon = epsilon + "0"
		}

	}

	gammai, _ := strconv.ParseInt(gamma, 2, 64)
	epsiloni, _ := strconv.ParseInt(epsilon, 2, 64)

	fmt.Println("Gamma:", gammai)
	fmt.Println("Epsilon:", epsiloni)
	fmt.Println("Power Consumption:", gammai*epsiloni)
}

func Part2() {

	o2, o2n := Part2LifeSupportCriteria("oxygen")
	co2, co2n := Part2LifeSupportCriteria("c02")
	fmt.Println("Oxygen:", o2, o2n)
	fmt.Println("CO2:", co2, co2n)
	fmt.Println("Life Support Rating:", o2n*co2n)
}

func Part2LifeSupportCriteria(supporttype string) (string, int64) {
	lines, _ := readLines("input.txt")
	positions := 12

	for i := 0; i < positions; i++ {

		if len(lines) == 1 {
			continue
		}

		zeros := 0
		ones := 0

		for _, line := range lines {
			if line[i] == 48 {
				zeros++
			} else {
				ones++
			}
		}

		var remove byte
		remove = 0

		if zeros > ones {
			// remove lines where i position = zero
			if supporttype == "oxygen" {
				remove = 48
			} else {
				remove = 49
			}
		} else {
			// remove lines where i position = one
			if supporttype == "oxygen" {
				remove = 49
			} else {
				remove = 48
			}
		}

		for t, line := range lines {
			if line[i] == remove {
				delete(lines, t)
			}
		}

	}

	value := ""
	for _, line := range lines {
		value = line
	}

	number, _ := strconv.ParseInt(value, 2, 64)

	return value, number
}

func readLines(path string) (map[int]string, error) {
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	var lines = make(map[int]string)
	scanner := bufio.NewScanner(file)
	i := 1
	for scanner.Scan() {
		lines[i] = scanner.Text()
		i++
	}

	return lines, scanner.Err()
}
