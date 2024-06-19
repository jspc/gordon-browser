package main

import (
	"fmt"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/theme"
	"fyne.io/fyne/v2/widget"
)

var (
	tabs   *container.AppTabs
	window fyne.Window
)

func main() {
	myApp := app.New()
	myApp.SetIcon(theme.FyneLogo())
	myApp.Settings().SetTheme(NewGordonTheme())

	contentWidth = fyne.MeasureText(fmt.Sprintf("%90s", ""), theme.TextSize(), fyne.TextStyle{}).Width

	window = myApp.NewWindow("Gordon")

	tabs = container.NewAppTabs(
		container.NewTabItemWithIcon("Home", theme.HomeIcon(), homeTab()),
	)

	window.SetContent(tabs)
	window.ShowAndRun()
}

func homeTab() fyne.CanvasObject {
	addrEntry := widget.NewEntry()
	addrEntry.Text = "//gordon.beasts.jspc.pw"

	f := widget.NewForm(
		widget.NewFormItem("Address", addrEntry),
	)

	f.OnSubmit = func() {
		tabs.Append(
			CreateTab(addrEntry.Text),
		)

		tabs.SelectIndex(len(tabs.Items) - 1)
	}

	return container.NewVBox(
		widget.NewLabelWithStyle("Welcome to Gordon", fyne.TextAlignLeading, fyne.TextStyle{
			Bold: true,
		}),
		widget.NewLabel("To get started, enter a gordon link below"),
		f,
	)
}
