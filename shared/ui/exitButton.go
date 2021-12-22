package sharedUi

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func ExitButton(w fyne.Window) *widget.Button {
	return widget.NewButtonWithIcon("Exit", theme.CancelIcon(), func() { w.Close() })
}
