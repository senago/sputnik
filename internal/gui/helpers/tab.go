package helpers

import "fyne.io/fyne/v2/container"

type Tab struct {
	*container.TabItem
	onSelected func()
}

func NewTab(ti *container.TabItem) *Tab {
	return &Tab{
		TabItem:    ti,
		onSelected: nil,
	}
}

func (tab *Tab) SetOnSelected(onSelected func()) *Tab {
	tab.onSelected = onSelected

	return tab
}

func (tab *Tab) OnSelected(ti *container.TabItem) {
	if ti.Text != tab.Text {
		return
	}

	if tab.onSelected != nil {
		tab.onSelected()
	}
}
