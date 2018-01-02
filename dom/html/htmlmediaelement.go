package html

import (
	"github.com/matthewmueller/joy/dom/webvvt"
	"github.com/matthewmueller/joy/dom/ms"
	"github.com/matthewmueller/joy/dom/eme"
	"github.com/matthewmueller/joy/dom/audio"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/mediastreams"
	"github.com/matthewmueller/joy/dom/dom"
)

type HTMLMediaElement interface {
	HTMLElement

	AddTextTrack(kind string, label *string, language *string) (t *webvvt.TextTrack)

	CanPlayType(kind string) (s string)

	Load()

	MsClearEffects()

	MsGetAsCastingSource() (i interface{})

	MsInsertAudioEffect(activatableClassId string, effectRequired bool, config *interface{})

	MsSetMediaKeys(mediaKeys *ms.MSMediaKeys)

	MsSetMediaProtectionManager(mediaProtectionManager *interface{})

	Pause()

	Play()

	SetMediaKeys(mediaKeys *eme.MediaKeys)

	AudioTracks() (audioTracks *audio.AudioTrackList)

	Autoplay() (autoplay bool)

	SetAutoplay(autoplay bool)

	Buffered() (buffered *dom.TimeRanges)

	Controls() (controls bool)

	SetControls(controls bool)

	CrossOrigin() (crossOrigin string)

	SetCrossOrigin(crossOrigin string)

	CurrentSrc() (currentSrc string)

	CurrentTime() (currentTime float32)

	SetCurrentTime(currentTime float32)

	DefaultMuted() (defaultMuted bool)

	SetDefaultMuted(defaultMuted bool)

	DefaultPlaybackRate() (defaultPlaybackRate float32)

	SetDefaultPlaybackRate(defaultPlaybackRate float32)

	Duration() (duration float32)

	Ended() (ended bool)

	Error() (err *utils.MediaError)

	Loop() (loop bool)

	SetLoop(loop bool)

	MediaKeys() (mediaKeys *eme.MediaKeys)

	MsAudioCategory() (msAudioCategory string)

	SetMsAudioCategory(msAudioCategory string)

	MsAudioDeviceType() (msAudioDeviceType string)

	SetMsAudioDeviceType(msAudioDeviceType string)

	MsGraphicsTrustStatus() (msGraphicsTrustStatus *ms.MSGraphicsTrust)

	MsKeys() (msKeys *ms.MSMediaKeys)

	MsPlayToDisabled() (msPlayToDisabled bool)

	SetMsPlayToDisabled(msPlayToDisabled bool)

	MsPlayToPreferredSourceURI() (msPlayToPreferredSourceUri string)

	SetMsPlayToPreferredSourceURI(msPlayToPreferredSourceUri string)

	MsPlayToPrimary() (msPlayToPrimary bool)

	SetMsPlayToPrimary(msPlayToPrimary bool)

	MsPlayToSource() (msPlayToSource interface{})

	MsRealTime() (msRealTime bool)

	SetMsRealTime(msRealTime bool)

	Muted() (muted bool)

	SetMuted(muted bool)

	NetworkState() (networkState uint8)

	Onencrypted() (onencrypted func(*dom.MediaEncryptedEvent))

	SetOnencrypted(onencrypted func(*dom.MediaEncryptedEvent))

	Onmsneedkey() (onmsneedkey func(*ms.MSMediaKeyNeededEvent))

	SetOnmsneedkey(onmsneedkey func(*ms.MSMediaKeyNeededEvent))

	Paused() (paused bool)

	PlaybackRate() (playbackRate float32)

	SetPlaybackRate(playbackRate float32)

	Played() (played *dom.TimeRanges)

	Preload() (preload string)

	SetPreload(preload string)

	ReadyState() (readyState interface{})

	Seekable() (seekable *dom.TimeRanges)

	Seeking() (seeking bool)

	Src() (src string)

	SetSrc(src string)

	SrcObject() (srcObject *mediastreams.MediaStream)

	SetSrcObject(srcObject *mediastreams.MediaStream)

	TextTracks() (textTracks *webvvt.TextTrackList)

	VideoTracks() (videoTracks *audio.VideoTrackList)

	Volume() (volume float32)

	SetVolume(volume float32)
}
