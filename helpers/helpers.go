package helpers

import "math/rand"

func IndexInto2dArray(arr []int, col int, row int, width int) int {
	return arr[col+row*width]
}

func PickStringRandomly(arr []string) string {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}
