package validators

import (
	"fmt"
	"reflect"
	"strconv"
)

// MinLen validates the length of a reflect.Value to be at least
// the value of tag.
func MinLen(name string, tag string, v reflect.Value) error {
	min, err := strconv.ParseInt(tag, 10, 32)
	if err != nil {
		return fmt.Errorf(
			"The tag %#q could not be converted to a numeric value",
			tag,
		)
	}

	switch v.Kind() {
	case reflect.Array:
	case reflect.Slice:
		if len(v.Interface().([]interface{})) < int(min) {
			return fmt.Errorf("Value %#q is shorter than %d", name, min)
		}
		return nil

	case reflect.Map:
		if len(v.Interface().(map[interface{}]interface{})) < int(min) {
			return fmt.Errorf("Value %#q is shorter than %d", name, min)
		}
		return nil

	case reflect.String:
		if len(v.Interface().(string)) < int(min) {
			return fmt.Errorf("Value %#q is shorter than %d", name, min)
		}
		return nil
	}

	return fmt.Errorf("Value %#q cannot be checked for length", name)
}

// MaxLen validates the length of a reflect.Value to be no higher than
// the value of tag.
func MaxLen(name string, tag string, v reflect.Value) error {
	max, err := strconv.ParseInt(tag, 10, 32)
	if err != nil {
		return fmt.Errorf(
			"The tag %#q could not be converted to a numeric value",
			tag,
		)
	}

	switch v.Kind() {
	case reflect.Array:
	case reflect.Slice:
		if len(v.Interface().([]interface{})) > int(max) {
			return fmt.Errorf("Value %#q is longer than %d", name, max)
		}
		return nil

	case reflect.Map:
		if len(v.Interface().(map[interface{}]interface{})) > int(max) {
			return fmt.Errorf("Value %#q is longer than %d", name, max)
		}
		return nil

	case reflect.String:
		if len(v.Interface().(string)) > int(max) {
			return fmt.Errorf("Value %#q is longer than %d", name, max)
		}
		return nil
	}

	return fmt.Errorf("Value %#q cannot be checked for length", name)
}
