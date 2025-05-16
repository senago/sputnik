package helpers

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type Canvas struct {
	widget.BaseWidget
	content *fyne.Container
	mapping []string
}

func NewCanvas() *Canvas {
	c := &Canvas{
		content: container.NewWithoutLayout(),
	}
	c.ExtendBaseWidget(c)
	return c
}

// CreateRenderer implements fyne.Widget.
func (c *Canvas) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(c.content)
}

// Add adds a new draggable object to the container
// at the specified position.
func (c *Canvas) AddAt(obj Object, pos fyne.Position) {
	draggable := NewDraggableItem(
		c,
		obj.CanvasObject,
	)

	c.content.Add(draggable)
	c.mapping = append(c.mapping, obj.id)

	draggable.Move(pos)
}

func (c *Canvas) RemoveAll() {
	c.content.RemoveAll()
}

func (c *Canvas) ObjectPositions() map[string]fyne.Position {
	positions := make(map[string]fyne.Position, len(c.content.Objects))

	for idx, obj := range c.content.Objects {
		positions[c.mapping[idx]] = obj.Position()
	}

	return positions
}
