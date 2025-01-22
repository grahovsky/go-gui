package gui

import (
	"go-gui/pkg/models"
	"go-gui/pkg/tasks"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

func WindowContent(window fyne.Window) *fyne.Container {
	input1, input2, input3, hint, startBtn := SetupControls(window)

	inputs := container.NewGridWithColumns(3,
		container.NewVBox(widget.NewLabel("Delay:"), input1),
		container.NewVBox(widget.NewLabel("Key:"), input2),
		container.NewVBox(widget.NewLabel("Window name:"), input3),
	)
	// inputs := container.NewGridWithColumns(2,
	// 	input1,
	// 	input2,
	// )
	// buttons := container.NewGridWithColumns(2, startBtn, stopBtn)
	buttons := container.NewGridWithColumns(1, startBtn)

	go RegisterHotkeys(startBtn)

	return container.NewVBox(
		inputs,
		buttons,
		widget.NewLabel("Description:"),
		hint,
	)
}

func SetupControls(window fyne.Window) (input1, input2, input3 *widget.Entry, hint *widget.TextGrid, startBtn *widget.Button) {
	input1 = widget.NewEntry()
	input1.SetPlaceHolder("Delay")
	input1.SetText(models.CurrentInput.Value1)
	input1.OnChanged = func(text string) {
		models.CurrentInput.Value1 = text
	}

	input2 = widget.NewEntry()
	input2.SetPlaceHolder("Key")
	input2.SetText(models.CurrentInput.Value2)
	input2.OnChanged = func(text string) {
		models.CurrentInput.Value2 = text
	}

	input3 = widget.NewEntry()
	input3.SetPlaceHolder("Window name")
	input3.SetText(models.CurrentInput.Value3)
	input3.OnChanged = func(text string) {
		models.CurrentInput.Value3 = text
	}

	hint = widget.NewTextGrid()
	hint.SetText("Здесь будет отображаться подсказка")

	startBtn = widget.NewButton("Start/Stop", func() {
		tasks.StartBackgroundTask(hint)
	})

	return input1, input2, input3, hint, startBtn
}
