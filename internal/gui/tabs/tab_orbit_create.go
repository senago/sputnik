package tabs

import (
	"context"
	"fmt"
	"strconv"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewCreateOrbitTab(
	insertOrbit port.InsertOrbit,
) *helpers.Tab {
	nameEntry := widget.NewEntry()
	heightEntry := widget.NewEntry()

	form := widget.NewForm(
		widget.NewFormItem("Name", nameEntry),
		widget.NewFormItem("Height (km)", heightEntry),
	)

	output := widget.NewLabel("")

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

		output.SetText(fmt.Sprintf("created orbit [%s]", orbit.Name))
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
