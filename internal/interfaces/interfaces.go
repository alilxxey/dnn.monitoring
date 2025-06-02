package interfaces

type DB interface {
	WriteGauge(name string, value float64) error
	Increment(name string, value int64) error
	GetValue(name string) (any, error)
	Exists(name string) (string, error)
}
