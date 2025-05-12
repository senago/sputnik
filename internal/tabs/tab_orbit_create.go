package tabs

import (
	"context"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/port"
)

func NewCreateOrbitTab(
	insertOrbit port.InsertOrbit,
) *container.TabItem {
	nameEntry := widget.NewEntry()
	heightEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Height (km)", heightEntry),
	)

	output := widget.NewEntry()

	form.OnSubmit = func() {
		heightKM, err := strconv.ParseInt(heightEntry.Text, 10, 64)
		if err != nil {
			output.SetText(fmt.Sprintf("invalid height: %s", err))
			return
		}

		orbit := domain.Orbit{
			ID:       domain.NewOrbitID(),
			Name:     nameEntry.Text,
			HeightKm: heightKM,
		}

		if err := insertOrbit(context.Background(), orbit); err != nil {
			output.SetText(fmt.Sprintf("insertOrbit: %s", err))
			return
		}

		output.SetText("successfully created orbit")
	}

	return container.NewTabItem(
		"orbit create",
		PadContainer(
			container.NewVBox(
				form,
				output,
			),
		),
	)
}
