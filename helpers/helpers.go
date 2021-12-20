package helpers

func indexInto2dArray(arr []int, col int, row int, width int) int {
	return arr[col+row*width]
}
