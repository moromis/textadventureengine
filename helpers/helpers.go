package helpers

import (
	"math/rand"
	"strings"
)

func IndexInto2dArray(arr []int, col int, row int, width int) int {
	return arr[col+row*width]
}

func PickStringRandomly(arr []string) string {
	randomIndex := rand.Intn(len(arr))
	return arr[randomIndex]
}

// FROM: https://www.jeremymorgan.com/tutorials/go/learn-golang-casing/
func TitleCase(input string) string {
	input = strings.ReplaceAll(input, "_", " ")
	return strings.Title(strings.ToLower(input))
}
