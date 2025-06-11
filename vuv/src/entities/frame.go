package main

type Frame struct {
	x_resolution int
	y_resolution int
	pixels [][]Pixel
}

func newFrame(x_resolution int, y_resolution int, pixels [][]Pixel) *Frame {
	return &Frame{
		x_resolution: x_resolution,
		y_resolution: y_resolution,
		pixels: pixels,
	}
}