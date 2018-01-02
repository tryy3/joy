package window

type WindowTimersExtension interface {
	ClearImmediate(handle int)

	SetImmediate(expression interface{}, args interface{}) (i int)
}
