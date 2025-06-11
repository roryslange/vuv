package main

type Metadata struct {
	x_resolution int
	y_resolution int
	frame_rate	float32
	frames_count int
}

func newMetadata(x_resolution int, y_resolution int, frame_rate float32, frames_count int) *Metadata {
	return &Metadata{
		x_resolution: x_resolution,
		y_resolution: y_resolution,
		frame_rate: frame_rate,
		frames_count: frames_count,
	}
}