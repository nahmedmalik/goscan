package main

import (
	"goscan/widgets"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func (myApp *Config) makeUI() {
	myApp.imageWidget = widgets.NewImageWidget("./test.jpg")
	myApp.Status = widgets.NewStatusWidget()

	button := widget.NewButton("Scan", func() {
		//app.imageWidget.LoadNewImage("scanned_image.png")
		ButtonClickHandler(myApp.imageWidget)
	})
	button.Importance = widget.HighImportance

	// file list
	// Create a directory path input
	directoryEntry := widget.NewEntry()
	directoryEntry.SetPlaceHolder("Enter directory path...")

	// Create a file list widget and add it to the window
	scanWidget := widgets.NewScanControlsWidget(myApp.MainWindow, func(directory string, pattern string) {
		OnScanAction(myApp.Status, myApp.imageWidget)
	})
	contentContainer := container.NewBorder(
		myApp.Status, myApp.imageWidget, nil, nil, nil,
	)

	myApp.MainWindow.SetContent(container.NewHSplit(scanWidget, contentContainer))

}
