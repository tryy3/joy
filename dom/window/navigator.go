package window

import (
	"github.com/matthewmueller/joy/macro"
	"github.com/matthewmueller/joy/dom/gamepad"
	"github.com/matthewmueller/joy/dom/eme"
	"github.com/matthewmueller/joy/dom/navigation"
	"github.com/matthewmueller/joy/dom/utils"
	"github.com/matthewmueller/joy/dom/mediastreams"
	"github.com/matthewmueller/joy/dom/authentication"
	"github.com/matthewmueller/joy/dom/serviceworker"
	"github.com/matthewmueller/joy/dom/geolocation"
)

type Navigator struct {
}

func (*Navigator) GetGamepads() (g []*gamepad.Gamepad) {
	macro.Rewrite("$_.getGamepads()")
	return g
}

func (*Navigator) JavaEnabled() (b bool) {
	macro.Rewrite("$_.javaEnabled()")
	return b
}

func (*Navigator) MsLaunchURI(uri string, successCallback *func(), noHandlerCallback *func()) {
	macro.Rewrite("$_.msLaunchUri($1, $2, $3)", uri, successCallback, noHandlerCallback)
}

func (*Navigator) RequestMediaKeySystemAccess(keySystem string, supportedConfigurations []*eme.MediaKeySystemConfiguration) (m *eme.MediaKeySystemAccess) {
	macro.Rewrite("await $_.requestMediaKeySystemAccess($1, $2)", keySystem, supportedConfigurations)
	return m
}

func (*Navigator) ConfirmSiteSpecificTrackingException(args *navigation.ConfirmSiteSpecificExceptionsInformation) (b bool) {
	macro.Rewrite("$_.confirmSiteSpecificTrackingException($1)", args)
	return b
}

func (*Navigator) ConfirmWebWideTrackingException(args *utils.ExceptionInformation) (b bool) {
	macro.Rewrite("$_.confirmWebWideTrackingException($1)", args)
	return b
}

func (*Navigator) RemoveSiteSpecificTrackingException(args *utils.ExceptionInformation) {
	macro.Rewrite("$_.removeSiteSpecificTrackingException($1)", args)
}

func (*Navigator) RemoveWebWideTrackingException(args *utils.ExceptionInformation) {
	macro.Rewrite("$_.removeWebWideTrackingException($1)", args)
}

func (*Navigator) StoreSiteSpecificTrackingException(args *utils.StoreSiteSpecificExceptionsInformation) {
	macro.Rewrite("$_.storeSiteSpecificTrackingException($1)", args)
}

func (*Navigator) StoreWebWideTrackingException(args *utils.StoreExceptionsInformation) {
	macro.Rewrite("$_.storeWebWideTrackingException($1)", args)
}

func (*Navigator) MsSaveBlob(blob interface{}, defaultName *string) (b bool) {
	macro.Rewrite("$_.msSaveBlob($1, $2)", blob, defaultName)
	return b
}

func (*Navigator) MsSaveOrOpenBlob(blob interface{}, defaultName *string) (b bool) {
	macro.Rewrite("$_.msSaveOrOpenBlob($1, $2)", blob, defaultName)
	return b
}

func (*Navigator) SendBeacon(url string, data *interface{}) (b bool) {
	macro.Rewrite("$_.sendBeacon($1, $2)", url, data)
	return b
}

func (*Navigator) GetUserMedia(constraints *mediastreams.MediaStreamConstraints, successCallback func(stream *mediastreams.MediaStream), errorCallback func(err *mediastreams.MediaStreamError)) {
	macro.Rewrite("$_.getUserMedia($1, $2, $3)", constraints, successCallback, errorCallback)
}

func (*Navigator) Authentication() (authentication *authentication.WebAuthentication) {
	macro.Rewrite("$_.authentication")
	return authentication
}

func (*Navigator) CookieEnabled() (cookieEnabled bool) {
	macro.Rewrite("$_.cookieEnabled")
	return cookieEnabled
}

func (*Navigator) GamepadInputEmulation() (gamepadInputEmulation *gamepad.GamepadInputEmulationType) {
	macro.Rewrite("$_.gamepadInputEmulation")
	return gamepadInputEmulation
}

func (*Navigator) SetGamepadInputEmulation(gamepadInputEmulation *gamepad.GamepadInputEmulationType) {
	macro.Rewrite("$_.gamepadInputEmulation = $1", gamepadInputEmulation)
}

func (*Navigator) Language() (language string) {
	macro.Rewrite("$_.language")
	return language
}

func (*Navigator) MaxTouchPoints() (maxTouchPoints int) {
	macro.Rewrite("$_.maxTouchPoints")
	return maxTouchPoints
}

func (*Navigator) MimeTypes() (mimeTypes *utils.MimeTypeArray) {
	macro.Rewrite("$_.mimeTypes")
	return mimeTypes
}

func (*Navigator) MsManipulationViewsEnabled() (msManipulationViewsEnabled bool) {
	macro.Rewrite("$_.msManipulationViewsEnabled")
	return msManipulationViewsEnabled
}

func (*Navigator) MsMaxTouchPoints() (msMaxTouchPoints int) {
	macro.Rewrite("$_.msMaxTouchPoints")
	return msMaxTouchPoints
}

func (*Navigator) MsPointerEnabled() (msPointerEnabled bool) {
	macro.Rewrite("$_.msPointerEnabled")
	return msPointerEnabled
}

func (*Navigator) Plugins() (plugins *utils.PluginArray) {
	macro.Rewrite("$_.plugins")
	return plugins
}

func (*Navigator) PointerEnabled() (pointerEnabled bool) {
	macro.Rewrite("$_.pointerEnabled")
	return pointerEnabled
}

func (*Navigator) ServiceWorker() (serviceWorker *serviceworker.ServiceWorkerContainer) {
	macro.Rewrite("$_.serviceWorker")
	return serviceWorker
}

func (*Navigator) Webdriver() (webdriver bool) {
	macro.Rewrite("$_.webdriver")
	return webdriver
}

func (*Navigator) AppCodeName() (appCodeName string) {
	macro.Rewrite("$_.appCodeName")
	return appCodeName
}

func (*Navigator) AppName() (appName string) {
	macro.Rewrite("$_.appName")
	return appName
}

func (*Navigator) AppVersion() (appVersion string) {
	macro.Rewrite("$_.appVersion")
	return appVersion
}

func (*Navigator) Platform() (platform string) {
	macro.Rewrite("$_.platform")
	return platform
}

func (*Navigator) Product() (product string) {
	macro.Rewrite("$_.product")
	return product
}

func (*Navigator) ProductSub() (productSub string) {
	macro.Rewrite("$_.productSub")
	return productSub
}

func (*Navigator) UserAgent() (userAgent string) {
	macro.Rewrite("$_.userAgent")
	return userAgent
}

func (*Navigator) Vendor() (vendor string) {
	macro.Rewrite("$_.vendor")
	return vendor
}

func (*Navigator) VendorSub() (vendorSub string) {
	macro.Rewrite("$_.vendorSub")
	return vendorSub
}

func (*Navigator) OnLine() (onLine bool) {
	macro.Rewrite("$_.onLine")
	return onLine
}

func (*Navigator) Geolocation() (geolocation *geolocation.Geolocation) {
	macro.Rewrite("$_.geolocation")
	return geolocation
}

func (*Navigator) HardwareConcurrency() (hardwareConcurrency uint64) {
	macro.Rewrite("$_.hardwareConcurrency")
	return hardwareConcurrency
}

func (*Navigator) MediaDevices() (mediaDevices *mediastreams.MediaDevices) {
	macro.Rewrite("$_.mediaDevices")
	return mediaDevices
}
