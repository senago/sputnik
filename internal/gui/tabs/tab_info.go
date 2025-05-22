package tabs

import (
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/gui/helpers"
)

func NewTabInfo() *helpers.Tab {
	label := widget.NewLabel("sputnik app")

	return helpers.NewTab(
		container.NewTabItem(
			"info",
			helpers.PadContainer(
				container.NewCenter(
					label,
				),
			),
		),
	)
}
