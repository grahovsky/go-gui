package utils

import (
	"fmt"
	"log/slog"
	"strconv"
	"time"

	"go-gui/pkg/config"
	"go-gui/pkg/models"

	"github.com/go-vgo/robotgo"
	"golang.org/x/exp/rand"
)

func RunLogic(keySet *models.KeyConfig) error {
	return SendKeyToWindow(keySet)
}

func SendKeyToWindow(keySet *models.KeyConfig) error {
	key := keySet.Value
	delay := keySet.Delay
	clip := keySet.Clip

	delayInt, err1 := strconv.Atoi(delay)
	if err1 != nil {
		return err1
	}
	clipInt, err2 := strconv.Atoi(clip)
	if err2 != nil {
		return err2
	}

	if err := robotgo.ActivePid(config.Settings.Window.ActivePid); err != nil {
		return err
	}

	randomDelay(delayInt)
	if !clipHit(clipInt) {
		return nil
	}

	robotgo.KeySleep = 39 + rand.Intn(15)
	robotgo.MouseSleep = 19 + rand.Intn(10)

	if key == "left" {
		robotgo.Click()
	} else if key == "right" {
		robotgo.Click("right")
	} else {
		if err := robotgo.KeyTap(key); err != nil {
			return err
		}
	}

	slog.Info("Key '%s' sent to window '%s' '%s' \n", key, config.Settings.Window.Name, time.Now())

	return nil
}

func GetActivePid() error {
	windowName := config.Settings.Window.Name
	pids, err := robotgo.FindIds(windowName)
	if err != nil || len(pids) == 0 {
		errf := fmt.Errorf("Error: window '%s' not found\n", windowName)
		return errf
	}
	config.Settings.Window.ActivePid = pids[0]

	return nil
}

func clipHit(chance int) bool {
	r := rand.Intn(chance) + 1
	return chance == r
}

func randomDelay(duration int) {
	// Рандомная задержка перед отправкой
	delayBase := time.Duration(duration*4/5) * time.Millisecond
	delayRand := time.Duration(rand.Intn(duration/5)) * time.Millisecond

	time.Sleep(delayBase + delayRand)
}
