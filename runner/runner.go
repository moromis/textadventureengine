package runner

import (
	"log"
	"strings"

	. "textadventureengine/constants"
	"textadventureengine/gameFileIO"
	"textadventureengine/helpers"
	"textadventureengine/runner/stateMachine"
	"textadventureengine/runner/worldManager"
	sharedUi "textadventureengine/shared/ui"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

func getMapTable() *widget.Table {
	var worldInstance = worldManager.GetWorldManager()
	var data = worldInstance.GetWorldTable()
	table := widget.NewTable(
		func() (int, int) {
			return len(data), len(data[0])
		},
		func() fyne.CanvasObject {
			return widget.NewLabel("wide content")
		},
		func(i widget.TableCellID, o fyne.CanvasObject) {
			currentLocation := worldInstance.GetCurrentRoom().Location
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

func openFileSelect(w fyne.Window, callback func(fyne.URI)) {
	dialog.ShowFileOpen(func(item fyne.URIReadCloser, err error) {
		if err != nil {
			log.Fatal(err)
		}
		if item != nil {
			uri := item.URI()
			path := uri.Path()
			stateMachine.SetupStateMachine(gameFileIO.ReadGameFileFromJson(path))
			callback(uri)
		}
		w.Close()
	}, w)
}

func OpenRunner(a fyne.App, game *Game) {
	// setup window
	w := a.NewWindow("TAE Runner")
	w.CenterOnScreen()

	// HEADER
	// buttons
	openMap := widget.NewButtonWithIcon("Open Map", theme.ComputerIcon(), func() { go openMapWindow(a) })
	openMap.Disable()
	// title
	title := widget.NewLabel("")
	title.TextStyle.Bold = true

	// OUTPUT BUFFER
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
	submit := widget.NewButtonWithIcon("Submit", theme.MailSendIcon(), func() { submitFunc(input.Text) })
	submit.Disable()

	// open game callback
	var openGameCallback = func(newTitleText string) {
		title.SetText(helpers.TitleCase(newTitleText))
		t = worldManager.GetWorldManager().PrintRoom(false)
		text.SetText(t)
		submit.Enable()
		input.Enable()
		openMap.Enable()
	}

	// OPEN FILE BUTTON
	openFile := widget.NewButtonWithIcon("Open Game File", theme.FolderOpenIcon(), func() {
		go openFileSelect(w, func(uri fyne.URI) {
			openGameCallback(helpers.TitleCase(strings.ReplaceAll(uri.Name(), uri.Extension(), "")))
		})
	})

	// setup game if we've been given one
	if game != nil {
		stateMachine.SetupStateMachine(game)
		openGameCallback(game.Title)
	}

	// set up the contents of the window
	w.SetContent(container.NewVBox(
		container.NewHBox(sharedUi.ExitButton(w), openFile, layout.NewSpacer(), openMap),
		title,
		textScroll,
		// layout.NewSpacer(),
		input,
		submit,
	))

	// focus the input field
	w.Canvas().Focus(input)

	// show and run the window
	w.Show()
}
