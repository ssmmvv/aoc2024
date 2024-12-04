package util

import (
	"testing"
)

func TestIntParsing(t *testing.T) {
	val := MustParseInt("123")
	if val != 123 {
		t.Errorf("Integer failed to parse, got %d, expected %d", val, 123)
	}
}

func TestIntParsePanic(t *testing.T) {
	defer func() {
		if r := recover(); r != nil {
			t.Log("Expected Panic was Caught")
		}
	}()
	MustParseInt("12ab34")
}
