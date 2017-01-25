package testlib

import (
	"testing"
)

func TestFullFunctionName(t *testing.T) {
	if expected, actual := "github.com/gigawattio/testlib.FullFunctionName", FullFunctionName(FullFunctionName); actual != expected {
		t.Fatalf("Expected value=%q but actual=%q", expected, actual)
	}
}

func TestFunctionName(t *testing.T) {
	if expected, actual := "FunctionName", FunctionName(FunctionName); actual != expected {
		t.Fatalf("Expected value=%q but actual=%q", expected, actual)
	}
}
