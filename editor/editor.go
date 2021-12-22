package editor

import (
	"log"

	"textadventureengine/constants"
	"textadventureengine/editor/fields"
	"textadventureengine/gameFileIO"
	"textadventureengine/runner"
	sharedUi "textadventureengine/shared/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/imdario/mergo"
)

// TODO: expand on this to enable testing
type Editor struct {
	RoomButtons []*widget.Button
}

var worldWidth = 1
var worldLayout = []*constants.Entity{
	{
		Name:     "Room 1",
		Location: [2]int{10, 10},
	},
}
var startingRoom = worldLayout[0]
var roomButtons []*widget.Button = nil

func saveRoom(room *constants.Entity, index int) {
	if len(worldLayout) > index {
		err := mergo.Merge(worldLayout[index], room, mergo.WithOverride)
		if err != nil {
			log.Fatal(err)
		}
		roomButtons[index].SetText(room.Name)
	}
}

func renderRoomsBorders(w fyne.Window) *fyne.Container {
	line := container.NewVBox(layout.NewSpacer())
	for i := 0; i <= len(worldLayout)/worldWidth+1; i++ {
		line.Add(layout.NewSpacer())
		line.Add(fields.NewRoom(w, func() {})) // TODO: dry
		line.Add(layout.NewSpacer())
	}
	line.Add(layout.NewSpacer())
	return line
}

func renderRooms(w fyne.Window) *fyne.Container {
	content := container.NewHBox(layout.NewSpacer())

	content.Add(renderRoomsBorders(w))
	rooms := container.NewVBox()
	rooms.Add(fields.NewRoom(w, func() {})) // TODO: dry
	for i := 0; i < len(worldLayout); i += worldWidth {
		line := container.NewHBox(layout.NewSpacer())
		for j := 0; j < worldWidth; j++ {
			room := fields.Room(w, worldLayout, i, func(e *constants.Entity) {
				saveRoom(e, i)
			})
			roomButtons = append(roomButtons, room)
			line.Add(layout.NewSpacer())
			line.Add(room)
			line.Add(layout.NewSpacer())
		}
		line.Add(layout.NewSpacer())
		rooms.Add(line)
	}
	rooms.Add(fields.NewRoom(w, func() {})) // TODO: dry
	content.Add(rooms)
	content.Add(renderRoomsBorders(w))
	content.Add(layout.NewSpacer())
	return content
}

func openGame(w fyne.Window, callback func(*constants.Game)) {
	dialog.ShowFileOpen(func(item fyne.URIReadCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if item != nil {
			callback(gameFileIO.ReadGameFileFromJson(item.URI().Path()))
		}
	}, w)
}

func saveGame(w fyne.Window, gameTitle string, callback func(*constants.Game)) {
	dialog.ShowFileSave(func(uc fyne.URIWriteCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if uc != nil {
			game := &constants.Game{
				FilePath:    uc.URI().Path(),
				Title:       gameTitle,
				WorldLayout: worldLayout,
				// TODO: this should be based on data from a select
				// which autofills/allows for selections
				// from rooms we've created
				StartingRoom: startingRoom,
				WorldWidth:   worldWidth,
			}
			gameFileIO.WriteGameFileToJson(game)
			callback(game)
		}
	}, w)
}

func OpenEditor(a fyne.App) []*widget.Button {
	// setup window
	w := a.NewWindow("TAE Editor")
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT))
	w.CenterOnScreen()

	var currentGame *constants.Game = nil

	// edit fields
	form := widget.NewForm()

	gameTitle := fields.Title()
	form.Append("Game Title", gameTitle)

	var roomOptions []string
	for _, room := range worldLayout {
		roomOptions = append(roomOptions, room.Name)
	}
	startingRoom := widget.NewSelect(roomOptions, func(s string) {
		for i := range worldLayout {
			if worldLayout[i].Name == s {
				startingRoom = worldLayout[i]
			}
		}
	})
	startingRoom.SetSelectedIndex(0)
	form.Append("Starting Room", startingRoom)

	// buttons
	test := widget.NewButtonWithIcon("Test", theme.MailSendIcon(), func() {
		if currentGame != nil {
			go runner.OpenRunner(a, currentGame)
		}
	})
	test.Disable()
	open := widget.NewButtonWithIcon("Open", theme.FolderOpenIcon(), func() {
		go openGame(w, func(g *constants.Game) {
			currentGame = g
			gameTitle.SetText(currentGame.Title)
			test.Enable()
		})
	})
	var saveCallback = func(game *constants.Game) {
		currentGame = game
		test.Enable()
	}
	save := widget.NewButtonWithIcon("Save", theme.DocumentSaveIcon(), func() { go saveGame(w, gameTitle.Text, saveCallback) })

	content := container.NewVBox(
		container.NewHBox(sharedUi.ExitButton(w), layout.NewSpacer(), open, test, save),
		form,
		layout.NewSpacer(),
		// TODO: this should map over all the rooms
		// that we've created to display them all --
		// also we need to pass in data based on index
		// so it can populate from saved data
		renderRooms(w),
		layout.NewSpacer(),
	)

	// set up the contents of the window
	w.SetContent(content)

	// show and run the window
	w.Show()

	return roomButtons
}
