package audio

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/html"
	"github.com/matthewmueller/joy/dom/mediastreams"
)

var _ AudioContext = (*OfflineAudioContext)(nil)
var _ event.EventTarget = (*OfflineAudioContext)(nil)

func New(numberofchannels uint, length uint, samplerate float32) *OfflineAudioContext {
	macro.Rewrite("new OfflineAudioContext($1, $2, $3)", numberofchannels, length, samplerate)
	return &OfflineAudioContext{}
}

type OfflineAudioContext struct {
}

func (*OfflineAudioContext) StartRendering() (a *AudioBuffer) {
	macro.Rewrite("await $_.startRendering()")
	return a
}

func (*OfflineAudioContext) Suspend(suspendTime float32) {
	macro.Rewrite("await $_.suspend($1)", suspendTime)
}

func (*OfflineAudioContext) Close() {
	macro.Rewrite("await $_.close()")
}

func (*OfflineAudioContext) CreateAnalyser() (a *AnalyserNode) {
	macro.Rewrite("$_.createAnalyser()")
	return a
}

func (*OfflineAudioContext) CreateBiquadFilter() (a *BiquadFilterNode) {
	macro.Rewrite("$_.createBiquadFilter()")
	return a
}

func (*OfflineAudioContext) CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *AudioBuffer) {
	macro.Rewrite("$_.createBuffer($1, $2, $3)", numberOfChannels, length, sampleRate)
	return a
}

func (*OfflineAudioContext) CreateBufferSource() (a *AudioBufferSourceNode) {
	macro.Rewrite("$_.createBufferSource()")
	return a
}

func (*OfflineAudioContext) CreateChannelMerger(numberOfInputs *uint) (a *ChannelMergerNode) {
	macro.Rewrite("$_.createChannelMerger($1)", numberOfInputs)
	return a
}

func (*OfflineAudioContext) CreateChannelSplitter(numberOfOutputs *uint) (a *ChannelSplitterNode) {
	macro.Rewrite("$_.createChannelSplitter($1)", numberOfOutputs)
	return a
}

func (*OfflineAudioContext) CreateConvolver() (a *ConvolverNode) {
	macro.Rewrite("$_.createConvolver()")
	return a
}

func (*OfflineAudioContext) CreateDelay(maxDelayTime *float32) (a *DelayNode) {
	macro.Rewrite("$_.createDelay($1)", maxDelayTime)
	return a
}

func (*OfflineAudioContext) CreateDynamicsCompressor() (a *DynamicsCompressorNode) {
	macro.Rewrite("$_.createDynamicsCompressor()")
	return a
}

func (*OfflineAudioContext) CreateGain() (a *GainNode) {
	macro.Rewrite("$_.createGain()")
	return a
}

func (*OfflineAudioContext) CreateIIRFilter(feedforward []float32, feedback []float32) (a *IIRFilterNode) {
	macro.Rewrite("$_.createIIRFilter($1, $2)", feedforward, feedback)
	return a
}

func (*OfflineAudioContext) CreateMediaElementSource(mediaElement html.HTMLMediaElement) (a *MediaElementAudioSourceNode) {
	macro.Rewrite("$_.createMediaElementSource($1)", mediaElement)
	return a
}

func (*OfflineAudioContext) CreateMediaStreamSource(mediaStream *mediastreams.MediaStream) (a *MediaStreamAudioSourceNode) {
	macro.Rewrite("$_.createMediaStreamSource($1)", mediaStream)
	return a
}

func (*OfflineAudioContext) CreateOscillator() (a *OscillatorNode) {
	macro.Rewrite("$_.createOscillator()")
	return a
}

func (*OfflineAudioContext) CreatePanner() (a *PannerNode) {
	macro.Rewrite("$_.createPanner()")
	return a
}

func (*OfflineAudioContext) CreatePeriodicWave(rl []float32, img []float32, constraints *PeriodicWaveConstraints) (p *PeriodicWave) {
	macro.Rewrite("$_.createPeriodicWave($1, $2, $3)", rl, img, constraints)
	return p
}

func (*OfflineAudioContext) CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (a *ScriptProcessorNode) {
	macro.Rewrite("$_.createScriptProcessor($1, $2, $3)", bufferSize, numberOfInputChannels, numberOfOutputChannels)
	return a
}

func (*OfflineAudioContext) CreateStereoPanner() (a *StereoPannerNode) {
	macro.Rewrite("$_.createStereoPanner()")
	return a
}

func (*OfflineAudioContext) CreateWaveShaper() (a *WaveShaperNode) {
	macro.Rewrite("$_.createWaveShaper()")
	return a
}

func (*OfflineAudioContext) DecodeAudioData(audioData []byte, successCallback *func(decodedData *AudioBuffer), errorCallback *func()) (a *AudioBuffer) {
	macro.Rewrite("await $_.decodeAudioData($1, $2, $3)", audioData, successCallback, errorCallback)
	return a
}

func (*OfflineAudioContext) Resume() {
	macro.Rewrite("await $_.resume()")
}

func (*OfflineAudioContext) AddEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.addEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*OfflineAudioContext) DispatchEvent(evt event.Event) (b bool) {
	macro.Rewrite("$_.dispatchEvent($1)", evt)
	return b
}

func (*OfflineAudioContext) RemoveEventListener(kind string, listener func(evt event.Event), useCapture bool) {
	macro.Rewrite("$_.removeEventListener($1, $2, $3)", kind, listener, useCapture)
}

func (*OfflineAudioContext) Length() (length uint) {
	macro.Rewrite("$_.length")
	return length
}

func (*OfflineAudioContext) Oncomplete() (oncomplete func(*OfflineAudioCompletionEvent)) {
	macro.Rewrite("$_.oncomplete")
	return oncomplete
}

func (*OfflineAudioContext) SetOncomplete(oncomplete func(*OfflineAudioCompletionEvent)) {
	macro.Rewrite("$_.oncomplete = $1", oncomplete)
}

func (*OfflineAudioContext) CurrentTime() (currentTime float32) {
	macro.Rewrite("$_.currentTime")
	return currentTime
}

func (*OfflineAudioContext) Destination() (destination *AudioDestinationNode) {
	macro.Rewrite("$_.destination")
	return destination
}

func (*OfflineAudioContext) Listener() (listener *AudioListener) {
	macro.Rewrite("$_.listener")
	return listener
}

func (*OfflineAudioContext) Onstatechange() (onstatechange func(event.Event)) {
	macro.Rewrite("$_.onstatechange")
	return onstatechange
}

func (*OfflineAudioContext) SetOnstatechange(onstatechange func(event.Event)) {
	macro.Rewrite("$_.onstatechange = $1", onstatechange)
}

func (*OfflineAudioContext) SampleRate() (sampleRate float32) {
	macro.Rewrite("$_.sampleRate")
	return sampleRate
}

func (*OfflineAudioContext) State() (state *AudioContextState) {
	macro.Rewrite("$_.state")
	return state
}
