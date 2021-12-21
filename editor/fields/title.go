package fields

import (
	"fyne.io/fyne/v2/widget"
)

var Title = func() *widget.Entry {
	input := widget.NewEntry()
	input.SetPlaceHolder("What's the title of your game?")
	return input
}
