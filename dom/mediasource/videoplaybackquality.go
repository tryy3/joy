package mediasource

import "github.com/matthewmueller/joy/macro"

type VideoPlaybackQuality struct {
}

func (*VideoPlaybackQuality) CorruptedVideoFrames() (corruptedVideoFrames uint) {
	macro.Rewrite("$_.corruptedVideoFrames")
	return corruptedVideoFrames
}

func (*VideoPlaybackQuality) CreationTime() (creationTime int) {
	macro.Rewrite("$_.creationTime")
	return creationTime
}

func (*VideoPlaybackQuality) DroppedVideoFrames() (droppedVideoFrames uint) {
	macro.Rewrite("$_.droppedVideoFrames")
	return droppedVideoFrames
}

func (*VideoPlaybackQuality) TotalFrameDelay() (totalFrameDelay float32) {
	macro.Rewrite("$_.totalFrameDelay")
	return totalFrameDelay
}

func (*VideoPlaybackQuality) TotalVideoFrames() (totalVideoFrames uint) {
	macro.Rewrite("$_.totalVideoFrames")
	return totalVideoFrames
}
