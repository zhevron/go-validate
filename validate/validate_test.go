package validate

import (
	"errors"
	"testing"
)

type testStruct struct {
	String string `minLen:"10" maxLen:"50"`
	Int    int    `min:"5"`
	Error  error  `nil:"false"`
}

func TestValidate(t *testing.T) {
	s := &testStruct{
		String: "abcdefabcdef",
		Int:    10,
		Error:  errors.New("test"),
	}
	if _, errs := Validate(s); len(errs) > 0 {
		t.Errorf("Number of errors did not match. Got %d, expected 0", len(errs))
		for _, err := range errs {
			t.Errorf("  %s", err.Error())
		}
	}
}

func TestValidate_Errors(t *testing.T) {
	s := &testStruct{
		String: "abc",
		Int:    0,
		Error:  nil,
	}
	if _, errs := Validate(s); len(errs) != 3 {
		t.Errorf("Number of errors did not match. Got %d, expected 3", len(errs))
		for _, err := range errs {
			t.Errorf("  %s", err.Error())
		}
	}
}
