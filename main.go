package main

import (
	"textadventureengine/editor"
	"textadventureengine/runner"
	sharedUi "textadventureengine/shared/ui"

	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// MAIN
func main() {
	// setup window
	a := app.New()
	w := a.NewWindow("Text Adventure Engine")
	w.CenterOnScreen()

	// HEADER
	// buttons
	openRunner := widget.NewButtonWithIcon("Open Game Runner", theme.MediaPlayIcon(), func() { go runner.OpenRunner(a, nil) })
	openEditor := widget.NewButtonWithIcon("Open Game Editor", theme.MediaRecordIcon(), func() { go editor.OpenEditor(a) })

	// set up the contents of the window
	w.SetContent(container.NewVBox(sharedUi.ExitButton(w), layout.NewSpacer(), layout.NewSpacer(), openRunner, layout.NewSpacer(), openEditor, layout.NewSpacer(), layout.NewSpacer()))

	// show and run the window
	w.ShowAndRun()
}
