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
		{7, 4},
		{4, 2},
		{8, 4},
		{3, 7},
		{256, 0},
	}

	var markets []int
	for _, location := range locations {
		coord := getClosestMarket(location, cityMap)
		markets = append(markets, coord)
	}

	fmt.Println("\n======== DashMart ========")
	fmt.Println("The nearest DashMart from each one of the locations are:\n", markets)
}

func getClosestMarket(coord [2]int, cityMap [8][8]string) int {
	if coord[0] >= len(cityMap[0]) || coord[1] >= len(cityMap) {
		return -1
	}

	var directions = [][2]int{
		{-1, 0},
		{1, 0},
		{0, -1},
		{0, 1},
	}
	queue := [][]int{{coord[0], coord[1], 0}}
	visited := make([][]bool, len(cityMap))
	for i := range visited {
		visited[i] = make([]bool, len(cityMap[0]))
	}

	visited[coord[1]][coord[0]] = true

	for len(queue) > 0 {
		node := queue[0]
		queue = queue[1:]

		x, y, dst := node[0], node[1], node[2]
		if cityMap[y][x] == "D" {
			return dst
		}

		for _, dir := range directions {
			newX, newY := x+dir[0], y+dir[1]

			if newX >= 0 && newX < len(cityMap[0]) && newY >= 0 && newY < len(cityMap) && !visited[newY][newX] {
				visited[newY][newX] = true
				queue = append(queue, []int{newX, newY, dst + 1})
			}
		}
	}

	return -1
}
