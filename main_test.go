package main

import "testing"

func TestW(t *testing.T) {
	res := w()
	if res != "Hello GitHub Actions. qikqiak.com is awesome" {
		t.Errorf("Greet() = %s; Expected Hello GitHub Actions. qikqiak.com is awesome", res)
	}
}
