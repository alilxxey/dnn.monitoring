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
func send400(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusBadRequest)

}

func run() error {
	mux := http.NewServeMux()
	mux.HandleFunc("/update/gauge/", getGaugeMetric)
	mux.HandleFunc("/update/counter/", getCounterMetric)
	mux.HandleFunc("/", send400)
	err := http.ListenAndServe("localhost:8080", mux)
	return err

}
