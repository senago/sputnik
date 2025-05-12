package tabs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

func NewSatelliteUpdateTab(
	getOrbits port.GetOrbits,
	getSatellitesByNameLike port.GetSatellitesByNameLike,
	updateSatellite port.UpdateSatellite,
) *container.TabItem {
	output := widget.NewEntry()

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
	orbitNameEntry := widget.NewSelectEntry(nil)
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

			orbitNameEntry.SetOptions([]string{satellite.Orbit.Name})
			orbitNameEntry.SetText(satellite.Orbit.Name)

			descriptionEntry.SetText(satellite.Description)
			typeEntry.SetText(satellite.Type)
		}
	}

	orbitNameEntry.OnCursorChanged = func() {
		orbits := resolveOrbits()

		orbitNameEntry.SetOptions(lo.Keys(orbits))
	}

	form.OnSubmit = func() {
		orbit := resolveOrbits()[orbitNameEntry.Text]

		satellite := domain.Satellite{
			ID:          domain.SatelliteIDFromString(satelliteID.Text),
			Orbit:       orbit,
			Name:        satelliteNameEntry.Text,
			Description: descriptionEntry.Text,
			Type:        typeEntry.Text,
		}

		if err := updateSatellite(context.Background(), satellite); err != nil {
			output.SetText(fmt.Sprintf("updateSatellite: %s", err))
			return
		}

		output.SetText("successfully updated satellite")
	}

	return container.NewTabItem(
		"satellite update",
		PadContainer(
			container.NewVBox(
				form,
				output,
			),
		),
	)
}
