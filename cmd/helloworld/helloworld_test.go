package main

import "testing"

func TestHello(t *testing.T) {

	emptyResult := hello("")

	if emptyResult != "Hello There!" {
		t.Errorf("hello(\"\") failed, expected %v, got %v", "Hello There!", emptyResult)
	}

}
