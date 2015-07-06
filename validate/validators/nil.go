package validators

import (
	"fmt"
	"reflect"
	"strings"
)

// Nil validates if a reflect.Value is nil or the zero value for it's type.
func Nil(name string, tag string, v reflect.Value) error {
	if strings.ToLower(tag) != "false" {
		if v.IsValid() {
			return fmt.Errorf("Value %#q is not nil", name)
		}
		return nil
	}
	if !v.IsValid() {
		return fmt.Errorf("Value %#q is nil", name)
	}
	if reflect.DeepEqual(v.Interface(), reflect.Zero(v.Type()).Interface()) {
		return fmt.Errorf("Value %#q is zero", name)
	}
	return nil
}
