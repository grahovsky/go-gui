package utils

import (
	"fmt"
	"strconv"
	"time"

	"go-gui/pkg/models"

	"github.com/go-vgo/robotgo"
	"golang.org/x/exp/rand"
)

func RunLogic() {
	SendKeyToWindow(models.CurrentInput.Value3, models.CurrentInput.Value2)
}

func SendKeyToWindow(windowName, key string) {
	pid, err := robotgo.FindIds(windowName)
	if err != nil || len(pid) == 0 {
		fmt.Printf("Error: window '%s' not found\n", windowName)
		return
	}

	robotgo.ActivePid(pid[0])

	value1, err1 := strconv.Atoi(models.CurrentInput.Value1)

	if err1 != nil {
		fmt.Println("Ошибка: некорректный ввод")
		return
	}

	randomDelay(value1)

	if key == "left" {
		robotgo.Click()
	} else if key == "right" {
		robotgo.Click("right")
	} else {
		robotgo.TypeStr(key)
	}

	fmt.Printf("Key '%s' sent to window '%s' '%s' \n", key, windowName, time.Now())
}

func randomDelay(duration int) {
	// Рандомная задержка перед отправкой
	delayBase := time.Duration(duration*4/5) * time.Millisecond
	delayRand := time.Duration(rand.Intn(duration/5)) * time.Millisecond

	time.Sleep(delayBase + delayRand)
}
