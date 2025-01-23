package gui

import (
	"strings"

	"go-gui/pkg/config"

	"fyne.io/fyne/v2/widget"
	hook "github.com/robotn/gohook"
)

func RegisterHotkeys(startBtn *widget.Button) {
	hotkey := strings.Split(config.Settings.HotKey, "+")
	hook.Register(hook.KeyDown, hotkey, func(e hook.Event) {
		startBtn.OnTapped()
	})

	s := hook.Start()
	<-hook.Process(s)
}
