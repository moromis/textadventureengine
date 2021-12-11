package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// GLOBALS
var WINDOW_WIDTH float32 = 640
var WINDOW_HEIGHT float32 = 480

func getMapTable() *widget.Table {
	var data = mapInstance.getMapTable()
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			currentLocation := mapInstance.getCurrentRoom().location
			newLabel := data[i.Row][i.Col]
			if i.Row == currentLocation[0] && i.Col == currentLocation[1] {
				newLabel += " *"
			}
			o.(*widget.Label).SetText(newLabel)
		})
	return table
}

func openMapWindow(a fyne.App) {
	// setup window
	w := a.NewWindow("Map")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	// map
	mapWidget := getMapTable()

	// set up the contents of the window
	w.SetContent(mapWidget)
	mapWidget.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	// show and run the window
	w.Show()
}

// MAIN
func main() {
	// setup state machine
	setupStateMachine()

	// setup window
	a := app.New()
	w := a.NewWindow("Runner") // TODO: replace with title of game?
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	openMapWindow(a)

	// HEADER
	// buttons
	exit := widget.NewButton("Exit", func() { w.Close() })
	openMap := widget.NewButton("Open Map", func() { go openMapWindow(a) })
	// title
	title := widget.NewLabel("{Game Title}") // TODO: get title from game files when loaded
	title.TextStyle.Bold = true

	// OUTPUT BUFFER
	var t = mapInstance.printRoom(false)
	text := widget.NewTextGrid()
	text.SetText(t)
	textScroll := container.NewVScroll(
		text,
	)
	textScroll.SetMinSize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT*2/3))

	// INPUT FIELD
	input := widget.NewEntry()
	submitFunc := func(inputText string) {
		t = t + "\n" + "> " + inputText // add the command to the output buffer
		text.SetText(t)                 // update the text
		input.Text = ""                 // clear the input
		input.Refresh()                 // refresh the input so it clears
		textScroll.ScrollToBottom()     // make sure we're always scrolled to the bottom (the most recent)
		input.Disable()                 // disable the input till we have a response

		// send response
		response := parseInput(inputText)

		// add response to output buffer
		t = t + "\n" + response + "\n"
		text.SetText(t)
		textScroll.ScrollToBottom() // make sure we're always scrolled to the bottom (the most recent)
		input.Enable()
	}
	input.SetPlaceHolder("Type Here")
	input.OnSubmitted = submitFunc

	// SUBMIT BUTTON
	submit := widget.NewButtonWithIcon("Submit", fyne.CurrentApp().Icon(), func() { submitFunc(input.Text) })

	// set up the contents of the window
	w.SetContent(container.NewVBox(
		container.NewHBox(exit, layout.NewSpacer(), openMap),
		title,
		textScroll,
		// layout.NewSpacer(),
		input,
		submit,
	))

	// focus the input field
	w.Canvas().Focus(input)

	// show and run the window
	w.ShowAndRun()
}
