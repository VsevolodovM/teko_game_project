package logic

import (
	"math/rand"
)

func ChooseRandomPlace(arr []int32, num int32) int32 {
	var empty_places []int32
	for i := range arr {
		if arr[i] == num {
			empty_places = append(empty_places, int32(i))
		}
	}
	rand_index := rand.Intn(len(empty_places))
	return empty_places[rand_index]
}

func OneDtotwoD(i int32) (uint32, uint32) {
	y := i / 5
	x := i - 5*y

	return uint32(x), uint32(y)
}

func AvailableNeighborCells(position int, gameState []int32) []int {
	neighbors := []int{}

	// Check if the cell to the right is available
	if position%5 < 4 && gameState[position+1] == 0 {
		neighbors = append(neighbors, position+1)
	}
	// Check if the cell to the left is available
	if position%5 > 0 && gameState[position-1] == 0 {
		neighbors = append(neighbors, position-1)
	}
	// Check if the cell above is available
	if position >= 5 && gameState[position-5] == 0 {
		neighbors = append(neighbors, position-5)
	}
	// Check if the cell below is available
	if position < 20 && gameState[position+5] == 0 {
		neighbors = append(neighbors, position+5)
	}

	// Check diagonal neighbors
	if position%5 < 4 && position >= 5 && gameState[position-4] == 0 {
		neighbors = append(neighbors, position-4) // Upper right diagonal
	}
	if position%5 > 0 && position >= 5 && gameState[position-6] == 0 {
		neighbors = append(neighbors, position-6) // Upper left diagonal
	}
	if position%5 < 4 && position < 20 && gameState[position+6] == 0 {
		neighbors = append(neighbors, position+6) // Lower right diagonal
	}
	if position%5 > 0 && position < 20 && gameState[position+4] == 0 {
		neighbors = append(neighbors, position+4) // Lower left diagonal
	}

	return neighbors
}

func isElementZero(arr []int, index int) bool {
	return arr[index] == 0
}
