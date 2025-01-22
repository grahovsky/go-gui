package gui

import (
	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
)

func RunApp() {
	myApp := app.New()
	window := myApp.NewWindow("Control Panel")

	content := WindowContent(window)

	window.Resize(fyne.NewSize(400, 300))
	window.SetContent(content)

	window.ShowAndRun()
}
