package handlers

import (
	"errors"
	"github.com/alilxxey/dnn.monitoring/internal/interfaces"
	"net/http"
	"strconv"
	"strings"
)

type HTTPHandler struct {
	db interfaces.DB
}

func New(db interfaces.DB) *HTTPHandler {
	return &HTTPHandler{db: db}
}

func parseURL(r *http.Request) ([]string, error) {
	raw := strings.Trim(r.URL.Path, "/")
	url := strings.Split(raw, "/")
	urlLen := len(url)
	if urlLen != 4 {
		return []string{"error"}, errors.New("unexpected length of URL")
	}
	return []string{url[2], url[3]}, nil

}

func (h *HTTPHandler) GetGaugeMetric(w http.ResponseWriter, r *http.Request) {
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
		writeErr := h.db.WriteGauge(data[0], f)
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

func (h *HTTPHandler) GetCounterMetric(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
	data, err := parseURL(r)
	if err == nil {
		f, convErr := strconv.ParseInt(data[1], 10, 64)
		if convErr != nil {
			w.WriteHeader(http.StatusBadRequest)
			return
		}
		writeErr := h.db.Increment(data[0], f)
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
