package tabs

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/dto"
	"github.com/senago/sputnik/internal/gui/helpers"
)

func NewTabSetup(
	applySettings func(settings dto.Settings) error,
) *helpers.Tab {
	helpText := canvas.NewText("Подключение к БД не установлено. Укажите DSN.", theme.Color(theme.ColorNameWarning))

	const defaultDSN = "host=localhost user=postgres database=sputnik password=admin"
	dsnEntry := widget.NewEntry()
	dsnEntry.SetText(defaultDSN)

	form := widget.NewForm(
		widget.NewFormItem("dsn", dsnEntry),
	)

	output := widget.NewLabel("")

	form.OnSubmit = func() {
		settings := dto.Settings{
			DSN: dsnEntry.Text,
		}

		if err := applySettings(settings); err != nil {
			output.SetText(err.Error())
		}
	}

	return helpers.NewTab(
		container.NewTabItem(
			"Настройка",
			helpers.PadContainer(
				container.NewVBox(
					container.NewCenter(
						helpText,
					),
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
