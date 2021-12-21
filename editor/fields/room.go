package fields

import (
	"fmt"
	"textadventureengine/runner/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func openCreateStuff(a fyne.App) {
	w := a.NewWindow("Create/Edit Entity")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT))

	// window content
	exit := widget.NewButton("Exit", func() { w.Close() })
	label := widget.NewLabel("Create/Edit Entity")
	entityType := widget.NewSelect([]string{"Animate", "Inanimate"}, func(s string) { println("yeah") })
	name := widget.NewEntry()
	name.SetPlaceHolder("Entity name")
	desc := widget.NewEntry()
	desc.SetPlaceHolder("Entity description")
	createStuff := widget.NewButton("Add Entity", func() { openCreateStuff(a) })

	w.SetContent(container.NewVBox(exit, label, layout.NewSpacer(), entityType, name, desc, createStuff, layout.NewSpacer()))
}

var _name string
var _desc string

// maybe could just use dialog.Custom{...}
func openRoomEdit(a fyne.App) {
	w := a.NewWindow("Create/Edit Room")
	// show the window
	w.Show()
	w.SetFixedSize(true)
	w.Resize(fyne.NewSize(constants.WINDOW_WIDTH, constants.WINDOW_HEIGHT))

	// window content
	label := widget.NewLabel("Create/Edit Room")
	name := widget.NewEntry()
	name.SetPlaceHolder("Room name")
	desc := widget.NewEntry()
	desc.SetPlaceHolder("Room description")
	createStuff := widget.NewButton("Add Entity", func() { openCreateStuff(a) })
	cancel := widget.NewButton("Cancel", func() { w.Close() })
	save := widget.NewButton("Save", func() {
		println(name.Text, desc.Text)
		w.Close()
	})

	w.SetContent(container.NewVBox(label, layout.NewSpacer(), name, desc, createStuff, layout.NewSpacer(), container.NewHBox(layout.NewSpacer(), cancel, save)))
}

func Room(a fyne.App, index int) *widget.Button {
	button := widget.NewButton(fmt.Sprintf("Room %d", index), func() { openRoomEdit(a) })
	return button
}
