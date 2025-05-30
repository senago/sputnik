package tabs

import (
	"context"
	"fmt"

	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

const tabInfoHelp = `
# **sputnik**

Группа ИУ8-84: Лебедева, Киорогло, Мирзоян

Весна 2025
`

func NewTabInfo(
	insertOrbit port.InsertOrbit,
	insertSatellite port.InsertSatellite,
	truncateAll port.TruncateAll,
) *helpers.Tab {
	label := widget.NewRichTextFromMarkdown(tabInfoHelp)

	output := widget.NewLabel("")

	buttonLoadDefaults := widget.NewButtonWithIcon(
		"Применить конфигурацию по умолчанию",
		theme.StorageIcon(),
		func() {
			orbits := []domain.Orbit{
				{
					ID:       domain.NewOrbitID(),
					Name:     "Полярная орбита",
					HeightKm: 50,
				},
				{
					ID:       domain.NewOrbitID(),
					Name:     "Высокоэллиптическая орбита",
					HeightKm: 250,
				},
				{
					ID:       domain.NewOrbitID(),
					Name:     "Геостационарная орбита",
					HeightKm: 425,
				},
			}

			satellites := []domain.Satellite{
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[0],
					Position:    domain.NewPosition(1027, 319),
					Name:        "АРКТИКА-М №2",
					Description: "16.12.2023",
					Type:        domain.SatelliteTypeResourceP,
				},

				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[1],
					Position:    domain.NewPosition(542, 145),
					Name:        "АРКТИКА-М №1",
					Description: "28.02.2021",
					Type:        domain.SatelliteTypeResourceP,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[1],
					Position:    domain.NewPosition(732, 755),
					Name:        "ЭЛЕКТРО-Л №3",
					Description: "24.12.2019",
					Type:        domain.SatelliteTypeResourceP,
				},

				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(454, 808),
					Name:        "КАНОПУС-В №5",
					Description: "27.12.2018",
					Type:        domain.SatelliteTypeKanopus,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(348, 695),
					Name:        "КАНОПУС-В №3",
					Description: "01.02.2018",
					Type:        domain.SatelliteTypeKanopus,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(283, 572),
					Name:        "КАНОПУС-В-ИК",
					Description: "14.07.2017",
					Type:        domain.SatelliteTypeKanopus,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(963, 33),
					Name:        "ЭЛЕКТРО-Л №2",
					Description: "11.12.2015",
					Type:        domain.SatelliteTypeElectroL,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(261, 424),
					Name:        "МЕТЕОР-М №2-4",
					Description: "29.02.2024",
					Type:        domain.SatelliteTypeMeteorM,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(281, 282),
					Name:        "МЕТЕОР-М №2-7",
					Description: "27.06.2023",
					Type:        domain.SatelliteTypeMeteorM,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(357, 129),
					Name:        "МЕТЕОР-М №2-2",
					Description: "05.07.2019",
					Type:        domain.SatelliteTypeMeteorM,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(1128, 186),
					Name:        "КОНДОР ФКА №2",
					Description: "2024",
					Type:        domain.SatelliteTypeKondor,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(1181, 412),
					Name:        "РЕСУРС-ПМ №1",
					Description: "2024, 2025",
					Type:        domain.SatelliteTypeResourceP,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(1152, 574),
					Name:        "РЕСУРС-П №5",
					Description: "2025",
					Type:        domain.SatelliteTypeResourceP,
				},
				{
					ID:          domain.NewSatelliteID(),
					Orbit:       orbits[2],
					Position:    domain.NewPosition(1047, 751),
					Name:        "РЕСУРС-П №4",
					Description: "31.03.2024",
					Type:        domain.SatelliteTypeResourceP,
				},
			}

			if err := truncateAll(context.Background()); err != nil {
				output.SetText(fmt.Sprintf("truncateAll: %v", err))
				return
			}

			for _, orbit := range orbits {
				if err := insertOrbit(context.Background(), orbit); err != nil {
					output.SetText(fmt.Sprintf("insertOrbit: %v", err))
					return
				}
			}

			for _, satellite := range satellites {
				if err := insertSatellite(context.Background(), satellite); err != nil {
					output.SetText(fmt.Sprintf("insertSatellite: %v", err))
					return
				}
			}

			output.SetText("Конфигурация по умолчанию загружена")
		},
	)

	buttonReset := widget.NewButtonWithIcon(
		"Сбросить",
		theme.DeleteIcon(),
		func() {
			if err := truncateAll(context.Background()); err != nil {
				output.SetText(fmt.Sprintf("truncateAll: %v", err))
				return
			}

			output.SetText("Сброс совершен")
		},
	)

	return helpers.NewTab(
		container.NewTabItem(
			"Панель управления",
			helpers.PadContainer(
				container.NewCenter(
					container.NewVBox(
						container.NewCenter(
							label,
						),
						container.NewHBox(
							buttonLoadDefaults,
							buttonReset,
						),
						container.NewCenter(
							output,
						),
					),
				),
			),
		),
	)
}
