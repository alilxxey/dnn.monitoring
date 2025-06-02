package handlers

import (
	"github.com/alilxxey/dnn.monitoring/internal/storage"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"net/url"
	"reflect"
	"testing"
)

func TestHTTPHandler_GetCounterMetric(t *testing.T) {
	tests := []struct {
		name     string
		endpoint string
		want     http.Response
	}{
		{
			name:     "wrong endpoint/no metric type",
			endpoint: "/update/",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "wrong endpoint/wrong metric type`",
			endpoint: "/update/ggg",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "wrong metric value type",
			endpoint: "/update/counter/1.1",
			want: http.Response{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			name:     "wrong endpoint/big depth",
			endpoint: "/update/gauge/1.44/asdf",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "send gauge metric to counter",
			endpoint: "/update/gauge/1.22",
			want: http.Response{
				StatusCode: http.StatusBadRequest,
			},
		},
		{
			name:     "OK/send counter metric",
			endpoint: "/update/counter/444",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send long counter metric",
			endpoint: "/update/counter/44222222222224",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send negative counter metric",
			endpoint: "/update/counter/-12333",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send counter metric with leading 0",
			endpoint: "/update/counter/-12333",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
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
	tests := []struct {
		name     string
		endpoint string
		want     http.Response
	}{
		{
			name:     "wrong endpoint/no metric type",
			endpoint: "/update/",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "wrong endpoint/wrong metric type`",
			endpoint: "/update/ggg",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "wrong endpoint/big depth",
			endpoint: "/update/gauge/1.44/asdf",
			want: http.Response{
				StatusCode: http.StatusNotFound,
			},
		},
		{
			name:     "OK/send gauge metric",
			endpoint: "/update/gauge/1.22",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "send counter metric to gauge",
			endpoint: "/update/gauge/444",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send int gauge metric",
			endpoint: "/update/gauge/122",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send negative gauge metric",
			endpoint: "/update/gauge/-1.4422",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send long metric",
			endpoint: "/update/gauge/1123123123.22443434",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
		{
			name:     "OK/send gauge metric as .1",
			endpoint: "/update/gauge/.22",
			want: http.Response{
				StatusCode: http.StatusOK,
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			request := httptest.NewRequest(http.MethodPost, test.endpoint, nil)
			// создаём новый Recorder
			db := storage.New()
			w := httptest.NewRecorder()
			h := HTTPHandler{db: db}
			h.GetGaugeMetric(w, request)

			res := w.Result()
			// проверяем код ответа
			assert.Equal(t, test.want.StatusCode, res.StatusCode)
			// получаем и проверяем тело запроса
			defer res.Body.Close()
			// require.NoError(t, err)
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
		{
			name: "ERR/#1",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "/update/counter/",
					},
				},
			},
			want:    []string{"error"},
			wantErr: true,
		},
		{
			name: "OK/#2",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "/update/counter/12",
					},
				},
			},
			want:    []string{"counter", "12"},
			wantErr: false,
		},
		{
			name: "ERR/#3",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "/update/counter/123/ggg",
					},
				},
			},
			want:    []string{"error"},
			wantErr: true,
		},
		{
			name: "OK/#4",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "/update/gauge/123123",
					},
				},
			},
			want:    []string{"gauge", "123123"},
			wantErr: false,
		},
		{
			name: "ERR/#5",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "/update/gauge/11111/12",
					},
				},
			},
			want:    []string{"error"},
			wantErr: true,
		},
		{
			name: "OK/#6",
			args: args{
				&http.Request{
					URL: &url.URL{
						Path: "///",
					},
				},
			},
			want:    []string{"error"},
			wantErr: true,
		},
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
