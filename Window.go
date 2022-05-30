package main

const defaultX = 17
const defaultY = 128
const defaultWidth = 632
const defaultHeight = defaultWidth + 30

// A Window
type Window struct {
	StartX        int
	StartY        int
	Width         int
	Height        int
	widthScaling  float64
	heightScaling float64
}
