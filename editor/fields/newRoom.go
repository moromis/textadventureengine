package fields

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func NewRoom(w fyne.Window, callback func()) *widget.Button {
	button := widget.NewButtonWithIcon("", fyne.CurrentApp().Settings().Theme().Icon(fyne.ThemeIconName("contentAdd")), callback)
	return button
}
