package handlers

import (
	"github.com/alilxxey/dnn.monitoring/internal/interfaces"
	"github.com/alilxxey/dnn.monitoring/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"reflect"
	"testing"
)

func TestHTTPHandler_GetCounterMetric(t *testing.T) {
	type fields struct {
		db interfaces.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}

	tests := []struct {
		name     string
		endpoint string
		want     http.Response
	}{
		{
			name:     "wrong endpoint",
			endpoint: "/status/asdf/",
			want: http.Response{
				StatusCode: 404,
			},
		},
		{
			name:     "wrong endpoint",
			endpoint: "/status/counter/1",
			want: http.Response{
				StatusCode: 404,
			},
		},
		// TODO: Add test cases.
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, test.endpoint, nil)
			// создаём новый Recorder
			db := storage.New()
			w := httptest.NewRecorder()
			h := HTTPHandler{db: db}
			h.GetCounterMetric(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.StatusCode, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			// require.NoError(t, err)
		})
	}
}

func TestHTTPHandler_GetGaugeMetric(t *testing.T) {
	type fields struct {
		db interfaces.DB
	}
	type args struct {
		w http.ResponseWriter
		r *http.Request
	}
	tests := []struct {
		name   string
		fields fields
		args   args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			h := &HTTPHandler{
				db: tt.fields.db,
			}
			h.GetGaugeMetric(tt.args.w, tt.args.r)
		})
	}
}

func TestNew(t *testing.T) {
	type args struct {
		db interfaces.DB
	}
	tests := []struct {
		name string
		args args
		want *HTTPHandler
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(tt.args.db); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_parseURL(t *testing.T) {
	type args struct {
		r *http.Request
	}
	tests := []struct {
		name    string
		args    args
		want    []string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := parseURL(tt.args.r)
			if (err != nil) != tt.wantErr {
				t.Errorf("parseURL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseURL() got = %v, want %v", got, tt.want)
			}
		})
	}
}
