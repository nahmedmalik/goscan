package main

import "fyne.io/fyne/v2"

func (app *Config) createMenuItems() {

	scanMenuItem := fyne.NewMenuItem("Open...", func() {})
	scanMenu := fyne.NewMenu("File", scanMenuItem)

	menu := fyne.NewMainMenu(scanMenu)

	app.MainWindow.SetMainMenu(menu)
}
