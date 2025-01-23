package gui

import (
	"log/slog"

	"go-gui/pkg/config"
	"go-gui/pkg/tasks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WindowContent(window fyne.Window) *fyne.Container {
	inputs, hint, startBtn := SetupControls(window)

	inputs = append(inputs, container.NewGridWithColumns(1, startBtn))
	canvas := []fyne.CanvasObject{}
	for _, input := range inputs {
		canvas = append(canvas, input)
	}
	canvas = append(canvas, widget.NewLabel("Description:"))
	canvas = append(canvas, hint)

	go RegisterHotkeys(startBtn)

	return container.NewVBox(
		canvas...,
	)
}

func SetupControls(window fyne.Window) (inputs []*fyne.Container, hint *widget.TextGrid, btn *widget.Button) {
	input := widget.NewEntry()
	input.SetPlaceHolder("Window name")
	input.SetText(config.Settings.Window.Name)
	input.OnChanged = func(text string) {
		config.Settings.Window.Name = text
	}
	newContainer := container.NewGridWithColumns(1,
		container.NewVBox(widget.NewLabel("Window name:"), input))
	inputs = append(inputs, newContainer)

	inputs = append(inputs, container.NewGridWithColumns(3, widget.NewLabel("Delay:"),
		widget.NewLabel("Key:"), widget.NewLabel("Clip:")))
	for i, keySet := range config.Settings.Keys {
		input1 := widget.NewEntry()
		input1.SetPlaceHolder("Delay")
		input1.SetText(keySet.Delay)
		input1.OnChanged = func(text string) {
			config.Settings.Keys[i].Delay = text
		}

		input2 := widget.NewEntry()
		input2.SetPlaceHolder("Key")
		input2.SetText(keySet.Value)
		input2.OnChanged = func(text string) {
			config.Settings.Keys[i].Value = text
		}

		input3 := widget.NewEntry()
		input3.SetPlaceHolder("Random")
		input3.SetText(keySet.Clip)
		input3.OnChanged = func(text string) {
			config.Settings.Keys[i].Clip = text
		}

		newContainer := container.NewGridWithColumns(3, input1, input2, input3)
		inputs = append(inputs, newContainer)
	}

	hint = widget.NewTextGrid()
	btn = widget.NewButton("Start/Stop", func() {
		if err := tasks.StartBackgroundTask(hint); err != nil {
			slog.Error(err.Error())
		}
	})

	return inputs, hint, btn
}
