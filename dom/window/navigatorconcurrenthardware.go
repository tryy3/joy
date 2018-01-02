package window

type NavigatorConcurrentHardware interface {
	HardwareConcurrency() (hardwareConcurrency uint64)
}
