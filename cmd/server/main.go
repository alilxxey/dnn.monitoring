package main

import (
	"net/http"
)

var DataBase MemStorage

func main() {
	DataBase.Fill()
	if err := run(); err != nil {
		panic(err)
	}
}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/gauge/", getGaugeMetric)
	mux.HandleFunc("/update/counter/", getCounterMetric)
	err := http.ListenAndServe("localhost:8080", mux)
	return err

}
