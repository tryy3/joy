package speech

import "github.com/matthewmueller/joy/macro"

type SpeechSynthesisVoice struct {
}

func (*SpeechSynthesisVoice) Default() (def bool) {
	macro.Rewrite("$_.default")
	return def
}

func (*SpeechSynthesisVoice) Lang() (lang string) {
	macro.Rewrite("$_.lang")
	return lang
}

func (*SpeechSynthesisVoice) LocalService() (localService bool) {
	macro.Rewrite("$_.localService")
	return localService
}

func (*SpeechSynthesisVoice) Name() (name string) {
	macro.Rewrite("$_.name")
	return name
}

func (*SpeechSynthesisVoice) VoiceURI() (voiceURI string) {
	macro.Rewrite("$_.voiceURI")
	return voiceURI
}
