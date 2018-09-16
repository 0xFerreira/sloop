package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func checkLoop(screenWidth, mouseX, activationMargin int) string {

	if mouseX+activationMargin > screenWidth {
		return "left"
	} else if mouseX-activationMargin < 0 {
		return "right"
	}

	return ""
}

func main() {

	var (
		activationMargin = 10
		jumpDelta        = 1
		checkFrequency   = 50
	)

	for {
		screenWidth, _ := robotgo.GetScreenSize()
		mouseX, mouseY := robotgo.GetMousePos()

		loop := checkLoop(screenWidth, mouseX, activationMargin)

		switch loop {
		case "left":
			robotgo.MoveMouse(activationMargin+jumpDelta, mouseY)
		case "right":
			robotgo.MoveMouse(screenWidth-(activationMargin+jumpDelta), mouseY)
		}

		time.Sleep(time.Duration(checkFrequency) * time.Millisecond)
	}

}
