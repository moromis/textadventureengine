package helpers

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestIndexInto2dArray(t *testing.T) {
	testArr := []int{1, 2, 3, 4, 5, 6} // 2x3 array
	width := 2

	tests := []struct {
		row    int
		col    int
		result int
	}{
		{0, 0, 1},
		{1, 0, 3},
		{0, 1, 2},
		{2, 1, 6},
	}

	for _, test := range tests {
		result := IndexInto2dArray(testArr, test.col, test.row, width)
		if result != test.result {
			t.Errorf("Indexing into array at (%d, %d) failed, got: %d, want: %d.", test.row, test.col, result, test.result)
		}
	}
}

func TestPickStringRandomly(t *testing.T) {
	test := []string{"a", "b"}
	selection := PickStringRandomly(test)
	assert.Contains(t, test, selection)
}

func TestTitleCase(t *testing.T) {
	test := "not_TITLE case aT AlL"
	output := TitleCase(test)
	assert.Equal(t, "Not Title Case At All", output)
}
