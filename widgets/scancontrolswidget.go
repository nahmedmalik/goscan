package widgets

import (
	"goscan/singleton"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/dialog"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

// ScanControlsWidget represents the custom widget with the specified controls
type ScanControlsWidget struct {
	widget.BaseWidget
	dirEntry         *widget.Entry
	filePatternEntry *widget.Entry
	dirButton        *widget.Button
	scanButton       *widget.Button
	onScan           func(directory string, pattern string)
	window           fyne.Window
}

// NewScanControlsWidget creates a new instance of the ScanControlsWidget
func NewScanControlsWidget(w fyne.Window, onAction func(directory string, pattern string)) *ScanControlsWidget {
	scanWidget := &ScanControlsWidget{
		dirEntry:         widget.NewEntry(),
		filePatternEntry: widget.NewEntry(),
		dirButton:        widget.NewButtonWithIcon("Select Directory", theme.FolderOpenIcon(), nil),
		scanButton:       widget.NewButton("Scan", nil),
		onScan:           onAction,
		window:           w,
	}
	scanWidget.dirEntry.SetPlaceHolder("Enter directory path...")
	scanWidget.filePatternEntry.SetPlaceHolder("Enter file name pattern...")

	// text entry change
	scanWidget.dirEntry.OnChanged = scanWidget.onDirectoryEntryChanged
	scanWidget.filePatternEntry.OnChanged = scanWidget.onFilePatternEntryChanged

	// load settings
	settings := singleton.GetInstance()
	scanWidget.dirEntry.SetText(settings.Directory)
	scanWidget.filePatternEntry.SetText(settings.FilePattern)

	// Set button actions
	scanWidget.dirButton.OnTapped = scanWidget.openDirectoryDialog
	scanWidget.scanButton.OnTapped = scanWidget.Scan
	scanWidget.scanButton.Importance = widget.HighImportance

	scanWidget.ExtendBaseWidget(scanWidget)
	return scanWidget
}

// CreateRenderer defines how the custom widget will be rendered
func (c *ScanControlsWidget) CreateRenderer() fyne.WidgetRenderer {

	// Create a vertical box layout with all the controls
	controls := container.NewVBox(
		c.dirEntry,
		c.dirButton,
		c.filePatternEntry,
		c.scanButton,
	)

	return widget.NewSimpleRenderer(controls)
}

func (c *ScanControlsWidget) MinSize() fyne.Size {
	return fyne.NewSize(300, 200) // Set your desired fixed width (300) and height (200)
}

// openDirectoryDialog opens a directory selection dialog
func (c *ScanControlsWidget) openDirectoryDialog() {
	dialog.ShowFolderOpen(func(uri fyne.ListableURI, err error) {
		if err == nil && uri != nil {
			c.dirEntry.SetText(uri.Path()) // Set the selected directory path
		}
	}, c.window)
}

// onFilePatternEntryChanged
func (c *ScanControlsWidget) onFilePatternEntryChanged(filePattern string) {
	settings := singleton.GetInstance()
	settings.FilePattern = filePattern
	settings.Save()
}

func (c *ScanControlsWidget) onDirectoryEntryChanged(directory string) {
	settings := singleton.GetInstance()
	settings.Directory = directory
	settings.Save()
}

// Scan triggers the action associated with the widget
func (c *ScanControlsWidget) Scan() {
	if c.onScan != nil {
		c.onScan(c.dirEntry.Text, c.filePatternEntry.Text)
	}
}
