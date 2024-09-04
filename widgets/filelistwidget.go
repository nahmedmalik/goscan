package widgets

import (
	"log"
	"os"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

// FileListWidget represents a custom widget to display files in a directory
type FileListWidget struct {
	widget.List
	files []string
}

// NewFileListWidget creates a new FileListWidget instance
func NewFileListWidget(directoryPath string) *FileListWidget {
	fileListWidget := &FileListWidget{
		files: listFiles(directoryPath), // Load initial file list
	}

	fileListWidget.List = *widget.NewList(
		func() int { return len(fileListWidget.files) },
		func() fyne.CanvasObject {
			return widget.NewLabel("")
		},
		func(id widget.ListItemID, item fyne.CanvasObject) {
			item.(*widget.Label).SetText(fileListWidget.files[id])
		},
	)

	// Ensure the widget will expand to fill available space
	fileListWidget.ExtendBaseWidget(fileListWidget)

	return fileListWidget
}

// Update refreshes the file list based on the given directory path
func (fl *FileListWidget) Update(directoryPath string) {
	fl.files = listFiles(directoryPath)
	fl.Refresh()
}

// listFiles reads the directory and returns a slice of file names
func listFiles(directoryPath string) []string {
	entries, err := os.ReadDir(directoryPath)
	if err != nil {
		log.Printf("Error reading directory: %v", err)
		return []string{}
	}

	var files []string
	for _, entry := range entries {
		if !entry.IsDir() { // Exclude directories, only list files
			files = append(files, entry.Name())
		}
	}

	return files
}
