package main

import (
	"testing"
)

func Test_checkLoop(t *testing.T) {
	type args struct {
		screenWidth      int
		mouseX           int
		activationMargin int
	}
	tests := []struct {
		name string
		args args
		want string
	}{
		{"Test Left Jump", args{100, 91, 10}, "left"},
		{"Test Right Jump", args{100, 1, 10}, "right"},
		{"Test No Jump", args{100, 50, 10}, ""},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := checkLoop(tt.args.screenWidth, tt.args.mouseX, tt.args.activationMargin); got != tt.want {
				t.Errorf("checkLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_run(t *testing.T) {
	type args struct {
		activationMargin *int
		jumpDelta        *int
	}

	var activationMargin = 0
	var jumpDelta = 1

	tests := []struct {
		name string
		args args
	}{
		{"Test single run", args{&activationMargin, &jumpDelta}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			run(tt.args.activationMargin, tt.args.jumpDelta)
		})
	}
}
