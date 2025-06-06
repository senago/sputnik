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

func NewSatelliteDeleteTab(
	getSatellitesByNameLike port.GetSatellitesByNameLike,
	deleteSatellites port.DeleteSatellites,
) *helpers.Tab {
	output := widget.NewLabel("")

	satelliteID := widget.NewEntry()
	satelliteNameEntry := widget.NewSelectEntry(nil)

	form := widget.NewForm(
		widget.NewFormItem("Name", satelliteNameEntry),
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
			satelliteNameEntry.TextStyle.Bold = true
			satelliteNameEntry.Refresh()

			satellite := satellites[0]

			satelliteID.SetText(satellite.ID.String())
		}
	}

	form.OnSubmit = func() {
		if err := deleteSatellites(
			context.Background(),
			[]domain.SatelliteID{
				domain.SatelliteIDFromString(satelliteID.Text),
			},
		); err != nil {
			output.SetText(fmt.Sprintf("deleteSatellites: %s", err))
			return
		}

		output.SetText(fmt.Sprintf("deleted satellite [%s]", satelliteID.Text))
	}

	return helpers.NewTab(
		container.NewTabItem(
			"Создание орбиты",
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
	)
}
