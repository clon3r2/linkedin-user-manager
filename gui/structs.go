package gui

import (
	"fmt"
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
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
	addPanel.initialize()
	searchPanel.initialize()

	a.MainContainer = container.NewGridWithColumns(2)
	a.MainContainer.Add(addPanel.panel)
	a.MainContainer.Add(searchPanel.panel)
}

type AddPanel struct {
	panel         *fyne.Container
	addNameButton *widget.Button
	addListButton *widget.Button
	nameEntry     *widget.Entry
}

func (panel *AddPanel) initialize() {
	panel.panel = container.NewVBox()
	panel.nameEntry = widget.NewEntry()
	panel.nameEntry.PlaceHolder = "new name ..."
	panel.nameEntry.OnChanged = func(input string) {
		if input == "" {
			panel.addNameButton.Disable()
		} else {
			panel.addNameButton.Enable()
		}
	}
	fmt.Println("before button init\n\n")
	panel.addNameButton = widget.NewButtonWithIcon("add", theme.ContentAddIcon(), func() {
		var nameType database.ListType
		database.DBConn.First(&nameType)
		database.DBConn.Create(&database.LinkedinUser{
			Name:     panel.nameEntry.Text,
			ListType: nameType,
		})
		fmt.Println("len(panel.nameEntry.Text) =", len(panel.nameEntry.Text))
		fmt.Println("panel.nameEntry.Text = ", panel.nameEntry.Text)

	})
	panel.addNameButton.Disable()
	fmt.Println("after button init\n\n")
	panel.panel.Add(panel.nameEntry)
	fmt.Println("after panel nameEntry add\n\n")
	panel.panel.Add(panel.addNameButton)
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
	panel.resultBox = widget.NewLabel("")
	panel.searchEntry = widget.NewEntry()
	panel.searchEntry.PlaceHolder = "search name ..."
	panel.searchEntry.OnChanged = func(input string) {
		if input == "" {
			panel.searchButton.Disable()
		} else {
			panel.searchButton.Enable()
		}
	}
	fmt.Println("before button init\n\n")
	panel.searchButton = widget.NewButtonWithIcon("search", theme.SearchIcon(), func() {
		var user database.LinkedinUser
		result := database.DBConn.Where("name = ?", panel.searchEntry.Text).Find(&user)
		//result := database.DBConn.Find(&user)
		if result.RowsAffected == 0 {
			fmt.Println("\n\ndidnt find name!")
			panel.resultBox.SetText(fmt.Sprintf("\n\ndid not found name: %v", panel.resultBox.Text))

		} else {
			fmt.Println()
			panel.resultBox.SetText(fmt.Sprintf("\n\nfound name: %v \n list type name = %v", panel.resultBox.Text, user.ListType.TypeName))
		}

	})
	panel.searchButton.Disable()
	panel.panel.Add(panel.searchEntry)
	panel.panel.Add(panel.searchButton)
	panel.panel.Add(panel.resultBox)

}
