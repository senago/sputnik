package tabs

import (
	"context"
	"log"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewTabInfo(
	insertOrbit port.InsertOrbit,
	insertSatellite port.InsertSatellite,
) *helpers.Tab {
	label := widget.NewLabel("sputnik")

	buttonLoadDefaults := widget.NewButton(
		"load defaults",
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

			for _, orbit := range orbits {
				if err := insertOrbit(context.Background(), orbit); err != nil {
					log.Printf("[loadDefaults] insertOrbit: %v", err)
					return
				}
			}

			for _, satellite := range satellites {
				if err := insertSatellite(context.Background(), satellite); err != nil {
					log.Printf("[loadDefaults] insertSatellite: %v", err)
					return
				}
			}
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
						buttonLoadDefaults,
					),
				),
			),
		),
	)
}
