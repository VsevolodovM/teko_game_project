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
