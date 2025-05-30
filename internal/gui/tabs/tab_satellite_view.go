package tabs

import (
	"context"
	"fmt"
	"image/color"
	"log"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
	"github.com/samber/lo"
	"github.com/senago/sputnik/internal/domain"
	"github.com/senago/sputnik/internal/gui/helpers"
	"github.com/senago/sputnik/internal/media"
	"github.com/senago/sputnik/internal/port"
)

// the hell is that?
const earthPadding = 300

func NewSatelliteViewTab(
	getSatellites port.GetSatellites,
	updateSatellites port.UpdateSatellites,
) *helpers.Tab {
	circlesContainer := container.NewStack()
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
		circlesContainer.RemoveAll()

		var err error
		satellites, err = getSatellites(context.Background())
		if err != nil {
			log.Println("[loadHandler] getSatellites", err)
			return
		}

		orbits := lo.UniqBy(
			lo.Map(
				satellites,
				func(s domain.Satellite, _ int) domain.Orbit {
					return s.Orbit
				},
			),
			func(o domain.Orbit) string {
				return o.Name
			},
		)

		orbitCircles := lo.Map(
			orbits,
			func(o domain.Orbit, _ int) fyne.CanvasObject {
				circle := canvas.NewCircle(color.Transparent)
				circle.StrokeColor = color.RGBA{
					R: 20,
					G: 50,
					B: 225,
					A: 85,
				}
				circle.StrokeWidth = 8

				label := widget.NewLabel(o.Name)
				label.Alignment = fyne.TextAlignCenter
				label.SizeName = theme.SizeNameCaptionText
				label.TextStyle.Bold = true

				return container.NewStack(
					helpers.PadContainerWithSize(
						label,
						fyne.NewSquareSize(theme.CaptionTextSize()/2),
					),
					circle,
				)
			},
		)

		for idx, circle := range orbitCircles {
			antiPadding := float32(orbits[idx].HeightKm * 3 / 4)

			circlesContainer.Add(
				helpers.PadContainerWithSize(
					circle,
					fyne.NewSize(earthPadding-antiPadding, earthPadding-antiPadding),
				),
			)
		}

		for _, s := range satellites {
			img := canvas.NewImageFromImage(media.GetSatelliteImage(s.Type))
			img.FillMode = canvas.ImageFillContain
			img.SetMinSize(fyne.NewSize(50, 50))

			text := widget.NewLabel(
				fmt.Sprintf(
					"[%s]\n[%s]",
					s.Name, s.Description,
				),
			)
			text.Alignment = fyne.TextAlignCenter
			text.SizeName = theme.SizeNameCaptionText

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

	image := canvas.NewImageFromImage(media.GetEarthImage())
	image.FillMode = canvas.ImageFillContain

	return helpers.NewTab(
		container.NewTabItem(
			"Представление",
			helpers.PadContainer(
				container.NewStack(
					helpers.PadContainerWithSize(
						image,
						fyne.NewSquareSize(earthPadding),
					),
					circlesContainer,
					satellitesCanvas,
				),
			),
		),
	).SetOnSelected(loadHandler)
}
