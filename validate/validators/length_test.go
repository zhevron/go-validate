package validators

import (
	"reflect"
	"strings"
	"testing"
)

func TestMinLen(t *testing.T) {
	s := "abc"
	v := reflect.ValueOf(s)
	err := MinLen("test", "1", v)
	if err != nil {
		t.Error(err)
	}
}

func TestMinLen_Error(t *testing.T) {
	s := "abc"
	v := reflect.ValueOf(s)
	err := MinLen("test", "10", v)
	if err == nil {
		t.Fatal("MinLen did not error with a lower value")
	}
	if !strings.Contains(err.Error(), "is shorter than") {
		t.Error(err)
	}
}

func TestMinLen_Invalid(t *testing.T) {
	n := 1
	v := reflect.ValueOf(n)
	err := MinLen("test", "1", v)
	if err == nil {
		t.Fatal("MinLen did not error with an invalid type")
	}
	if !strings.Contains(err.Error(), "cannot be checked for length") {
		t.Error(err)
	}
}

func TestMaxLen(t *testing.T) {
	s := "abc"
	v := reflect.ValueOf(s)
	err := MaxLen("test", "10", v)
	if err != nil {
		t.Error(err)
	}
}

func TestMaxLen_Error(t *testing.T) {
	s := "abc"
	v := reflect.ValueOf(s)
	err := MaxLen("test", "1", v)
	if err == nil {
		t.Fatal("MaxLen did not error with an higher value")
	}
	if !strings.Contains(err.Error(), "is longer than") {
		t.Error(err)
	}
}

func TestMaxLen_Invalid(t *testing.T) {
	n := 1
	v := reflect.ValueOf(n)
	err := MaxLen("test", "1", v)
	if err == nil {
		t.Fatal("MaxLen did not error with an invalid type")
	}
	if !strings.Contains(err.Error(), "cannot be checked for length") {
		t.Error(err)
	}
}
