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

func run(activationMargin, jumpDelta *int) {
	screenWidth, _ := robotgo.GetScreenSize()
	mouseX, mouseY := robotgo.GetMousePos()

	loop := checkLoop(screenWidth, mouseX, *activationMargin)

	switch loop {
	case "left":
		robotgo.MoveMouse(*activationMargin+*jumpDelta, mouseY)
	case "right":
		robotgo.MoveMouse(screenWidth-(*activationMargin+*jumpDelta), mouseY)
	}
}

func main() {

	activationMargin := flag.Int("activationMargin", 10, "Jump activation margin")
	jumpDelta := flag.Int("jumpDelta", 1, "Jump delta distance from activation margin")
	hz := flag.Int("hz", 50, "Pointer position check frequency")
	d := flag.Bool("d", false, "Daemonize and return control")

	flag.Parse()

	if *d {
		for {
			run(activationMargin, jumpDelta)
			time.Sleep(time.Duration(*hz) * time.Millisecond)
		}
	} else {
		run(activationMargin, jumpDelta)
	}
}
