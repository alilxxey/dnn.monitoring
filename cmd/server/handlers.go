package main

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"
)

func parseURL(r *http.Request) ([]string, error) {
	raw := strings.Trim(r.URL.Path, "/")
	url := strings.Split(raw, "/")
	urlLen := len(url)
	if urlLen != 4 {
		return []string{"error"}, errors.New("unexpected length of URL\n")
	}
	return []string{url[2], url[3]}, nil

}

func getGaugeMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data, err := parseURL(r)
	if err == nil {
		f, convErr := strconv.ParseFloat(data[1], 64)
		if convErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		writeErr := DataBase.WriteGauge(data[0], f)
		if writeErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return

		}
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}

}

func getCounterMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data, err := parseURL(r)
	if err == nil {
		fmt.Println(data)
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte(err.Error()))
		return
	}
}
