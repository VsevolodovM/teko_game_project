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

	if position%5 < 4 && gameState[position+1] == 0 {
		neighbors = append(neighbors, position+1)
	}
	if position%5 > 0 && gameState[position-1] == 0 {
		neighbors = append(neighbors, position-1)
	}
	if position >= 5 && gameState[position-5] == 0 {
		neighbors = append(neighbors, position-5)
	}
	if position < 20 && gameState[position+5] == 0 {
		neighbors = append(neighbors, position+5)
	}

	return neighbors
}

func GetWinningPosition(gameState []int32) int {
	for i := 0; i < len(gameState); i += 5 {
		row := gameState[i : i+5]
		if checkRowForWin(row) {
			for j := 0; j < len(row); j++ {
				if row[j] == 0 {
					return i + j
				}
			}
		}
	}

	return -1
}

func checkRowForWin(row []int32) bool {
	if row[0] == row[1] && row[1] == row[2] && row[0] != 0 {
		return true
	} else if row[1] == row[2] && row[2] == row[3] && row[1] != 0 {
		return true
	} else if row[2] == row[3] && row[3] == row[4] && row[2] != 0 {
		return true
	}
	return false
}
