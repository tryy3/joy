package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
)

var _ AudioNode = (*PannerNode)(nil)
var _ event.EventTarget = (*PannerNode)(nil)

type PannerNode struct {
}

func (*PannerNode) SetOrientation(x float32, y float32, z float32) {
	macro.Rewrite("$_.setOrientation($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetPosition(x float32, y float32, z float32) {
	macro.Rewrite("$_.setPosition($1, $2, $3)", x, y, z)
}

func (*PannerNode) SetVelocity(x float32, y float32, z float32) {
	macro.Rewrite("$_.setVelocity($1, $2, $3)", x, y, z)
}

func (*PannerNode) Connect(destination AudioNode, output *uint, input *uint) (a AudioNode) {
	macro.Rewrite("$_.connect($1, $2, $3)", destination, output, input)
	return a
}

func (*PannerNode) Disconnect() {
	macro.Rewrite("$_.disconnect()")
}

func (*PannerNode) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*PannerNode) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*PannerNode) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*PannerNode) ConeInnerAngle() (coneInnerAngle float32) {
	macro.Rewrite("$_.coneInnerAngle")
	return coneInnerAngle
}

func (*PannerNode) SetConeInnerAngle(coneInnerAngle float32) {
	macro.Rewrite("$_.coneInnerAngle = $1", coneInnerAngle)
}

func (*PannerNode) ConeOuterAngle() (coneOuterAngle float32) {
	macro.Rewrite("$_.coneOuterAngle")
	return coneOuterAngle
}

func (*PannerNode) SetConeOuterAngle(coneOuterAngle float32) {
	macro.Rewrite("$_.coneOuterAngle = $1", coneOuterAngle)
}

func (*PannerNode) ConeOuterGain() (coneOuterGain float32) {
	macro.Rewrite("$_.coneOuterGain")
	return coneOuterGain
}

func (*PannerNode) SetConeOuterGain(coneOuterGain float32) {
	macro.Rewrite("$_.coneOuterGain = $1", coneOuterGain)
}

func (*PannerNode) DistanceModel() (distanceModel *DistanceModelType) {
	macro.Rewrite("$_.distanceModel")
	return distanceModel
}

func (*PannerNode) SetDistanceModel(distanceModel *DistanceModelType) {
	macro.Rewrite("$_.distanceModel = $1", distanceModel)
}

func (*PannerNode) MaxDistance() (maxDistance float32) {
	macro.Rewrite("$_.maxDistance")
	return maxDistance
}

func (*PannerNode) SetMaxDistance(maxDistance float32) {
	macro.Rewrite("$_.maxDistance = $1", maxDistance)
}

func (*PannerNode) PanningModel() (panningModel *PanningModelType) {
	macro.Rewrite("$_.panningModel")
	return panningModel
}

func (*PannerNode) SetPanningModel(panningModel *PanningModelType) {
	macro.Rewrite("$_.panningModel = $1", panningModel)
}

func (*PannerNode) RefDistance() (refDistance float32) {
	macro.Rewrite("$_.refDistance")
	return refDistance
}

func (*PannerNode) SetRefDistance(refDistance float32) {
	macro.Rewrite("$_.refDistance = $1", refDistance)
}

func (*PannerNode) RolloffFactor() (rolloffFactor float32) {
	macro.Rewrite("$_.rolloffFactor")
	return rolloffFactor
}

func (*PannerNode) SetRolloffFactor(rolloffFactor float32) {
	macro.Rewrite("$_.rolloffFactor = $1", rolloffFactor)
}

func (*PannerNode) ChannelCount() (channelCount uint) {
	macro.Rewrite("$_.channelCount")
	return channelCount
}

func (*PannerNode) SetChannelCount(channelCount uint) {
	macro.Rewrite("$_.channelCount = $1", channelCount)
}

func (*PannerNode) ChannelCountMode() (channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode")
	return channelCountMode
}

func (*PannerNode) SetChannelCountMode(channelCountMode *ChannelCountMode) {
	macro.Rewrite("$_.channelCountMode = $1", channelCountMode)
}

func (*PannerNode) ChannelInterpretation() (channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation")
	return channelInterpretation
}

func (*PannerNode) SetChannelInterpretation(channelInterpretation *ChannelInterpretation) {
	macro.Rewrite("$_.channelInterpretation = $1", channelInterpretation)
}

func (*PannerNode) Context() (context AudioContext) {
	macro.Rewrite("$_.context")
	return context
}

func (*PannerNode) NumberOfInputs() (numberOfInputs uint) {
	macro.Rewrite("$_.numberOfInputs")
	return numberOfInputs
}

func (*PannerNode) NumberOfOutputs() (numberOfOutputs uint) {
	macro.Rewrite("$_.numberOfOutputs")
	return numberOfOutputs
}
