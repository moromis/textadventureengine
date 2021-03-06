package main

import (
	"textadventureengine/editor"
	"textadventureengine/runner"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// GLOBALS
// TODO: move to preferences
var WINDOW_WIDTH float32 = 640
var WINDOW_HEIGHT float32 = 480

// MAIN
func main() {
	// setup window
	a := app.New()
	w := a.NewWindow("Text Adventure Engine")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	// HEADER
	// buttons
	exit := widget.NewButton("Exit", func() { w.Close() })
	openRunner := widget.NewButton("Open Game Runner", func() { go runner.OpenRunner(a, nil) })
	openEditor := widget.NewButton("Open Game Editor", func() { go editor.OpenEditor(a) })

	// set up the contents of the window
	w.SetContent(container.NewVBox(container.NewHBox(exit), layout.NewSpacer(), layout.NewSpacer(), openRunner, layout.NewSpacer(), openEditor, layout.NewSpacer(), layout.NewSpacer()))

	// show and run the window
	w.ShowAndRun()
}
