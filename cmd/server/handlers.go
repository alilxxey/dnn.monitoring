package main

import (
	"errors"
	"net/http"
	"strings"
)

func parseURL(r *http.Request) ([]string, error) {
	raw := strings.Trim(r.URL.Path, "/")
	url := strings.Split(raw, "/")
	urlLen := len(url)
	if urlLen != 4 {
		return []string{"error"}, errors.New("unexpected length of URL")
	}
	return []string{url[2], url[3]}, nil

}

func getGaugeMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, err := parseURL(r)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}

}

func getCounterMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	_, err := parseURL(r)
	if err == nil {
		w.WriteHeader(http.StatusOK)
	} else {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}
