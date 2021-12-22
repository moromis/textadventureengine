package fields

import (
	"textadventureengine/constants"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

func OpenEntityEditorWindow(w fyne.Window, entityDetails *constants.Entity, title string, showEntityTypeSelect bool, saveEntityCallback func(*constants.Entity)) {
	// window content
	label := widget.NewLabel(title)
	// get entity type select options
	entityTypeOptions := make([]string, len(constants.ENTITY_TYPES), len(constants.ENTITY_TYPES))
	for idx, value := range constants.ENTITY_TYPES {
		entityTypeOptions[idx] = value
	}
	entityType := widget.NewSelect(entityTypeOptions, func(s string) { println("yeah") })
	// entityType.SetSelectedIndex()
	if !showEntityTypeSelect {
		entityType.Hide()
	}
	name := widget.NewEntry()
	name.SetPlaceHolder("Name")
	desc := widget.NewMultiLineEntry()
	desc.SetPlaceHolder("Description")
	if entityDetails != nil {
		name.SetText(entityDetails.Name)
		desc.SetText(entityDetails.Desc)
	}
	createStuff := widget.NewButton("Add Entity", func() {
		OpenEntityEditorWindow(w, entityDetails, "Create/Edit Entity", true, func(e *constants.Entity) {})
	})

	content := container.NewVBox(label, layout.NewSpacer(), entityType, name, desc, createStuff, layout.NewSpacer())
	dialog.ShowCustomConfirm(title, "Save", "Cancel", content, func(bool) {
		saveEntityCallback(&constants.Entity{
			Name: name.Text,
			Desc: desc.Text,
		})
	}, w)
}
