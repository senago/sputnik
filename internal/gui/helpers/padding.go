package helpers

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
)

func PadContainer(obj fyne.CanvasObject) *fyne.Container {
	return PadContainerWithSize(obj, fyne.NewSize(16, 16))
}

func PadContainerWithSize(obj fyne.CanvasObject, s fyne.Size) *fyne.Container {
	rect := canvas.NewRectangle(color.Transparent)
	rect.SetMinSize(s)
	return container.NewBorder(rect, rect, rect, rect, obj)
}
