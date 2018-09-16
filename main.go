package main

import (
	"flag"
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

func run(activationMargin, jumpDelta, hz *int) {

	for {
		screenWidth, _ := robotgo.GetScreenSize()
		mouseX, mouseY := robotgo.GetMousePos()

		loop := checkLoop(screenWidth, mouseX, *activationMargin)

		switch loop {
		case "left":
			robotgo.MoveMouse(*activationMargin+*jumpDelta, mouseY)
		case "right":
			robotgo.MoveMouse(screenWidth-(*activationMargin+*jumpDelta), mouseY)
		}

		time.Sleep(time.Duration(*hz) * time.Millisecond)
	}

}

func main() {

	activationMargin := flag.Int("activationMargin", 10, "Jump activation margin")
	jumpDelta := flag.Int("jumpDelta", 1, "Jump delta distance from activation margin")
	hz := flag.Int("hz", 50, "Pointer position check frequency")

	flag.Parse()

	run(activationMargin, jumpDelta, hz)
}
