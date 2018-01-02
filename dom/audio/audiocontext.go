package audio

import (
	"github.com/matthewmueller/joy/dom/event"
	"github.com/matthewmueller/joy/dom/mediastreams"
	"github.com/matthewmueller/joy/dom/html"
)

type AudioContext interface {
	event.
		EventTarget

	Close()

	CreateAnalyser() (a *AnalyserNode)

	CreateBiquadFilter() (b *BiquadFilterNode)

	CreateBuffer(numberOfChannels uint, length uint, sampleRate float32) (a *AudioBuffer)

	CreateBufferSource() (a *AudioBufferSourceNode)

	CreateChannelMerger(numberOfInputs *uint) (c *ChannelMergerNode)

	CreateChannelSplitter(numberOfOutputs *uint) (c *ChannelSplitterNode)

	CreateConvolver() (c *ConvolverNode)

	CreateDelay(maxDelayTime *float32) (d *DelayNode)

	CreateDynamicsCompressor() (d *DynamicsCompressorNode)

	CreateGain() (g *GainNode)

	CreateIIRFilter(feedforward []float32, feedback []float32) (i *IIRFilterNode)

	CreateMediaElementSource(mediaElement html.HTMLMediaElement) (m *MediaElementAudioSourceNode)

	CreateMediaStreamSource(mediaStream *mediastreams.MediaStream) (m *MediaStreamAudioSourceNode)

	CreateOscillator() (o *OscillatorNode)

	CreatePanner() (p *PannerNode)

	CreatePeriodicWave(rl []float32, img []float32, constraints *PeriodicWaveConstraints) (p *PeriodicWave)

	CreateScriptProcessor(bufferSize *uint, numberOfInputChannels *uint, numberOfOutputChannels *uint) (s *ScriptProcessorNode)

	CreateStereoPanner() (s *StereoPannerNode)

	CreateWaveShaper() (w *WaveShaperNode)

	DecodeAudioData(audioData []byte, successCallback *func(decodedData *AudioBuffer), errorCallback *func()) (a *AudioBuffer)

	Resume()

	Suspend(suspendTime float32)

	CurrentTime() (currentTime float32)

	Destination() (destination *AudioDestinationNode)

	Listener() (listener *AudioListener)

	Onstatechange() (onstatechange func(event.Event))

	SetOnstatechange(onstatechange func(event.Event))

	SampleRate() (sampleRate float32)

	State() (state *AudioContextState)
}
