package main

import (
	"fmt"
	"goscan/singleton"
	"goscan/widgets"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
	"time"
)

const FileExt = ".jpg"

// const scan_command = []string {
// 	"scanimage",
// 	"--format=jpeg",           // Set output format to PNG
// 	"--resolution", "600",     // Set the resolution (DPI)
// 	"-x", "210",               // Width in mm (210mm for A4)
// 	"-y", "297"                // Height in mm (297mm for A4)
// 	// "--mode", "Color"        // Color mode, can be "Gray" or "Lineart" as well
// }

// Scan handles the button click event
func OnScanAction(status *widgets.StatusWidget, imageWidget *widgets.ImageWidget) {

	outputFile := getScanFilePath()
	//outputFile := "scanned_image.jpg"

	status.SetWarn("Scanning...   " + outputFile)

	// Execute the scanimage command
	// cmd := exec.Command("scanimage", "--format=jpeg", "--output-file", outputFile)
	cmd := exec.Command(
		"scanimage",
		"--format=jpeg",
		"--resolution", "600",
		"-p",
		"--output-file", outputFile)
	fmt.Printf("os.Stdout: %v\n", os.Stdout)
	if err := cmd.Run(); err != nil {
		status.SetError("Error: " + err.Error())
		return
	}

	// Check if the file was created
	if _, err := os.Stat(outputFile); os.IsNotExist(err) {
		status.SetError("Error: " + err.Error())
		return
	}

	// Create and display the custom image widget
	imageWidget.LoadNewImage(outputFile)
	status.SetSuccess("Scanned    " + outputFile)

	// window.SetContent(imageWidget)

}

func getScanFilePath() string {
	settings := singleton.GetInstance()
	directory := settings.Directory
	filepattern := ReplaceDateTimePlaceholder(settings.FilePattern)
	return filepath.Join(directory, filepattern)
}

func ReplaceDateTimePlaceholder(input string) string {
	// Define the date-time format (e.g., "2006-01-02 15:04:05")
	// Modify the format string to fit your needs
	format := "20060102-150405" // Example: "YYYY-MM-DD HH-MM-SS"

	// Get the current time and format it
	currentTime := time.Now().Format(format)

	// Replace [DATETIME] with the formatted date-time string
	result := strings.Replace(input, "[DATETIME]", currentTime, -1)

	return result + FileExt
}
