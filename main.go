package main

import (
	"fmt"
)

func main() {
	cityMap := [8][8]string{
		{" ", " ", "X", " ", "D", " ", "X", " "},
		{"X", "D", " ", " ", "X", " ", " ", "D"},
		{" ", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", " ", " ", " ", " ", "X"},
		{" ", " ", " ", "D", " ", "X", " ", " "},
		{" ", "X", " ", " ", " ", " ", "X", " "},
		{"D", " ", " ", " ", " ", " ", " ", " "},
		{" ", " ", " ", "X", " ", " ", "D", " "},
	}
	locations := [5][2]int{
		{1, 64},
		{4, 1},
		{8, 4},
		{3, 8},
		{256, 0},
	}

	var markets []int
	for _, location := range locations {
		coord := getClosestMarket(location, cityMap)
		markets = append(markets, coord)
	}

	fmt.Println("\n======== DashMart ========")
	fmt.Println("The nearest DashMart from each one of the locations are: ", markets)
}

func getClosestMarket(coord [2]int, cityMap [8][8]string) int {
	if coord[0] > len(cityMap[0]) || coord[1] > len(cityMap) {
		return -1
	}

	return 2
}
