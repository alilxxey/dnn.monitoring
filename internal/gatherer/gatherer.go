package gatherer

import (
	"math/rand"
	"runtime"
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

type GaugeRecord struct {
	Alloc         float64
	BuckHashSys   float64
	Frees         float64
	GCCPUFraction float64
	GCSys         float64
	HeapAlloc     float64
	HeapIdle      float64
	HeapInuse     float64
	HeapObjects   float64
	HeapReleased  float64
	HeapSys       float64
	LastGC        float64
	Lookups       float64
	MCacheInuse   float64
	MCacheSys     float64
	MSpanInuse    float64
	MSpanSys      float64
	Mallocs       float64
	NextGC        float64
	NumForcedGC   float64
	NumGC         float64
	OtherSys      float64
	PauseTotalNs  float64
	StackInuse    float64
	StackSys      float64
	Sys           float64
	TotalAlloc    float64
	RandomValue   float64
}
type CounterRecord struct {
	PollCount int64
}

type MetricsRecord struct {
	gauge   GaugeRecord
	counter CounterRecord
}

func (m *MetricsRecord) ReadStats() {
	var ms runtime.MemStats
	runtime.ReadMemStats(&ms)
	m.gauge = GaugeRecord{
		Alloc:         float64(ms.Alloc),
		BuckHashSys:   float64(ms.BuckHashSys),
		Frees:         float64(ms.Frees),
		GCCPUFraction: ms.GCCPUFraction,
		GCSys:         float64(ms.GCSys),
		HeapAlloc:     float64(ms.HeapAlloc),
		HeapIdle:      float64(ms.HeapIdle),
		HeapInuse:     float64(ms.HeapInuse),
		HeapObjects:   float64(ms.HeapObjects),
		HeapReleased:  float64(ms.HeapReleased),
		HeapSys:       float64(ms.HeapSys),
		LastGC:        float64(ms.LastGC),
		Lookups:       float64(ms.Lookups),
		MCacheInuse:   float64(ms.MCacheInuse),
		MCacheSys:     float64(ms.MCacheSys),
		MSpanInuse:    float64(ms.MSpanInuse),
		MSpanSys:      float64(ms.MSpanSys),
		Mallocs:       float64(ms.Mallocs),
		NextGC:        float64(ms.NextGC),
		NumForcedGC:   float64(ms.NumForcedGC),
		NumGC:         float64(ms.NumGC),
		OtherSys:      float64(ms.OtherSys),
		PauseTotalNs:  float64(ms.PauseTotalNs),
		StackInuse:    float64(ms.StackInuse),
		StackSys:      float64(ms.StackSys),
		Sys:           float64(ms.Sys),
		TotalAlloc:    float64(ms.TotalAlloc),
		RandomValue:   rand.Float64(),
	}
	m.counter = CounterRecord{
		PollCount: 1,
	}
}
