package editor

import (
	"log"

	"textadventureengine/editor/fields"
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

func openGame(a fyne.App, callback func(*structs.Game)) {
	w := a.NewWindow("Save Game")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	dialog.ShowFileOpen(func(item fyne.URIReadCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if item != nil {
			callback(gameFileIO.ReadGameFileFromJson(item.URI().Path()))
		}
		w.Close()
	}, w)
}

func saveGame(a fyne.App, gameTitle string, callback func(*structs.Game)) {
	w := a.NewWindow("Save Game")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(WINDOW_WIDTH, WINDOW_HEIGHT))
	dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if uc != nil {
			game := &structs.Game{
				FilePath: uc.URI().Path(),
				Title:    gameTitle,
			}
			gameFileIO.WriteGameFileToJson(game)
			callback(game)
		}
		w.Close()
	}, w)
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
		go openGame(a, func(g *structs.Game) {
			currentGame = g
			test.Enable()
		})
	})

	// edit fields
	gameTitle := fields.Title()

	var saveCallback = func(game *structs.Game) {
		currentGame = game
		test.Enable()
	}
	// save button
	save := widget.NewButton("Save", func() { go saveGame(a, gameTitle.Text, saveCallback) })
	// save.Disable()

	// set up the contents of the window
	w.SetContent(container.NewVBox(
		container.NewHBox(exit, layout.NewSpacer(), open, test),
		gameTitle,
		layout.NewSpacer(),
		container.NewHBox(layout.NewSpacer(), save),
	))

	// show and run the window
	w.Show()
}
