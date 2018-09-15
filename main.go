package main

import (
	"time"

	"github.com/go-vgo/robotgo"
)

func main() {

	for {
		screenWidth, _ := robotgo.GetScreenSize()

		mouseX, mouseY := robotgo.GetMousePos()

		if mouseX+10 > screenWidth {
			robotgo.MoveMouse(11, mouseY)
		} else if mouseX-10 < 0 {
			robotgo.MoveMouse(screenWidth-11, mouseY)
		}

		time.Sleep(50 * time.Millisecond)
	}

}
