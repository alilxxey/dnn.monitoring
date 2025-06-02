package storage

import (
	"errors"
	"fmt"
)

const (
	Alloc         = "Alloc"
	BuckHashSys   = "BuckHashSys"
	Frees         = "Frees"
	GCCPUFraction = "GCCPUFraction"
	GCSys         = "GCSys"
	HeapAlloc     = "HeapAlloc"
	HeapIdle      = "HeapIdle"
	HeapInuse     = "HeapInuse"
	HeapObjects   = "HeapObjects"
	HeapReleased  = "HeapReleased"
	HeapSys       = "HeapSys"
	LastGC        = "LastGC"
	Lookups       = "Lookups"
	MCacheInuse   = "MCacheInuse"
	MCacheSys     = "MCacheSys"
	MSpanInuse    = "MSpanInuse"
	MSpanSys      = "MSpanSys"
	Mallocs       = "Mallocs"
	NextGC        = "NextGC"
	NumForcedGC   = "NumForcedGC"
	NumGC         = "NumGC"
	OtherSys      = "OtherSys"
	PauseTotalNs  = "PauseTotalNs"
	StackInuse    = "StackInuse"
	StackSys      = "StackSys"
	Sys           = "Sys"
	TotalAlloc    = "TotalAlloc"
	RandomValue   = "RandomValue"
	PollCount     = "PollCount"
	gauge         = "gauge"
	counter       = "counter"
)

var GaugeMetricNames = []string{
	Alloc,
	BuckHashSys,
	Frees,
	GCCPUFraction,
	GCSys,
	HeapAlloc,
	HeapIdle,
	HeapInuse,
	HeapObjects,
	HeapReleased,
	HeapSys,
	LastGC,
	Lookups,
	MCacheInuse,
	MCacheSys,
	MSpanInuse,
	MSpanSys,
	Mallocs,
	NextGC,
	NumForcedGC,
	NumGC,
	OtherSys,
	PauseTotalNs,
	StackInuse,
	StackSys,
	Sys,
	TotalAlloc,
	RandomValue,
}
var CounterMetricNames = []string{
	PollCount,
}

type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func New() *MemStorage {
	s := MemStorage{
		gauge:   make(map[string]float64),
		counter: make(map[string]int64),
	}
	return &s
}

func (s *MemStorage) Exists(name string) (string, error) {
	_, ok := s.gauge[name]
	if ok {
		return "gauge", nil
	}
	_, ok1 := s.counter[name]
	if ok1 {
		return "counter", nil
	}
	return "", errors.New("metric doesn't exist")
}

func (s *MemStorage) WriteGauge(name string, value float64) error {
	s.gauge[name] = value
	return nil
}

func (s *MemStorage) Increment(name string, value int64) error {
	s.counter[name] += value
	return nil
}

func (s *MemStorage) GetValue(name string) (any, error) {
	t, err := s.Exists(name)
	if err != nil {
		return "", err
	}
	switch t {
	case gauge:
		fmt.Println(s.gauge[name])
		return s.gauge[name], nil
	case counter:
		fmt.Println(s.counter[name])
		return s.counter[name], nil
	}
	return "", errors.New("something went wrong")
}
