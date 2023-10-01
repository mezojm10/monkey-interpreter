package tests

import (
	"testing"

	"github.com/mezojm10/monkey-interpreter/internal/object"
)

func TestStringHashKey(t *testing.T) {
	hello1 := &object.String{Value: "Hello world"}
	hello2 := &object.String{Value: "Hello world"}
	diff1 := &object.String{Value: "My name is johnny"}
	diff2 := &object.String{Value: "My name is johnny"}

	if hello1.HashKey() != hello2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}

	if diff1.HashKey() != diff2.HashKey() {
		t.Errorf("strings with the same content have different hash keys")
	}

	if hello1.HashKey() == diff2.HashKey() {
		t.Errorf("strings with the different content have the same hash keys")
	}
}
