package main

import (
	"textadventureengine/constants"
	"textadventureengine/editor"
	"textadventureengine/runner"
	sharedUi "textadventureengine/shared/ui"

	"fyne.io/fyne/v2"
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
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT))
	w.CenterOnScreen()

	// HEADER
	// buttons
	title := widget.NewLabel("Text Adventure Engine")
	openRunner := widget.NewButtonWithIcon("Open Game Runner", theme.MediaPlayIcon(), func() { go runner.OpenRunner(a, nil) })
	openEditor := widget.NewButtonWithIcon("Open Game Editor", theme.MediaRecordIcon(), func() { go editor.OpenEditor(a) })

	// set up the contents of the window
	w.SetContent(
		container.NewVBox(
			container.NewHBox(sharedUi.ExitButton(w)),
			container.NewCenter(title),
			layout.NewSpacer(),
			container.NewCenter(openRunner),
			container.NewCenter(openEditor),
			layout.NewSpacer(),
			layout.NewSpacer(),
		),
	)

	// show and run the window
	w.ShowAndRun()
}
