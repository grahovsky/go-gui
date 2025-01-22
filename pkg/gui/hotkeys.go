package gui

import (
	"fyne.io/fyne/v2/widget"
	hook "github.com/robotn/gohook"
)

func RegisterHotkeys(startBtn *widget.Button) {
	hook.Register(hook.KeyDown, []string{"alt", "s"}, func(e hook.Event) {
		startBtn.OnTapped()
	})

	s := hook.Start()
	<-hook.Process(s)
}
