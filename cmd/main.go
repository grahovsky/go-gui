package main

import (
	"os"

	"go-gui/pkg/gui"
)

func main() {
	os.Setenv("FYNE_RENDERER", "software")

	gui.RunApp()
}
