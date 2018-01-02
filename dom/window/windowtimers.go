package window

type WindowTimers interface {
	ClearImmediate(handle int)

	SetImmediate(expression interface{}, args interface{}) (i int)

	ClearInterval(handle int)

	ClearTimeout(handle int)

	SetInterval(handler interface{}, timeout *interface{}, args interface{}) (i int)

	SetTimeout(handler interface{}, timeout *interface{}, args interface{}) (i int)
}
