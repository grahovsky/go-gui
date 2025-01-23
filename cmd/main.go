package main

import (
	"go-gui/pkg/config"
	"go-gui/pkg/gui"
)

func main() {
	config.Init()
	gui.RunApp()
}
