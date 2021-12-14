package main

import (
	"testing"
)

func TestIsNumberChosen(t *testing.T) {
	s := []int{1}
	str := 1

	chosen := IsNumberChosen(s, str)

	if chosen == false {
		t.Fatal("Chosen number should have matched")
	}

	s = []int{1}
	str = 2
	chosen = IsNumberChosen(s, str)

	if chosen == true {
		t.Fatal("Chosen number should not have matched")
	}

}

func TestSumOfMarkedBalls(t *testing.T) {
	winningCards := [][]int{{1, 1, 2, 3, 4}, {1, 1, 2, 3, 4}}
	chosenBalls := []int{2, 3}

	sum := SumOfUnmarkedBalls(winningCards, chosenBalls)

	if sum != 12 {
		t.Fatal("Sum was incorrect")
	}
}
