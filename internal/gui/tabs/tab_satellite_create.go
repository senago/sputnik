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

	satelliteNameEntry := widget.NewEntry()
	orbitNameEntry := widget.NewSelect(nil, nil)
	descriptionEntry := widget.NewEntry()
	typeEntry := widget.NewSelect(
		domain.AllSatelliteType(),
		nil,
	)

	satelliteNameEntry.SetPlaceHolder("Satellite name")
	descriptionEntry.SetPlaceHolder("Satellite description")

	form := widget.NewForm(
		widget.NewFormItem("Name", satelliteNameEntry),
		widget.NewFormItem("Orbit Name", orbitNameEntry),
		widget.NewFormItem("Description", descriptionEntry),
		widget.NewFormItem("Type", typeEntry),
	)

	form.SubmitText = "Create"

	form.OnSubmit = func() {
		orbit := resolveOrbits()[orbitNameEntry.Selected]

		satellite := domain.Satellite{
			ID:          domain.NewSatelliteID(),
			Orbit:       orbit,
			Name:        satelliteNameEntry.Text,
			Description: descriptionEntry.Text,
			Type:        typeEntry.Selected,
		}

		if err := insertSatellite(context.Background(), satellite); err != nil {
			output.SetText(fmt.Sprintf("insertSatellite: %s", err))
			return
		}

		output.SetText("successfully created satellite")
	}

	loadOrbits := func() {
		orbits := resolveOrbits()
		orbitNameEntry.SetOptions(lo.Keys(orbits))
	}

	return helpers.NewTab(
		container.NewTabItem(
			"create satellite",
			helpers.PadContainer(
				container.NewVBox(
					form,
					container.NewCenter(
						output,
					),
				),
			),
		),
	).SetOnSelected(loadOrbits)
}
