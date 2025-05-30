package gui

import (
	"context"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/driver/desktop"
	"github.com/senago/sputnik/internal/dto"
	"github.com/senago/sputnik/internal/gui/helpers"
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
	window.Resize(fyne.NewSize(1600, 900))

	fyneTabs := container.NewAppTabs()

	baseTabs := []*helpers.Tab{
		tabs.NewTabInfo(
			deps.PortInsertOrbit(),
			deps.PortInsertSatellite(),
			deps.PortTruncateAll(),
		),
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
	}

	appTabs := func() []*helpers.Tab {
		if !deps.IsConnectedToDB() {
			return []*helpers.Tab{
				tabs.NewTabSetup(func(settings dto.Settings) error {
					if err := deps.ConnectDB(context.Background(), settings.DSN); err != nil {
						return err
					}

					redrawTabs(fyneTabs, baseTabs)

					return nil
				}),
			}
		}

		return baseTabs
	}()

	redrawTabs(fyneTabs, appTabs)

	window.SetContent(fyneTabs)

	window.CenterOnScreen()

	for _, shortcut := range closeShortcuts {
		window.Canvas().AddShortcut(shortcut, func(fyne.Shortcut) { window.Close() })
	}

	return window
}

func redrawTabs(tabsContainer *container.AppTabs, tabs []*helpers.Tab) {
	for _, item := range tabsContainer.Items {
		tabsContainer.Remove(item)
	}

	for _, tab := range tabs {
		tabsContainer.Append(tab.TabItem)
	}

	tabsContainer.OnSelected = func(ti *container.TabItem) {
		for _, appTab := range tabs {
			appTab.OnSelected(ti)
		}
	}
}
