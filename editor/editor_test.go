package editor

import (
	"testing"

	"fyne.io/fyne/v2/test"

	"github.com/stretchr/testify/assert"
)

func TestEditor(t *testing.T) {
	buttons := OpenEditor(test.NewApp())

	// there should only be one room when we start the editor
	assert.Equal(t, len(buttons), 1)

	// tap the room button
	test.Tap(buttons[0])
}
