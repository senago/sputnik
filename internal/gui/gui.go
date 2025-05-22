package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/senago/sputnik/internal/gui/tabs"
	"github.com/senago/sputnik/internal/ioc"
)

var closeShortcuts = []fyne.Shortcut{
	&desktop.CustomShortcut{KeyName: fyne.KeyC, Modifier: fyne.KeyModifierControl},
	&desktop.CustomShortcut{KeyName: fyne.KeyZ, Modifier: fyne.KeyModifierControl},
	&desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierControl},
}

func New(deps *ioc.Container) fyne.Window {
	fyneApp := app.New()

	window := fyneApp.NewWindow("sputnik")
	window.Resize(fyne.NewSize(1280, 720))

	window.SetContent(container.NewAppTabs(
		tabs.NewSatelliteViewTab(
			deps.PortGetSatellites(),
			deps.PortUpdateSatellite(),
		),
		tabs.NewSatelliteCreateTab(
			deps.PortGetOrbits(),
			deps.PortInsertSatellite(),
		),
		tabs.NewCreateOrbitTab(
			deps.PortInsertOrbit(),
		),
		tabs.NewSatelliteUpdateTab(
			deps.PortGetOrbits(),
			deps.PortGetSatellitesByNameLike(),
			deps.PortUpdateSatellite(),
		),
		tabs.NewSatelliteDeleteTab(
			deps.PortGetSatellitesByNameLike(),
			deps.PortDeleteSatellites(),
		),
	))

	window.CenterOnScreen()

	for _, shortcut := range closeShortcuts {
		window.Canvas().AddShortcut(shortcut, func(fyne.Shortcut) { window.Close() })
	}

	return window
}
