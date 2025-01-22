package tasks

import (
	"fmt"

	"go-gui/pkg/models"
	"go-gui/pkg/utils"

	"fyne.io/fyne/v2/widget"
)

var (
	stopChan  chan bool
	isRunning bool
)

func StartBackgroundTask(hint *widget.TextGrid) {
	if !isRunning {

		isRunning = true
		stopChan = make(chan bool)
		hint.SetText(fmt.Sprintf("Запущено с параметрами: %s и %s", models.CurrentInput.Value1, models.CurrentInput.Value2))

		go func() {
			for {
				select {
				case <-stopChan:
					hint.SetText("Остановлено")
					return
				default:
					utils.RunLogic()
				}
			}
		}()
	} else if isRunning {
		isRunning = false
		close(stopChan)
		hint.SetText("Остановлено")
	}
}
