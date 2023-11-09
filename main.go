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
