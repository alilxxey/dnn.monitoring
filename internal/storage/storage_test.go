package storage

import (
	"reflect"
	"testing"
)

func TestMemStorage_Exists(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    string
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			got, err := s.Exists(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("Exists() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("Exists() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStorage_GetValue(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name string
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    any
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			got, err := s.GetValue(tt.args.name)
			if (err != nil) != tt.wantErr {
				t.Errorf("GetValue() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("GetValue() got = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestMemStorage_Increment(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name  string
		value int64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			if err := s.Increment(tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("Increment() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestMemStorage_WriteGauge(t *testing.T) {
	type fields struct {
		gauge   map[string]float64
		counter map[string]int64
	}
	type args struct {
		name  string
		value float64
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &MemStorage{
				gauge:   tt.fields.gauge,
				counter: tt.fields.counter,
			}
			if err := s.WriteGauge(tt.args.name, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("WriteGauge() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestNew(t *testing.T) {
	tests := []struct {
		name string
		want *MemStorage
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := New(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("New() = %v, want %v", got, tt.want)
			}
		})
	}
}
