package test

import (
	"math/rand"
)

func ChooseRandomPlace(arr []int32) int32 {
	var empty_places []int32
	for i := range arr {
		if arr[i] == 0 {
			empty_places = append(empty_places, int32(i))
		}
	}
	rand_index := rand.Intn(len(empty_places))
	return empty_places[rand_index]
}

func OneDtotwoD(i int32) (int32, int32) {
	y := i / 5
	x := i - 5*y

	return x, y
}
