package editor

import (
	"log"

	"textadventureengine/gameFileIO"
	"textadventureengine/runner"
	"textadventureengine/structs"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

// GLOBALS
// TODO: move to preferences
var WINDOW_WIDTH float32 = 640
var WINDOW_HEIGHT float32 = 480

func openFile(a fyne.App, callback func(*structs.Game)) {
	w := a.NewWindow("Save Game")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	dialog.ShowFileOpen(func(item fyne.URIReadCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		callback(gameFileIO.ReadGameFileFromJson(item.URI().Path()))
		w.Close()
	}, w)
}

func openFileSave(a fyne.App, callback func(fyne.URI)) {
	w := a.NewWindow("Save Game")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	// dialog.ShowFileSave(func(item fyne.URIReadCloser, err error) {
	// 	if err != nil {
	// 		log.Fatal(err)
	// 	}
	// 	uri := item.URI()
	// 	path := uri.Path()
	// 	// TODO: do something with entities
	// 	var mapLayout, mapWidth, startingRoom, _, inventory = gameFileIO.ReadGameFileFromJson(path)
	// 	stateMachine.SetupStateMachine(mapLayout, mapWidth, startingRoom, inventory)
	// 	callback(uri)
	// 	w.Close()
	// }, w)
}

func OpenEditor(a fyne.App) {
	// setup window
	w := a.NewWindow("TAE Editor") // TODO: replace with title of game?
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))

	var currentGame *structs.Game = nil

	// HEADER
	// buttons
	exit := widget.NewButton("Exit", func() { w.Close() })
	test := widget.NewButton("Test", func() { go runner.OpenRunner(a, currentGame) })
	test.Disable()
	open := widget.NewButton("Open", func() {
		go openFile(a, func(g *structs.Game) {
			currentGame = g
			test.Enable()
		})
	})
	save := widget.NewButton("Save", func() { go openFileSave(a, func(u fyne.URI) {}) })
	save.Disable()

	// set up the contents of the window
	w.SetContent(container.NewVBox(
		container.NewHBox(exit, layout.NewSpacer(), open, test, save),
	))

	// show and run the window
	w.Show()
}
