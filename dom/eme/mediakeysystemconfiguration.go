package eme

type MediaKeySystemConfiguration struct {
	audioCapabilities     *[]*MediaKeySystemMediaCapability
	distinctiveIdentifier *MediaKeysRequirement
	initDataTypes         *[]string
	persistentState       *MediaKeysRequirement
	videoCapabilities     *[]*MediaKeySystemMediaCapability
}
