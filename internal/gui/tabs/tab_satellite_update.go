package tabs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewSatelliteUpdateTab(
	getOrbits port.GetOrbits,
	getSatellitesByNameLike port.GetSatellitesByNameLike,
	updateSatellites port.UpdateSatellites,
) *helpers.Tab {
	output := widget.NewLabel("")

	resolveOrbits := func() map[string]domain.Orbit {
		orbits, err := getOrbits(context.Background())
		if err != nil {
			output.SetText(fmt.Sprintf("getOrbits: %s", err))
			return nil
		}

		return lo.SliceToMap(
			orbits,
			func(orbit domain.Orbit) (string, domain.Orbit) {
				return orbit.Name, orbit
			},
		)
	}

	satelliteID := widget.NewEntry()
	satelliteNameEntry := widget.NewSelectEntry(nil)
	orbitNameEntry := widget.NewSelect(nil, nil)
	descriptionEntry := widget.NewEntry()
	typeEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Name", satelliteNameEntry),
		widget.NewFormItem("Orbit Name", orbitNameEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Type", typeEntry),
	)

	satelliteNameEntry.OnChanged = func(s string) {
		satellites, err := getSatellitesByNameLike(context.Background(), s)
		if err != nil {
			output.SetText(fmt.Sprintf("getSatellitesByNameLike: %s", err))
			return
		}

		satelliteNameEntry.SetOptions(lo.Map(
			satellites,
			func(s domain.Satellite, _ int) string {
				return s.Name
			},
		))

		if len(satellites) == 1 {
			satellite := satellites[0]

			satelliteID.SetText(satellite.ID.String())

			orbitNameEntry.SetSelected(satellite.Orbit.Name)

			descriptionEntry.SetText(satellite.Description)
			typeEntry.SetText(satellite.Type)
		}
	}

	form.OnSubmit = func() {
		orbit := resolveOrbits()[orbitNameEntry.Selected]

		satellite := domain.Satellite{
			ID:          domain.SatelliteIDFromString(satelliteID.Text),
			Orbit:       orbit,
			Name:        satelliteNameEntry.Text,
			Description: descriptionEntry.Text,
			Type:        typeEntry.Text,
		}

		if err := updateSatellites(context.Background(), []domain.Satellite{satellite}); err != nil {
			output.SetText(fmt.Sprintf("updateSatellite: %s", err))
			return
		}

		output.SetText(fmt.Sprintf("updated satellite [%s]", satellite.Name))
	}

	loadOrbits := func() {
		orbits := resolveOrbits()
		orbitNameEntry.SetOptions(lo.Keys(orbits))
	}

	return helpers.NewTab(
		container.NewTabItem(
			"Обновление спутника",
			helpers.PadContainer(
				container.NewVBox(
					helpers.PadContainerWithSize(
						form,
						fyne.NewSize(300, 0),
					),
					container.NewCenter(
						output,
					),
				),
			),
		),
	).SetOnSelected(loadOrbits)
}
