package tasks

import (
	"fmt"

	"go-gui/pkg/config"
	"go-gui/pkg/utils"

	"fyne.io/fyne/v2/widget"
)

var (
	stopChan  chan bool
	isRunning bool
)

func StartBackgroundTask(hint *widget.TextGrid) (result error) {
	stopChan = make(chan bool)
	if !isRunning {

		hint.SetText(fmt.Sprintf("Start with params: %s, %s", config.Settings.Keys[0].Value, config.Settings.Keys[0].Delay))
		if err := utils.GetActivePid(); err != nil {
			hint.SetText(err.Error())
			return err
		}
		isRunning = true

		for _, keySet := range config.Settings.Keys {
			go func() {
				for {
					select {
					case <-stopChan:
						hint.SetText("Stopped")
						return
					default:
						if err := utils.RunLogic(&keySet); err != nil {
							result = err
							fmt.Println(err)
							close(stopChan)
						}
					}
				}
			}()
		}
	} else if isRunning {
		isRunning = false
		close(stopChan)
		hint.SetText("Stopped")
	}

	return result
}
