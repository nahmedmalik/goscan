package widgets

import (
	"net/url"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// StatusWidget is a custom widget to display status messages
type StatusWidget struct {
	widget.BaseWidget
	label *widget.Label
	// hyperlink *widget.Hyperlink
}

// NewStatusWidget creates a new StatusWidget
func NewStatusWidget() *StatusWidget {

	statusWidget := &StatusWidget{
		label: widget.NewLabel("Click Scan Button to Start Scanning"),
		// hyperlink: widget.NewHyperlink("", FileURI("")),
		// hyperlink: widget.NewHyperlink("link", &url.URL{Path: "file:///home/nahmed/Pictures/Scans/SimpleScan-20240903-212424.jpg"}),
	}

	// statusWidget.hyperlink.OnTapped = statusWidget.openFileFromHyperLink

	statusWidget.ExtendBaseWidget(statusWidget)
	return statusWidget
}

// CreateRenderer implements the fyne.Widget interface
func (w *StatusWidget) CreateRenderer() fyne.WidgetRenderer {
	controls := container.NewHBox(
		w.label,
		// w.hyperlink,
	)
	return widget.NewSimpleRenderer(controls)
}

func (w *StatusWidget) SetInfo(message string) {
	w.label.SetText(message)
	w.label.Importance = widget.MediumImportance
	w.label.Refresh()
}
func (w *StatusWidget) SetError(message string) {
	w.label.SetText(message)
	w.label.Importance = widget.DangerImportance
	w.label.Refresh()
}
func (w *StatusWidget) SetSuccess(message string) {
	w.label.SetText(message)
	w.label.Importance = widget.SuccessImportance
	w.label.Refresh()
}

// func (w *StatusWidget) SetFileLink(message string, url string) {
// 	w.label.Hidden = true
// 	// w.hyperlink.SetText(message + url)
// 	// w.hyperlink.SetURLFromString(url)
// 	w.label.Importance = widget.HighImportance
// 	w.label.Refresh()
// }

func (w *StatusWidget) SetWarn(message string) {
	w.label.SetText(message)
	w.label.Importance = widget.WarningImportance
	w.label.Refresh()
}

func FileURI(fileUri string) *url.URL {
	uri, _ := url.Parse(fileUri)
	return uri
}
