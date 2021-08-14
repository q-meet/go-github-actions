package main

import "testing"

func TestW(t *testing.T) {
	res := w()
	if res != "Hello GitHub Actions. qikqiak.com is awesome1" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. qikqiak.com is awesome1", res)
	}
}

func TestGetVersion1(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{name: "test1", want: "0.0.1"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GetVersion(); got != tt.want {
				t.Errorf("GetVersion() = %v, want-v %v", got, tt.want)
			}
		})
	}
}
