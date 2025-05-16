package helpers

import "fyne.io/fyne/v2"

type Object struct {
	fyne.CanvasObject

	id string
}

func NewObject(id string, obj fyne.CanvasObject) Object {
	return Object{
		CanvasObject: obj,
		id:           id,
	}
}
