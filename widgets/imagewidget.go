package widgets

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

// ImageWidget is a custom widget to display an image
type ImageWidget struct {
	widget.BaseWidget
	image *canvas.Image
}

// NewImageWidget creates a new instance of ImageWidget
func NewImageWidget(imagePath string) *ImageWidget {
	// Load the image from the file path
	img := canvas.NewImageFromFile(imagePath)
	img.SetMinSize(fyne.Size{Width: 800, Height: 800})
	img.FillMode = canvas.ImageFillContain // Set fill mode to contain the image within bounds

	widget := &ImageWidget{
		image: img,
	}

	// Initialize the widget
	widget.ExtendBaseWidget(widget)
	return widget
}

// LoadNewImage loads a new image into the existing widget
func (w *ImageWidget) LoadNewImage(imagePath string) {
	// Update the image object with a new file
	w.image.File = imagePath

	// Refresh the widget to display the new image
	w.image.Refresh()
}

// CreateRenderer returns a new widget renderer for the image widget
func (w *ImageWidget) CreateRenderer() fyne.WidgetRenderer {
	// Create a container to render the image
	container := container.NewStack(w.image)
	return &imageWidgetRenderer{container}
}

// imageWidgetRenderer implements the fyne.WidgetRenderer interface
type imageWidgetRenderer struct {
	container *fyne.Container
}

// Layout lays out the components of the renderer
func (r *imageWidgetRenderer) Layout(size fyne.Size) {
	r.container.Resize(size)
}

// BackgroundColor returns the background color of the widget
func (r *imageWidgetRenderer) BackgroundColor() color.Color {
	return color.RGBA{G: 100, A: 100}
}

// MinSize returns the minimum size of the renderer
func (r *imageWidgetRenderer) MinSize() fyne.Size {
	return r.container.MinSize()
}

// Refresh refreshes the renderer to update the display
func (r *imageWidgetRenderer) Refresh() {
	canvas.Refresh(r.container)
}

// Objects returns the objects to render
func (r *imageWidgetRenderer) Objects() []fyne.CanvasObject {
	return r.container.Objects
}

// Destroy cleans up resources when the widget is destroyed
func (r *imageWidgetRenderer) Destroy() {}
