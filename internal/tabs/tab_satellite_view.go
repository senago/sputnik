package tabs

import (
	"context"
	"fmt"
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

func NewSatelliteViewTab(
	getSatellites port.GetSatellites,
) *container.TabItem {
	satellitesContainer := container.NewVBox()

	renderHandler := func() {
		satellites, err := getSatellites(context.Background())
		if err == nil {
			satellitesContainer.Objects = lo.Map(
				satellites,
				func(s domain.Satellite, _ int) fyne.CanvasObject {
					return canvas.NewText(
						fmt.Sprintf(
							"name=[%s] desc=[%s] orbitName=[%s]",
							s.Name, s.Description, s.Orbit.Name,
						),
						color.White,
					)
				},
			)

			satellitesContainer.Refresh()
		}
	}

	renderButton := widget.NewButton("render", renderHandler)

	return container.NewTabItem(
		"satellite view",
		PadContainer(
			container.New(
				layout.NewVBoxLayout(),
				renderButton,
				satellitesContainer,
			),
		),
	)
}
