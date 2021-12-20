package main

import (
	"log"

	"textadventureengine/gameFileIO"
	"textadventureengine/mapManager"
	"textadventureengine/stateMachine"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// GLOBALS
var WINDOW_WIDTH float32 = 640
var WINDOW_HEIGHT float32 = 480

func getMapTable() *widget.Table {
	var mapInstance = mapManager.GetMap()
	var data = mapInstance.GetMapTable()
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			currentLocation := mapInstance.GetCurrentRoom().Location
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

	// show the window
	w.Show()
}

func openFileSelect(a fyne.App, callback func()) {
	w := a.NewWindow("Open Game (*.tae)")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	dialog.ShowFileOpen(func(item fyne.URIReadCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		uri := item.URI()
		filename := uri.Path()
		// TODO: do something with entities
		var mapLayout, mapWidth, startingRoom, _, inventory = gameFileIO.ReadGameFileFromJson(filename)
		stateMachine.SetupStateMachine(mapLayout, mapWidth, startingRoom, inventory)
		callback()
		w.Close()
	}, w)
}

// MAIN
func main() {
	// setup state machine
	// TODO: move this to after loading game file
	// stateMachine.SetupStateMachine()

	// setup window
	a := app.New()
	w := a.NewWindow("Runner") // TODO: replace with title of game?
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	// TODO: remove, just for testing
	// openMapWindow(a)

	// HEADER
	// buttons
	exit := widget.NewButton("Exit", func() { w.Close() })
	openMap := widget.NewButton("Open Map", func() { go openMapWindow(a) })
	openMap.Disable()
	// title
	title := widget.NewLabel("{Game Title}") // TODO: get title from game files when loaded
	title.TextStyle.Bold = true

	// OUTPUT BUFFER
	// TODO: move this to after setting up state machine
	//       also, this could be a custom action
	var t = ""
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
		response := stateMachine.ParseInput(inputText)

		// add response to output buffer
		t = t + "\n" + response + "\n"
		text.SetText(t)
		textScroll.ScrollToBottom() // make sure we're always scrolled to the bottom (the most recent)
		input.Enable()
	}
	input.SetPlaceHolder("Type Here")
	input.Disable()
	input.OnSubmitted = submitFunc

	// SUBMIT BUTTON
	submit := widget.NewButtonWithIcon("Submit", fyne.CurrentApp().Icon(), func() { submitFunc(input.Text) })
	submit.Disable()

	// OPEN FILE BUTTON
	openFile := widget.NewButton("Open Game File", func() {
		go openFileSelect(a, func() {
			t = mapManager.GetMap().PrintRoom(false)
			text.SetText(t)
			submit.Enable()
			input.Enable()
			openMap.Enable()
		})
	})

	// set up the contents of the window
	w.SetContent(container.NewVBox(
		container.NewHBox(exit, openFile, layout.NewSpacer(), openMap),
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
