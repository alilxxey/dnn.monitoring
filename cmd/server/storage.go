package main

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

var gaugeMetricNames = []string{
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
var counterMetricNames = []string{
	PollCount,
}

type DB interface {
	WriteGauge(name string, value float64) error
	Increment(name string, value int64) error
	GetValue(name string) (any, error)
	Exists(name string) bool
	Fill()
}
type MemStorage struct {
	gauge   map[string]float64
	counter map[string]int64
}

func (s *MemStorage) Fill() {
	if s.gauge == nil && s.counter == nil {
		s.gauge = make(map[string]float64)
		s.counter = make(map[string]int64)
	} else {
		panic("already exists")
	}
	for _, v := range gaugeMetricNames {
		s.gauge[v] = 0.0
	}
	for _, v := range counterMetricNames {
		s.counter[v] = 0
	}

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
	if _, err := s.Exists(name); err != nil {
		return err
	}
	s.gauge[name] = value
	return nil
}

func (s *MemStorage) Increment(name string, value int64) error {
	if _, err := s.Exists(name); err != nil {
		return err
	}
	s.counter[name] += value
	return nil
}

func (s *MemStorage) GetValue(name string) any {
	t, err := s.Exists(name)
	if err != nil {
		return err
	}
	switch t {
	case gauge:
		fmt.Println(s.gauge[name])
		return s.gauge[name]
	case counter:
		fmt.Println(s.counter[name])
		return s.counter[name]
	}
	return errors.New("something went wrong")
}
