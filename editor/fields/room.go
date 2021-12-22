package fields

import (
	"fmt"
	"textadventureengine/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

func Room(w fyne.Window, rooms []*constants.Entity, index int, saveRoomCallback func(*constants.Entity)) *widget.Button {
	buttonTitle := fmt.Sprintf("Room %d", index+1)
	windowTitle := fmt.Sprintf("Create Room %d", index+1)
	room := rooms[index]
	if room != nil {
		if room.Name != "" {
			buttonTitle = room.Name
			windowTitle = fmt.Sprintf("Edit %s", room.Name)
		}
	}
	button := widget.NewButton(buttonTitle, func() {
		OpenEntityEditorWindow(w, room, windowTitle, false, saveRoomCallback)
	})
	return button
}
