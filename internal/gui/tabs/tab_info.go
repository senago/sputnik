package tabs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewTabInfo(
	insertOrbit port.InsertOrbit,
	insertSatellite port.InsertSatellite,
	truncateAll port.TruncateAll,
) *helpers.Tab {
	label := widget.NewRichTextFromMarkdown("# **sputnik**")
	output := widget.NewLabel("")

	buttonLoadDefaults := widget.NewButtonWithIcon(
		"load defaults",
		theme.StorageIcon(),
		func() {
			orbits := []domain.Orbit{
				{
					ID:       domain.NewOrbitID(),
					Name:     "orbit100",
					HeightKm: 100,
				},
				{
					ID:       domain.NewOrbitID(),
					Name:     "orbit200",
					HeightKm: 200,
				},
			}

			satellites := []domain.Satellite{
				{
					ID:    domain.NewSatelliteID(),
					Orbit: orbits[0],
					Position: domain.Position{
						X: 336,
						Y: 272,
					},
					Name:        "satellite1",
					Description: "resource-p",
					Type:        domain.SatelliteTypeResourceP,
				},
				{
					ID:    domain.NewSatelliteID(),
					Orbit: orbits[1],
					Position: domain.Position{
						X: 818,
						Y: 92,
					},
					Name:        "satellite2",
					Description: "kondor",
					Type:        domain.SatelliteTypeKondor,
				},
				{
					ID:    domain.NewSatelliteID(),
					Orbit: orbits[1],
					Position: domain.Position{
						X: 820,
						Y: 482,
					},
					Name:        "satellite3",
					Description: "kanopus",
					Type:        domain.SatelliteTypeKanopus,
				},
			}

			if err := truncateAll(context.Background()); err != nil {
				output.SetText(fmt.Sprintf("truncateAll: %v", err))
				return
			}

			for _, orbit := range orbits {
				if err := insertOrbit(context.Background(), orbit); err != nil {
					output.SetText(fmt.Sprintf("insertOrbit: %v", err))
					return
				}
			}

			for _, satellite := range satellites {
				if err := insertSatellite(context.Background(), satellite); err != nil {
					output.SetText(fmt.Sprintf("insertSatellite: %v", err))
					return
				}
			}

			output.SetText("loaded defaults")
		},
	)

	buttonReset := widget.NewButtonWithIcon(
		"reset",
		theme.DeleteIcon(),
		func() {
			if err := truncateAll(context.Background()); err != nil {
				output.SetText(fmt.Sprintf("truncateAll: %v", err))
				return
			}

			output.SetText("reset")
		},
	)

	return helpers.NewTab(
		container.NewTabItem(
			"info",
			helpers.PadContainer(
				container.NewCenter(
					container.NewVBox(
						container.NewCenter(
							label,
						),
						container.NewHBox(
							buttonLoadDefaults,
							buttonReset,
						),
						container.NewCenter(
							output,
						),
					),
				),
			),
		),
	)
}
