package tabs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewSatelliteCreateTab(
	getOrbits port.GetOrbits,
	insertSatellite port.InsertSatellite,
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

	satelliteNameEntry := widget.NewEntry()
	orbitNameEntry := widget.NewSelectEntry(nil)
	descriptionEntry := widget.NewEntry()
	typeEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Name", satelliteNameEntry),
		widget.NewFormItem("Orbit Name", orbitNameEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Type", typeEntry),
	)

	orbitNameEntry.OnCursorChanged = func() {
		orbits := resolveOrbits()

		orbitNameEntry.SetOptions(lo.Keys(orbits))
	}

	form.OnSubmit = func() {
		orbit := resolveOrbits()[orbitNameEntry.Text]

		satellite := domain.Satellite{
			ID:          domain.NewSatelliteID(),
			Orbit:       orbit,
			Name:        satelliteNameEntry.Text,
			Description: descriptionEntry.Text,
			Type:        typeEntry.Text,
		}

		if err := insertSatellite(context.Background(), satellite); err != nil {
			output.SetText(fmt.Sprintf("insertSatellite: %s", err))
			return
		}

		output.SetText("successfully created satellite")
	}

	return container.NewTabItem(
		"create satellite",
		helpers.PadContainer(
			container.NewVBox(
				form,
				output,
			),
		),
	)
}
