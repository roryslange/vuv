package main

type Pixel struct {
	color int
	nextFrameDiff int
}

func newPixel(color int, nextFrameDiff int) *Pixel {
	return &Pixel{
		color: color,
		nextFrameDiff: nextFrameDiff,
	}
}