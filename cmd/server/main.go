package main

import (
	"github.com/alilxxey/dnn.monitoring/internal/handlers"
	"github.com/alilxxey/dnn.monitoring/internal/storage"
	"net/http"
)

func main() {
	if err := run(); err != nil {
		panic(err)
	}
}

func send404(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotFound)

}

func run() error {
	db := storage.New()
	h := handlers.New(db)
	mux := http.NewServeMux()

	mux.HandleFunc("/update/counter/", h.GetCounterMetric)
	mux.HandleFunc("/update/gauge/", h.GetGaugeMetric)
	mux.HandleFunc("/", send404)
	err := http.ListenAndServe("localhost:8080", mux)
	return err

}
