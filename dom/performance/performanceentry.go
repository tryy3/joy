package performance

type PerformanceEntry interface {
	Duration() (duration int)

	EntryType() (entryType string)

	Name() (name string)

	StartTime() (startTime int)
}
