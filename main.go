// main.go
package main

import (
	"goscan/widgets"
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

type Config struct {
	App      fyne.App
	InfoLog  *log.Logger
	ErrorLog *log.Logger
	Status   *widgets.StatusWidget

	imageWidget *widgets.ImageWidget

	MainWindow fyne.Window
}

var myApp Config

func main() {
	// Create a new application
	fyneApp := app.NewWithID("com.naveim.goscan")
	myApp.App = fyneApp

	// create loggers
	myApp.InfoLog = log.New(os.Stdout, "INFO\t", log.Ldate|log.Ltime)
	myApp.ErrorLog = log.New(os.Stdout, "ERROR\t", log.Ldate|log.Ltime|log.Lshortfile)
	myApp.Status = widgets.NewStatusWidget()
	myApp.Status.SetInfo("Read to Scan")

	// imageWidget, button := myApp.makeUI()

	// create window
	myApp.MainWindow = fyneApp.NewWindow("Simple Document Scanner")
	myApp.MainWindow.Resize(fyne.NewSize(800, 800))
	myApp.MainWindow.SetMaster()

	myApp.makeUI()
	myApp.createMenuItems()

	// Show and run the application window
	myApp.MainWindow.ShowAndRun()
}
