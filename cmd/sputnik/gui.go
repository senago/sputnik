package main

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/senago/sputnik/internal/tabs"
)

var closeShortcuts = []fyne.Shortcut{
	&desktop.CustomShortcut{KeyName: fyne.KeyC, Modifier: fyne.KeyModifierControl},
	&desktop.CustomShortcut{KeyName: fyne.KeyZ, Modifier: fyne.KeyModifierControl},
	&desktop.CustomShortcut{KeyName: fyne.KeyW, Modifier: fyne.KeyModifierControl},
}

func startApp(appConfig *App) {
	fyneApp := app.New()

	window := fyneApp.NewWindow("sputnik")
	window.Resize(fyne.NewSize(600, 300))
	window.SetContent(container.NewAppTabs(
		tabs.NewSatelliteViewTab(
			appConfig.PortGetSatellites(),
		),
		tabs.NewSatelliteCreateTab(
			appConfig.PortGetOrbits(),
			appConfig.PortInsertSatellite(),
		),
		tabs.NewCreateOrbitTab(
			appConfig.PortInsertOrbit(),
		),
		tabs.NewSatelliteUpdateTab(
			appConfig.PortGetOrbits(),
			appConfig.PortGetSatellitesByNameLike(),
			appConfig.PortUpdateSatellite(),
		),
		tabs.NewSatelliteDeleteTab(
			appConfig.PortGetSatellitesByNameLike(),
			appConfig.PortDeleteSatellites(),
		),
	))
	window.CenterOnScreen()

	for _, shortcut := range closeShortcuts {
		window.Canvas().AddShortcut(shortcut, func(fyne.Shortcut) { window.Close() })
	}

	window.ShowAndRun()
}
