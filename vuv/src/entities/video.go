package main

type Video struct {
	metadata Metadata
	frames []Frame
}

func newVideo(metadata Metadata, frames []Frame) *Video {
	return &Video{
		metadata: metadata,
		frames: frames,
	}
}
