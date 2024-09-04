package main

import (
	"goscan/widgets"
	"os"
	"os/exec"
)

// ButtonClickHandler handles the button click event
func ButtonClickHandler(imageWidget *widgets.ImageWidget) {
	outputFile := "scanned_image.png"

	// Execute the scanimage command
	cmd := exec.Command("scanimage", "--format=png", "--output-file", outputFile)
	if err := cmd.Run(); err != nil {
		// window.SetContent(widget.NewLabel("Error: " + err.Error()))
		return
	}

	// Check if the file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		// window.SetContent(widget.NewLabel("Error: Image file not created"))
		return
	}

	// Create and display the custom image widget
	imageWidget.LoadNewImage(outputFile)

	// Update the window content to display the widget

	// window.SetContent(imageWidget)

}
