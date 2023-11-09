package main

import (
	"fmt"
	"main/database"
	"main/gui"
)

func init() {
	database.InitializeDatabase()
	database.Migrate()

	gui.InitializeGUI()
}

func main() {
	fmt.Printf("app window = %v", gui.MainApp)

	gui.MainApp.Window.SetContent(gui.MainApp.MainContainer)
	gui.MainApp.Window.Show()
	gui.MainApp.A.Run()
}

// TODO: reset all entry boxes after taping on any button
// TODO: validation on entry texts
// TODO: show dialog boxes for all button actions
// TODO: check name duplicates on lists and inform client
