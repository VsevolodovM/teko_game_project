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
func TwoDtoOneD(x, y uint32) int32 {
	return int32(y*5 + x)
}

func AvailableNeighborCells(position int, gameState [25]int32) []int {
	neighbors := []int{}

	// cell to the right
	if position%5 < 4 && gameState[position+1] == 0 {
		neighbors = append(neighbors, position+1)
	}
	// cell to the left
	if position%5 > 0 && gameState[position-1] == 0 {
		neighbors = append(neighbors, position-1)
	}
	// cell above
	if position >= 5 && gameState[position-5] == 0 {
		neighbors = append(neighbors, position-5)
	}
	// cell below
	if position < 20 && gameState[position+5] == 0 {
		neighbors = append(neighbors, position+5)
	}

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
