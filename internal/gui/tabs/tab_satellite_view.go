package tabs

import (
	"context"
	"fmt"
	"log"
	"path"

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
) *helpers.Tab {
	satellitesCanvas := helpers.NewCanvas()

	var satellites []domain.Satellite

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

	satellitesCanvas.SetOnChange(saveHandler)

	loadHandler := func() {
		satellitesCanvas.RemoveAll()

		var err error
		satellites, err = getSatellites(context.Background())
		if err != nil {
			log.Println("[loadHandler] getSatellites", err)
			return
		}

		for _, s := range satellites {
			img := canvas.NewImageFromFile(path.Join("./media", s.Type+".png"))
			img.FillMode = canvas.ImageFillContain
			img.SetMinSize(fyne.NewSize(50, 50))

			text := widget.NewLabel(
				fmt.Sprintf(
					"[%s] [%s] [%s]",
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

	image := canvas.NewImageFromFile("./media/earth.png")
	image.FillMode = canvas.ImageFillContain

	return helpers.NewTab(
		container.NewTabItem(
			"satellite view",
			helpers.PadContainer(
				container.NewStack(
					helpers.PadContainerWithSize(
						image,
						fyne.NewSize(150, 150),
					),
					satellitesCanvas,
				),
			),
		),
	).SetOnSelected(loadHandler)
}
