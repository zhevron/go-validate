package validators

import (
	"reflect"
	"strings"
	"testing"
)

func TestNil_True(t *testing.T) {
	n := error(nil)
	v := reflect.ValueOf(n)
	err := Nil("test", "true", v)
	if err != nil {
		t.Error(err)
	}
}

func TestNil_NotNil(t *testing.T) {
	n := 1
	v := reflect.ValueOf(n)
	err := Nil("test", "true", v)
	if err == nil {
		t.Fatal("Nil(true) did not error with a non-nil value")
	}
	if !strings.Contains(err.Error(), "is not nil") {
		t.Error(err)
	}
}

func TestNil_False(t *testing.T) {
	n := 1
	v := reflect.ValueOf(n)
	err := Nil("test", "false", v)
	if err != nil {
		t.Error(err)
	}
}

func TestNil_Zero(t *testing.T) {
	n := 0
	v := reflect.ValueOf(n)
	err := Nil("test", "false", v)
	if err == nil {
		t.Fatal("Nil(false) did not error with a zero value")
	}
	if !strings.Contains(err.Error(), "is zero") {
		t.Error(err)
	}
}
