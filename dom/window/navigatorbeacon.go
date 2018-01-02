package window

type NavigatorBeacon interface {
	SendBeacon(url string, data *interface{}) (b bool)
}
