package tabs

import (
	"context"
	"fmt"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/port"
)

func NewSatelliteViewTab(
	getSatellites port.GetSatellites,
	updateSatellites port.UpdateSatellites,
) *container.TabItem {
	satellitesCanvas := helpers.NewCanvas()

	var satellites []domain.Satellite

	loadHandler := func() {
		satellitesCanvas.RemoveAll()

		var err error
		satellites, err = getSatellites(context.Background())
		if err != nil {
			log.Println("[loadHandler] getSatellites", err)
			return
		}

		for _, s := range satellites {
			img := canvas.NewImageFromFile("./media/resource-p.png")
			img.FillMode = canvas.ImageFillContain
			img.SetMinSize(fyne.NewSize(50, 50))

			text := widget.NewLabel(
				fmt.Sprintf(
					"%s, %s, %s",
					s.Name, s.Description, s.Orbit.Name,
				),
			)
			text.Alignment = fyne.TextAlignCenter

			obj := container.NewVBox(
				img,
				text,
			)

			satellitesCanvas.AddAt(
				helpers.NewObject(s.Name, obj),
				fyne.NewPos(s.Position.X, s.Position.Y),
			)
		}
	}

	loadButton := widget.NewButton("load", loadHandler)
	loadButton.Resize(loadButton.MinSize())

	saveHandler := func() {
		positions := satellitesCanvas.ObjectPositions()

		for idx := range satellites {
			if pos, ok := positions[satellites[idx].Name]; ok {
				satellites[idx] = satellites[idx].SetPosition(domain.Position{
					X: pos.X,
					Y: pos.Y,
				})
			}
		}

		if err := updateSatellites(context.Background(), satellites); err != nil {
			log.Println("[saveHandler] updateSatellites:", err)
		}
	}

	saveButton := widget.NewButton("save", saveHandler)
	saveButton.Resize(saveButton.MinSize())

	image := canvas.NewImageFromFile("./media/earth.png")
	image.FillMode = canvas.ImageFillContain

	return container.NewTabItem(
		"satellite view",
		helpers.PadContainer(
			container.NewBorder(
				container.NewHBox(
					loadButton,
					saveButton,
				),
				nil,
				nil,
				nil,
				container.NewStack(
					image,
					satellitesCanvas,
				),
			),
		),
	)
}
