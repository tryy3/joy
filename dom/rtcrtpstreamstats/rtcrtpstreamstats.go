package rtcrtpstreamstats

import "github.com/matthewmueller/joy/dom/webrtc"

type RTCRTPStreamStats struct {
	*webrtc.RTCStats

	associateStatsId *string
	codecId          *string
	firCount         *uint
	isRemote         *bool
	mediaTrackId     *string
	nackCount        *uint
	pliCount         *uint
	sliCount         *uint
	ssrc             *string
	transportId      *string
}
