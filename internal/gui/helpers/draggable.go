package helpers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/widget"
)

type DraggableItem struct {
	widget.BaseWidget

	parent fyne.Widget
	object fyne.CanvasObject
}

func NewDraggableItem(parent fyne.Widget, obj fyne.CanvasObject) *DraggableItem {
	d := &DraggableItem{
		parent: parent,
		object: obj,
	}

	d.ExtendBaseWidget(d)
	d.Resize(obj.MinSize())

	return d
}

func (d *DraggableItem) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(d.object)
}

func (d *DraggableItem) Dragged(event *fyne.DragEvent) {
	d.Move(fyne.NewPos(event.Dragged.DX, event.Dragged.DY))
}

func (d *DraggableItem) Move(pos fyne.Position) {
	diff := d.Position().AddXY(pos.X, pos.Y)

	diff.X = max(diff.X, 0)
	diff.X = min(diff.X, d.parent.Size().Width-d.Size().Width)

	diff.Y = max(diff.Y, 0)
	diff.Y = min(diff.Y, d.parent.Size().Height-d.Size().Height)

	d.BaseWidget.Move(diff)
}

func (d *DraggableItem) DragEnd() {}
