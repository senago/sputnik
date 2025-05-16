package helpers

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func PadContainer(c fyne.CanvasObject) *fyne.Container {
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(fyne.NewSize(16, 16))
	return container.NewBorder(rect, rect, rect, rect, c)
}
