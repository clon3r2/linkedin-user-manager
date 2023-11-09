package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"main/database"
)

type App struct {
	A             fyne.App
	Window        fyne.Window
	MainContainer *fyne.Container
}

func (a *App) initApp() {
	a.A = app.New()
	a.Window = a.A.NewWindow("linkedin users")
	a.Window.Resize(fyne.Size{
		Width:  1024,
		Height: 500,
	})
	addPanel := AddPanel{}
	searchPanel := SearchPanel{}
	fmt.Println("before add panel init")
	addPanel.initialize(&a.Window)
	searchPanel.initialize()

	a.MainContainer = container.NewGridWithColumns(2)
	a.MainContainer.Add(addPanel.panel)
	a.MainContainer.Add(searchPanel.panel)
}

type AddPanel struct {
	panel         *fyne.Container
	NamePanel     *fyne.Container
	ListPanel     *fyne.Container
	addNameButton *widget.Button
	addListButton *widget.Button
	nameEntry     *widget.Entry
	listEntry     *widget.Entry
	typeSelect    *widget.Select
}

func (panel *AddPanel) initialize(win *fyne.Window) {
	panel.panel = container.NewGridWithRows(2)
	panel.namePanelInit(win)
	panel.initListPanel(win)
	panel.panel.Add(panel.ListPanel)
	panel.panel.Add(panel.NamePanel)
}

func (panel *AddPanel) namePanelInit(win *fyne.Window) {
	panel.NamePanel = container.NewVBox()
	var allListTypes []database.ListType
	var allTypeNames []string
	database.DBConn.Find(&allListTypes)
	for i := 0; i < len(allListTypes); i++ {
		allTypeNames = append(allTypeNames, allListTypes[i].TypeName)
	}
	panel.nameEntry = widget.NewEntry()
	panel.typeSelect = widget.NewSelect(allTypeNames, func(s string) {})
	panel.typeSelect.PlaceHolder = "select a list type ..."
	panel.typeSelect.OnChanged = func(s string) {
		if panel.typeSelect.Selected == "" {
			panel.addNameButton.Disable()
		} else if panel.nameEntry.Text != "" {
			panel.addNameButton.Enable()
		}

	}
	panel.nameEntry.PlaceHolder = "new name ..."
	panel.nameEntry.OnChanged = func(input string) {
		if input == "" {
			panel.addNameButton.Disable()
		} else if panel.typeSelect.Selected != "" {
			panel.addNameButton.Enable()
		}
	}
	panel.addNameButton = widget.NewButtonWithIcon("add", theme.ContentAddIcon(), func() {
		var nameType database.ListType
		database.DBConn.Where("type_name = ?", panel.typeSelect.Selected).Find(&nameType)
		database.DBConn.Create(&database.LinkedinUser{
			Name:     panel.nameEntry.Text,
			ListType: nameType,
		})
	})
	panel.addNameButton.Disable()

	panel.NamePanel.Add(panel.typeSelect)
	panel.NamePanel.Add(panel.nameEntry)
	panel.NamePanel.Add(panel.addNameButton)
}

func (panel *AddPanel) initListPanel(win *fyne.Window) {
	panel.ListPanel = container.NewVBox()
	panel.listEntry = widget.NewEntry()
	panel.listEntry.PlaceHolder = "add a new list name ...."
	panel.listEntry.OnChanged = func(input string) {
		if input == "" {
			panel.addListButton.Disable()
		} else {
			panel.addListButton.Enable()
		}
	}
	panel.addListButton = widget.NewButtonWithIcon("add", theme.ContentAddIcon(), func() {
		database.DBConn.Create(&database.ListType{TypeName: panel.listEntry.Text})
		oldListTypes := panel.typeSelect.Options
		panel.typeSelect.SetOptions(append(oldListTypes, panel.listEntry.Text))
		panel.listEntry.SetText("")
		dialog.NewInformation("message", "new list type added successfuly", *win).Show()

	})
	panel.addListButton.Disable()

	panel.ListPanel.Add(panel.listEntry)
	panel.ListPanel.Add(panel.addListButton)
}

type SearchPanel struct {
	panel        *fyne.Container
	searchEntry  *widget.Entry
	searchButton *widget.Button
	resultBox    *widget.Label
}

/*
result = &{Config:0xc0000f2750 Error:<nil> RowsAffected:1 Statement:0xc00171a000 clone:0}
user =  {Model:{ID:1 CreatedAt:2023-11-09 15:07:53.31087526 +0330 +0330 UpdatedAt:2023-11-09 15:07:53.31087526 +0330 +0330 DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:qsd ListTypeID:1 ListType:{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} TypeName:}}

result = &{Config:0xc000124630 Error:<nil> RowsAffected:0 Statement:0xc000213a40 clone:0}
user =  {Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC                       DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} Name:    ListTypeID:0 ListType:{Model:{ID:0 CreatedAt:0001-01-01 00:00:00 +0000 UTC UpdatedAt:0001-01-01 00:00:00 +0000 UTC DeletedAt:{Time:0001-01-01 00:00:00 +0000 UTC Valid:false}} TypeName:}}
*/
func (panel *SearchPanel) initialize() {
	panel.panel = container.NewVBox()
	panel.resultBox = widget.NewLabelWithStyle("salam", fyne.TextAlignCenter, fyne.TextStyle{Bold: true})
	panel.resultBox.Importance = widget.HighImportance
	panel.resultBox.TextStyle = fyne.TextStyle{
		Bold: true,
	}

	panel.searchEntry = widget.NewEntry()
	panel.searchEntry.PlaceHolder = "search name ..."
	panel.searchEntry.OnChanged = func(input string) {
		if input == "" {
			panel.searchButton.Disable()
		} else {
			panel.searchButton.Enable()
		}
	}
	panel.searchButton = widget.NewButtonWithIcon("search", theme.SearchIcon(), func() {
		var user database.LinkedinUser
		result := database.DBConn.
			Preload("ListType").
			Where("name = ?", panel.searchEntry.Text).
			Find(&user)
		if result.RowsAffected == 0 {
			panel.resultBox.SetText("did not found name")
		} else {
			panel.resultBox.SetText(fmt.Sprintf("\n\nfound name: %v \nlist type name = %v", panel.searchEntry.Text, user.ListType.TypeName))
		}

	})
	panel.searchButton.Disable()
	panel.panel.Add(panel.searchEntry)
	panel.panel.Add(panel.searchButton)
	panel.panel.Add(panel.resultBox)

}
