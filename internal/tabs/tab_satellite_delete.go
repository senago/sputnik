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

func NewSatelliteDeleteTab(
	getSatellitesByNameLike port.GetSatellitesByNameLike,
	deleteSatellites port.DeleteSatellites,
) *container.TabItem {
	output := widget.NewEntry()

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

		output.SetText(fmt.Sprintf("successfully deleted satellite id=[%v]", satelliteID.Text))
	}

	return container.NewTabItem(
		"satellite delete",
		PadContainer(
			container.NewVBox(
				form,
				output,
			),
		),
	)
}
