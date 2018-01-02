package utils

import "github.com/matthewmueller/joy/macro"

type PerfWidgetExternal struct {
}

func (*PerfWidgetExternal) AddEventListener(eventType string, callback func()) {
	macro.Rewrite("$_.addEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) GetMemoryUsage() (u uint) {
	macro.Rewrite("$_.getMemoryUsage()")
	return u
}

func (*PerfWidgetExternal) GetProcessCPUUsage() (u uint) {
	macro.Rewrite("$_.getProcessCpuUsage()")
	return u
}

func (*PerfWidgetExternal) GetRecentCPUUsage(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentCpuUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentFrames(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentFrames($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentMemoryUsage(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentMemoryUsage($1)", last)
	return i
}

func (*PerfWidgetExternal) GetRecentPaintRequests(last uint64) (i interface{}) {
	macro.Rewrite("$_.getRecentPaintRequests($1)", last)
	return i
}

func (*PerfWidgetExternal) RemoveEventListener(eventType string, callback func()) {
	macro.Rewrite("$_.removeEventListener($1, $2)", eventType, callback)
}

func (*PerfWidgetExternal) RepositionWindow(x int, y int) {
	macro.Rewrite("$_.repositionWindow($1, $2)", x, y)
}

func (*PerfWidgetExternal) ResizeWindow(width uint, height uint) {
	macro.Rewrite("$_.resizeWindow($1, $2)", width, height)
}

func (*PerfWidgetExternal) ActiveNetworkRequestCount() (activeNetworkRequestCount uint) {
	macro.Rewrite("$_.activeNetworkRequestCount")
	return activeNetworkRequestCount
}

func (*PerfWidgetExternal) AverageFrameTime() (averageFrameTime float32) {
	macro.Rewrite("$_.averageFrameTime")
	return averageFrameTime
}

func (*PerfWidgetExternal) AveragePaintTime() (averagePaintTime float32) {
	macro.Rewrite("$_.averagePaintTime")
	return averagePaintTime
}

func (*PerfWidgetExternal) ExtraInformationEnabled() (extraInformationEnabled bool) {
	macro.Rewrite("$_.extraInformationEnabled")
	return extraInformationEnabled
}

func (*PerfWidgetExternal) IndependentRenderingEnabled() (independentRenderingEnabled bool) {
	macro.Rewrite("$_.independentRenderingEnabled")
	return independentRenderingEnabled
}

func (*PerfWidgetExternal) IrDisablingContentString() (irDisablingContentString string) {
	macro.Rewrite("$_.irDisablingContentString")
	return irDisablingContentString
}

func (*PerfWidgetExternal) IrStatusAvailable() (irStatusAvailable bool) {
	macro.Rewrite("$_.irStatusAvailable")
	return irStatusAvailable
}

func (*PerfWidgetExternal) MaxCPUSpeed() (maxCpuSpeed uint) {
	macro.Rewrite("$_.maxCpuSpeed")
	return maxCpuSpeed
}

func (*PerfWidgetExternal) PaintRequestsPerSecond() (paintRequestsPerSecond uint) {
	macro.Rewrite("$_.paintRequestsPerSecond")
	return paintRequestsPerSecond
}

func (*PerfWidgetExternal) PerformanceCounter() (performanceCounter uint64) {
	macro.Rewrite("$_.performanceCounter")
	return performanceCounter
}

func (*PerfWidgetExternal) PerformanceCounterFrequency() (performanceCounterFrequency uint64) {
	macro.Rewrite("$_.performanceCounterFrequency")
	return performanceCounterFrequency
}
